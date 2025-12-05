package handlers

import (
	"net/http"
	"net/url"

	"code.cloudfoundry.org/korifi/api/presenter"
	"code.cloudfoundry.org/korifi/api/routing"
)

const (
	OrgQuotaPath = "/v3/organization_quotas/${guid}"
)

type CFOrgQuotaRepository interface{}

type OrgQuota struct {
	apiBaseURL   url.URL
	orgQuotaRepo CFOrgQuotaRepository
}

func NewOrgQuota(apiBaseURL url.URL, orgQuotaRepo CFOrgQuotaRepository) *OrgQuota {
	return &OrgQuota{
		apiBaseURL:   apiBaseURL,
		orgQuotaRepo: orgQuotaRepo,
	}
}

func (h *OrgQuota) get(r *http.Request) (*routing.Response, error) {
	return routing.NewResponse(http.StatusOK).WithBody(presenter.ForOrgQuota(h.apiBaseURL)), nil
}

func (h *OrgQuota) AuthenticatedRoutes() []routing.Route {
	return []routing.Route{
		{Method: "GET", Pattern: OrgQuotaPath, Handler: h.get},
	}
}

func (h *OrgQuota) UnauthenticatedRoutes() []routing.Route {
	return nil
}
