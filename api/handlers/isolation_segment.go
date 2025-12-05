package handlers

import (
	"net/http"
	"net/url"

	"code.cloudfoundry.org/korifi/api/presenter"
	"code.cloudfoundry.org/korifi/api/routing"
)

const (
	IsolationSegmentsPath = "/v3/isolation_segments"

// IsolationSegmentPath  = "/v3/isolation_segments/{guid}"
)

type CFIsolationSegmentRepository interface{}

type IsolationSegment struct {
	apiBaseURL           url.URL
	isolationSegmentRepo CFIsolationSegmentRepository
}

func NewIsolationSegment(apiBaseURL url.URL, isolationSegmentRepo CFIsolationSegmentRepository) *IsolationSegment {
	return &IsolationSegment{
		apiBaseURL:           apiBaseURL,
		isolationSegmentRepo: isolationSegmentRepo,
	}
}

//func (h *IsolationSegment) get(r *http.Request) (*routing.Response, error) {
//	return routing.NewResponse(http.StatusOK).WithBody(presenter.ForIsolationSegment(h.apiBaseURL)), nil
//}

func (h *IsolationSegment) list(r *http.Request) (*routing.Response, error) {
	return routing.NewResponse(http.StatusOK).WithBody(presenter.ForIsolationSegment(h.apiBaseURL)), nil
}

func (h *IsolationSegment) AuthenticatedRoutes() []routing.Route {
	return []routing.Route{
		{Method: "GET", Pattern: IsolationSegmentsPath, Handler: h.list},
		//		{Method: "GET", Pattern: IsolationSegmentPath, Handler: h.get},
	}
}

func (h *IsolationSegment) UnauthenticatedRoutes() []routing.Route {
	return nil
}
