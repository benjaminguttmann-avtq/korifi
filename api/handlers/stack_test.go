package handlers_test

import (
	"errors"
	"net/http"
	"time"

	. "code.cloudfoundry.org/korifi/api/handlers"
	"code.cloudfoundry.org/korifi/api/handlers/fake"
	"code.cloudfoundry.org/korifi/api/payloads"
	"code.cloudfoundry.org/korifi/api/repositories"
	. "code.cloudfoundry.org/korifi/tests/matchers"
	"code.cloudfoundry.org/korifi/tools"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	var (
		stackRepo        *fake.StackRepository
		req              *http.Request
		requestValidator *fake.RequestValidator
	)

	BeforeEach(func() {
		stackRepo = new(fake.StackRepository)
		requestValidator = new(fake.RequestValidator)

		apiHandler := NewStack(*serverURL, stackRepo, requestValidator)
		routerBuilder.LoadRoutes(apiHandler)
	})

	JustBeforeEach(func() {
		routerBuilder.Build().ServeHTTP(rr, req)
	})

	Describe("the GET /v3/stacks endpoint", func() {
		BeforeEach(func() {
			stackRepo.ListStacksReturns([]repositories.StackRecord{
				{
					Name:        "io.buildpacks.stacks.jammy",
					Description: "Jammy Stack",
					CreatedAt:   time.UnixMilli(1000),
					UpdatedAt:   tools.PtrTo(time.UnixMilli(2000)),
				},
				{
					Name:        "io.buildpacks.stacks.noble",
					Description: "Noble Stack",
					CreatedAt:   time.UnixMilli(1000),
					UpdatedAt:   tools.PtrTo(time.UnixMilli(2000)),
				},
			}, nil)

			var err error
			req, err = http.NewRequestWithContext(ctx, "GET", "/v3/stacks", nil)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the stacks for the default builder", func() {
			Expect(stackRepo.ListStacksCallCount()).To(Equal(1))
			_, actualAuthInfo, _ := stackRepo.ListStacksArgsForCall(0)
			Expect(actualAuthInfo).To(Equal(authInfo))

			Expect(rr).To(HaveHTTPStatus(http.StatusOK))
			Expect(rr).To(HaveHTTPHeaderWithValue("Content-Type", "application/json"))
			Expect(rr).To(HaveHTTPBody(SatisfyAll(
				MatchJSONPath("$.pagination.total_results", BeEquivalentTo(2)),
				MatchJSONPath("$.pagination.first.href", "https://api.example.org/v3/stacks"),
				MatchJSONPath("$.resources", HaveLen(2)),
				MatchJSONPath("$.resources[0].name", "io.buildpacks.stacks.jammy"),
			)))
		})
		When("there is some other error fetching the stacks", func() {
			BeforeEach(func() {
				stackRepo.ListStacksReturns([]repositories.StackRecord{}, errors.New("unknown!"))
			})

			It("returns an error", func() {
				expectUnknownError()
			})
		})

		When("filtering query params are provided", func() {
			BeforeEach(func() {
				requestValidator.DecodeAndValidateURLValuesStub = decodeAndValidateURLValuesStub(&payloads.StackList{
					Names: "a1,a2",
				})
			})

			It("passes them to the repository", func() {
				Expect(stackRepo.ListStacksCallCount()).To(Equal(1))
				_, _, message := stackRepo.ListStacksArgsForCall(0)

				Expect(message.Names).To(ConsistOf("a1", "a2"))
			})
		})
	})
})
