package presenter

import (
	"net/url"

	"code.cloudfoundry.org/korifi/api/repositories"
	// "code.cloudfoundry.org/korifi/api/repositories/include"
)

// TODO: repetition with handler endpoint?
const isolationSegmentsBase = "/v3/isolation_segments"

type IsolationSegmentResponse struct{}

func ForIsolationSegment(isolationSegment repositories.IsolationSegmentRecord, apiBaseURL url.URL) IsolationSegmentResponse {
	return IsolationSegmentResponse{}
}
