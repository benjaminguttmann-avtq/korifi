package presenter

import (
	"net/url"
)

type OrgQuotaResponse struct{}

func ForOrgQuota(apiBaseURL url.URL) OrgQuotaResponse {
	return OrgQuotaResponse{}
}
