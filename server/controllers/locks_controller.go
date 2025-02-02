package controllers

import (
	"fmt"
	"github.com/runatlantis/atlantis/server/controllers/templates"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/runatlantis/atlantis/server/core/locking"
	"github.com/runatlantis/atlantis/server/events"
	"github.com/runatlantis/atlantis/server/events/models"
	"github.com/runatlantis/atlantis/server/events/vcs"
	"github.com/runatlantis/atlantis/server/logging"
)

// LocksController handles all requests relating to Atlantis locks.
type LocksController struct {
	AtlantisVersion    string
	AtlantisURL        *url.URL
	Locker             locking.Locker
	Logger             logging.SimpleLogging
	ApplyLocker        locking.ApplyLocker
	VCSClient          vcs.Client
	LockDetailTemplate templates.TemplateWriter
	WorkingDir         events.WorkingDir
	WorkingDirLocker   events.WorkingDirLocker
	Backend            locking.Backend
	DeleteLockCommand  events.DeleteLockCommand
}

// LockApply handles creating a global apply lock.
// If Lock already exists it will be a no-op
func (l *LocksController) LockApply(w http.ResponseWriter, r *http.Request) {
	lock, err := l.ApplyLocker.LockApply()
	if err != nil {
		l.respond(w, logging.Error, http.StatusInternalServerError, "creating apply lock failed with: %s", err)
		return
	}

	l.respond(w, logging.Info, http.StatusOK, "Apply Lock is acquired on %s", lock.Time.Format("2006-01-02 15:04:05"))
}

// UnlockApply handles releasing a global apply lock.
// If Lock doesn't exists it will be a no-op
func (l *LocksController) UnlockApply(w http.ResponseWriter, r *http.Request) {
	err := l.ApplyLocker.UnlockApply()
	if err != nil {
		l.respond(w, logging.Error, http.StatusInternalServerError, "deleting apply lock failed with: %s", err)
		return
	}

	l.respond(w, logging.Info, http.StatusOK, "Deleted apply lock")
}

// GetLock is the GET /locks/{id} route. It renders the lock detail view.
func (l *LocksController) GetLock(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		l.respond(w, logging.Warn, http.StatusBadRequest, "No lock id in request")
		return
	}

	idUnencoded, err := url.QueryUnescape(id)
	if err != nil {
		l.respond(w, logging.Warn, http.StatusBadRequest, "Invalid lock id: %s", err)
		return
	}
	lock, err := l.Locker.GetLock(idUnencoded)
	if err != nil {
		l.respond(w, logging.Error, http.StatusInternalServerError, "Failed getting lock: %s", err)
		return
	}
	if lock == nil {
		l.respond(w, logging.Info, http.StatusNotFound, "No lock found at id %q", idUnencoded)
		return
	}

	// get queues locks for this lock details page
	var queue models.ProjectLockQueue
	queue, _ = l.Locker.GetQueueByLock(lock.Project, lock.Workspace)
	lockDetailQueue := GetQueueItemIndexData(queue)
	owner, repo := models.SplitRepoFullName(lock.Project.RepoFullName)
	viewData := templates.LockDetailData{
		LockKeyEncoded:  id,
		LockKey:         idUnencoded,
		PullRequestLink: lock.Pull.URL,
		LockedBy:        lock.Pull.Author,
		Workspace:       lock.Workspace,
		AtlantisVersion: l.AtlantisVersion,
		CleanedBasePath: l.AtlantisURL.Path,
		RepoOwner:       owner,
		RepoName:        repo,
		Queue:           lockDetailQueue,
	}

	err = l.LockDetailTemplate.Execute(w, viewData)
	if err != nil {
		l.Logger.Err(err.Error())
	}
}

// DeleteLock handles deleting the lock at id and commenting back on the
// pull request that the lock has been deleted.
func (l *LocksController) DeleteLock(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok || id == "" {
		l.respond(w, logging.Warn, http.StatusBadRequest, "No lock id in request")
		return
	}

	idUnencoded, err := url.PathUnescape(id)
	if err != nil {
		l.respond(w, logging.Warn, http.StatusBadRequest, "Invalid lock id %q. Failed with error: %s", id, err)
		return
	}

	lock, dequeuedLock, err := l.DeleteLockCommand.DeleteLock(idUnencoded)
	if err != nil {
		l.respond(w, logging.Error, http.StatusInternalServerError, "deleting lock failed with: %s", err)
		return
	}

	if lock == nil {
		l.respond(w, logging.Info, http.StatusNotFound, "No lock found at id %q", idUnencoded)
		return
	}

	// NOTE: Because BaseRepo was added to the PullRequest model later, previous
	// installations of Atlantis will have locks in their DB that do not have
	// this field on PullRequest. We skip commenting in this case.
	if lock.Pull.BaseRepo != (models.Repo{}) {
		if err := l.Backend.UpdateProjectStatus(lock.Pull, lock.Workspace, lock.Project.Path, models.DiscardedPlanStatus); err != nil {
			l.Logger.Err("unable to update project status: %s", err)
		}

		// Once the lock has been deleted, comment back on the pull request.
		comment := fmt.Sprintf("**Warning**: The plan for dir: `%s` workspace: `%s` was **discarded** via the Atlantis UI.\n\n"+
			"To `apply` this plan you must run `plan` again.", lock.Project.Path, lock.Workspace)
		if err = l.VCSClient.CreateComment(lock.Pull.BaseRepo, lock.Pull.Num, comment, ""); err != nil {
			l.Logger.Warn("failed commenting on pull request: %s", err)
		}
		if dequeuedLock != nil {
			l.Logger.Warn("dequeued lock: %s", dequeuedLock)
			l.commentOnDequeuedPullRequests(*dequeuedLock)
		}
	} else {
		l.Logger.Debug("skipping commenting on pull request and deleting workspace because BaseRepo field is empty")
	}
	l.respond(w, logging.Info, http.StatusOK, "Deleted lock id %q", id)
}

// respond is a helper function to respond and log the response. lvl is the log
// level to log at, code is the HTTP response code.
func (l *LocksController) respond(w http.ResponseWriter, lvl logging.LogLevel, responseCode int, format string, args ...interface{}) {
	response := fmt.Sprintf(format, args...)
	l.Logger.Log(lvl, response)
	w.WriteHeader(responseCode)
	fmt.Fprintln(w, response)
}

func (l *LocksController) commentOnDequeuedPullRequests(dequeuedLock models.ProjectLock) {
	planVcsMessage := models.BuildCommentOnDequeuedPullRequest([]models.ProjectLock{dequeuedLock})
	if commentErr := l.VCSClient.CreateComment(dequeuedLock.Pull.BaseRepo, dequeuedLock.Pull.Num, planVcsMessage, ""); commentErr != nil {
		l.Logger.Err("unable to comment on PR %d: %s", dequeuedLock.Pull.Num, commentErr)
	}
}

func GetQueueItemIndexData(q models.ProjectLockQueue) []templates.QueueItemIndexData {
	var queueIndexDataList []templates.QueueItemIndexData
	for _, projectLock := range q {
		queueIndexDataList = append(queueIndexDataList, templates.QueueItemIndexData{
			LockPath:      "Not yet acquired",
			RepoFullName:  projectLock.Project.RepoFullName,
			PullNum:       projectLock.Pull.Num,
			Path:          projectLock.Project.Path,
			Workspace:     projectLock.Workspace,
			Time:          projectLock.Time,
			TimeFormatted: projectLock.Time.Format("02-01-2006 15:04:05"),
			PullURL:       projectLock.Pull.URL,
			Author:        projectLock.Pull.Author,
		})
	}
	return queueIndexDataList
}
