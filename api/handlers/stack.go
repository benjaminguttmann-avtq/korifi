package handlers

import (
	"context"
	"net/http"
	"net/url"

	"code.cloudfoundry.org/korifi/api/authorization"
	apierrors "code.cloudfoundry.org/korifi/api/errors"
	"code.cloudfoundry.org/korifi/api/payloads"
	"code.cloudfoundry.org/korifi/api/presenter"
	"code.cloudfoundry.org/korifi/api/repositories"
	"code.cloudfoundry.org/korifi/api/routing"

	"github.com/go-logr/logr"
)

const (
	StacksPath = "/v3/stacks"
)

type StackRepository interface {
	ListStacks(ctx context.Context, authInfo authorization.Info, messages repositories.ListStacksMessage) ([]repositories.StackRecord, error)
}

type Stack struct {
	serverURL        url.URL
	stackRepo        StackRepository
	requestValidator RequestValidator
}

func NewStack(
	serverURL url.URL,
	stackRepo StackRepository,
	requestValidator RequestValidator,
) *Stack {
	return &Stack{
		serverURL:        serverURL,
		stackRepo:        stackRepo,
		requestValidator: requestValidator,
	}
}

func (h *Stack) list(r *http.Request) (*routing.Response, error) {
	authInfo, _ := authorization.InfoFromContext(r.Context())
	logger := logr.FromContextOrDiscard(r.Context()).WithName("handlers.build.list")

	payload := new(payloads.StackList)
	err := h.requestValidator.DecodeAndValidateURLValues(r, payload)
	if err != nil {
		return nil, apierrors.LogAndReturn(logger, err, "Unable to decode request query parameters")
	}

	//	stacks, err := h.stackRepo.ListStacks(r.Context(), authInfo)
	stackList, err := h.stackRepo.ListStacks(r.Context(), authInfo, payload.ToMessage())
	if err != nil {
		return nil, apierrors.LogAndReturn(logger, err, "Failed to fetch buildpacks from Kubernetes")
	}

	//	return routing.NewResponse(http.StatusOK).WithBody(presenter.ForList(presenter.ForStack, stacks, h.serverURL, *r.URL)), nil
	return routing.NewResponse(http.StatusOK).WithBody(presenter.ForList(presenter.ForStack, stackList, h.serverURL, *r.URL)), nil
}

func (h *Stack) UnauthenticatedRoutes() []routing.Route {
	return nil
}

func (h *Stack) AuthenticatedRoutes() []routing.Route {
	return []routing.Route{
		{Method: "GET", Pattern: StacksPath, Handler: h.list},
	}
}
