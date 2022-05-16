// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/locking (interfaces: Locker)

package mocks

import (
	"reflect"
	"time"

	pegomock "github.com/petergtz/pegomock"
	locking "github.com/runatlantis/atlantis/server/core/locking"
	models "github.com/runatlantis/atlantis/server/events/models"
)

type MockLocker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockLocker(options ...pegomock.Option) *MockLocker {
	mock := &MockLocker{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockLocker) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockLocker) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockLocker) TryLock(p models.Project, workspace string, pull models.PullRequest, user models.User) (locking.TryLockResponse, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{p, workspace, pull, user}
	result := pegomock.GetGenericMockFrom(mock).Invoke("TryLock", params, []reflect.Type{reflect.TypeOf((*locking.TryLockResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 locking.TryLockResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(locking.TryLockResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) Unlock(key string) (*models.ProjectLock, *models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{key}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Unlock", params, []reflect.Type{reflect.TypeOf((**models.ProjectLock)(nil)).Elem(), reflect.TypeOf((**models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *models.ProjectLock
	var ret1 *models.ProjectLock
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(*models.ProjectLock)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockLocker) List() (map[string]models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("List", params, []reflect.Type{reflect.TypeOf((*map[string]models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 map[string]models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string]models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) ListQueues() (map[string][]models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ListQueues", params, []reflect.Type{reflect.TypeOf((*map[string][]models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 map[string][]models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string][]models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) GetQueueByLock(project models.Project, workspace string) ([]models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{project, workspace}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetQueueByLock", params, []reflect.Type{reflect.TypeOf((*[]models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) UnlockByPull(repoFullName string, pullNum int) ([]models.ProjectLock, models.DequeueStatus, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{repoFullName, pullNum}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UnlockByPull", params, []reflect.Type{reflect.TypeOf((*[]models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*models.DequeueStatus)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []models.ProjectLock
	var ret1 models.DequeueStatus
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(models.DequeueStatus)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockLocker) GetLock(key string) (*models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{key}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetLock", params, []reflect.Type{reflect.TypeOf((**models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) VerifyWasCalledOnce() *VerifierMockLocker {
	return &VerifierMockLocker{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockLocker) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockLocker {
	return &VerifierMockLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockLocker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockLocker {
	return &VerifierMockLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockLocker) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockLocker {
	return &VerifierMockLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockLocker struct {
	mock                   *MockLocker
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockLocker) TryLock(p models.Project, workspace string, pull models.PullRequest, user models.User) *MockLocker_TryLock_OngoingVerification {
	params := []pegomock.Param{p, workspace, pull, user}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "TryLock", params, verifier.timeout)
	return &MockLocker_TryLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_TryLock_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_TryLock_OngoingVerification) GetCapturedArguments() (models.Project, string, models.PullRequest, models.User) {
	p, workspace, pull, user := c.GetAllCapturedArguments()
	return p[len(p)-1], workspace[len(workspace)-1], pull[len(pull)-1], user[len(user)-1]
}

func (c *MockLocker_TryLock_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Project, _param1 []string, _param2 []models.PullRequest, _param3 []models.User) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Project, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Project)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(models.PullRequest)
		}
		_param3 = make([]models.User, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(models.User)
		}
	}
	return
}

func (verifier *VerifierMockLocker) Unlock(key string) *MockLocker_Unlock_OngoingVerification {
	params := []pegomock.Param{key}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Unlock", params, verifier.timeout)
	return &MockLocker_Unlock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_Unlock_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_Unlock_OngoingVerification) GetCapturedArguments() string {
	key := c.GetAllCapturedArguments()
	return key[len(key)-1]
}

func (c *MockLocker_Unlock_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockLocker) List() *MockLocker_List_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "List", params, verifier.timeout)
	return &MockLocker_List_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_List_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_List_OngoingVerification) GetCapturedArguments() {
}

func (c *MockLocker_List_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockLocker) ListQueues() *MockLocker_ListQueues_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ListQueues", params, verifier.timeout)
	return &MockLocker_ListQueues_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_ListQueues_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_ListQueues_OngoingVerification) GetCapturedArguments() {
}

func (c *MockLocker_ListQueues_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockLocker) GetQueueByLock(project models.Project, workspace string) *MockLocker_GetQueueByLock_OngoingVerification {
	params := []pegomock.Param{project, workspace}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetQueueByLock", params, verifier.timeout)
	return &MockLocker_GetQueueByLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_GetQueueByLock_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_GetQueueByLock_OngoingVerification) GetCapturedArguments() (models.Project, string) {
	project, workspace := c.GetAllCapturedArguments()
	return project[len(project)-1], workspace[len(workspace)-1]
}

func (c *MockLocker_GetQueueByLock_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Project, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Project, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Project)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockLocker) UnlockByPull(repoFullName string, pullNum int) *MockLocker_UnlockByPull_OngoingVerification {
	params := []pegomock.Param{repoFullName, pullNum}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UnlockByPull", params, verifier.timeout)
	return &MockLocker_UnlockByPull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_UnlockByPull_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_UnlockByPull_OngoingVerification) GetCapturedArguments() (string, int) {
	repoFullName, pullNum := c.GetAllCapturedArguments()
	return repoFullName[len(repoFullName)-1], pullNum[len(pullNum)-1]
}

func (c *MockLocker_UnlockByPull_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}

func (verifier *VerifierMockLocker) GetLock(key string) *MockLocker_GetLock_OngoingVerification {
	params := []pegomock.Param{key}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetLock", params, verifier.timeout)
	return &MockLocker_GetLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_GetLock_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_GetLock_OngoingVerification) GetCapturedArguments() string {
	key := c.GetAllCapturedArguments()
	return key[len(key)-1]
}

func (c *MockLocker_GetLock_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}
