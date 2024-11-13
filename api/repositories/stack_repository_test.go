package repositories_test

import (
	"context"

	. "code.cloudfoundry.org/korifi/api/repositories"
	"code.cloudfoundry.org/korifi/api/repositories/fake"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("StackRepository", func() {
	var (
		stackRepo *StackRepository
		sorter    *fake.StackSorter
	)

	BeforeEach(func() {
		sorter = new(fake.StackSorter)
		sorter.SortStub = func(records []StackRecord, _ string) []StackRecord {
			return records
		}

		stackRepo = NewStackRepository(
			builderNames,
			userClientFactory,
			rootNamespace,
			sorter,
		)
	})
	Describe("ListStacks", func() {
		var message ListStacksMessage

		BeforeEach(func() {
			message = ListStacksMessage{OrderBy: "names"}
		})

		When("a controller with the configured BuilderName exists", func() {
			var stacks []StackRecord

			BeforeEach(func() {
				createBuilderInfoWithCleanup(ctx, builderNames[0], "io.buildpack.stacks.jammy", []buildpackInfo{
					{name: "paketo-stacks/stack-1-1", version: "1.1"},
					{name: "paketo-stacks/stack-1-2", version: "1.2"},
				})
				var err error
				stacks, err = stackRepo.ListStacks(context.Background(), authInfo, message)
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns all stacks", func() {
				Expect(stacks).To(ConsistOf(
					MatchFields(IgnoreExtras, Fields{
						"Name": Equal("io.buildpack.stacks.jammy"),
					}),
				))
			})

			It("sorts the stacks", func() {
				Expect(sorter.SortCallCount()).To(Equal(1))
				sortedStacks, field := sorter.SortArgsForCall(0)
				Expect(field).To(Equal("names"))
				Expect(sortedStacks).To(ConsistOf(
					MatchFields(IgnoreExtras, Fields{
						"Name": Equal("io.buildpack.stacks.jammy"),
					}),
				))
			})
		})
	})
})
