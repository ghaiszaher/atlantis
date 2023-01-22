// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsPullReqStatus() models.PullReqStatus {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.PullReqStatus))(nil)).Elem()))
	var nullValue models.PullReqStatus
	return nullValue
}

func EqModelsPullReqStatus(value models.PullReqStatus) models.PullReqStatus {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.PullReqStatus
	return nullValue
}

func NotEqModelsPullReqStatus(value models.PullReqStatus) models.PullReqStatus {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue models.PullReqStatus
	return nullValue
}

func ModelsPullReqStatusThat(matcher pegomock.ArgumentMatcher) models.PullReqStatus {
	pegomock.RegisterMatcher(matcher)
	var nullValue models.PullReqStatus
	return nullValue
}
