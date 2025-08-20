package handlers

import (
	"context"
	//	"fmt"
	//	"net/http"
	//	"net/url"
	//
	//	"github.com/go-logr/logr"
	//
	"code.cloudfoundry.org/korifi/api/authorization"
	//	apierrors "code.cloudfoundry.org/korifi/api/errors"
	//	"code.cloudfoundry.org/korifi/api/payloads"
	//	"code.cloudfoundry.org/korifi/api/presenter"
	"code.cloudfoundry.org/korifi/api/repositories"
)

const (
	IsolationSegmentsPath = "/v3/isolation_segments"
)

//counterfeiter:generate -o fake -fake-name CFIsolationSegmentRepository . CFIsolationSegmentRepository

type CFIsolationSegmentRepository interface {
	ListIsolationSegments(context.Context, authorization.Info) (repositories.ListResult[repositories.IsolationSegmentRecord], error)
}
