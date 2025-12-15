package handlers

import (
	"net/http"
	"net/url"

	"code.cloudfoundry.org/korifi/api/presenter"
	"code.cloudfoundry.org/korifi/api/repositories"
	"code.cloudfoundry.org/korifi/api/routing"
)

const (
	SpaceQuotasPath = "/v3/space_quotas"
)

type SpaceQuota struct {
	apiBaseURL url.URL
}

func NewSpaceQuota(apiBaseURL url.URL) *SpaceQuota {
	return &SpaceQuota{
		apiBaseURL: apiBaseURL,
	}
}

func (h *SpaceQuota) create(r *http.Request) (*routing.Response, error) {
	authInfo, _ := authorization.InfoFromContext(r.Context())
	logger := logr.FromContextOrDiscard(r.Context()).WithName("handlers.space-quota.create")

	payload := new(payloads.SpaceQuotaCreate)
	if err := h.requestValidator.DecodeAndValidateJSONPayload(r, payload); err != nil {
		return nil, apierrors.LogAndReturn(logger, err, "failed to decode payload")
	}

	spaceQuota, err := h.spaceQuotaRepo.CreateSpaceQuota(r.Context(), authInfo, payload.ToMessage())
	if err != nil {
		return nil, apierrors.LogAndReturn(logger, err, "Failed to create space group", "Space Quota Name", payload.DisplayName)
	}

	return routing.NewResponse(http.StatusCreated).WithBody(presenter.ForSpaceQuota(spaceQuota, h.serverURL)), nil
}

func (h *SpaceQuota) AuthenticatedRoutes() []routing.Route {
	return []routing.Route{
		{Method: "POST", Pattern: SpaceQuotasPath, Handler: h.create},
	}
}

func (h *SpaceQuota) UnauthenticatedRoutes() []routing.Route {
	return nil
}
