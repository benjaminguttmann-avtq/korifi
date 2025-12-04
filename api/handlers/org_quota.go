package handlers

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-logr/logr"

	"code.cloudfoundry.org/korifi/api/authorization"
	apierrors "code.cloudfoundry.org/korifi/api/errors"
	"code.cloudfoundry.org/korifi/api/payloads"
	"code.cloudfoundry.org/korifi/api/presenter"
	"code.cloudfoundry.org/korifi/api/repositories"
	"code.cloudfoundry.org/korifi/api/routing"
	"code.cloudfoundry.org/korifi/api/tools/singleton"
)

const (
	OrgQuotaPath = " /v3/organization_quotas/${guid}"
)

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
