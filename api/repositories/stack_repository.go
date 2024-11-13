package repositories

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"code.cloudfoundry.org/korifi/api/authorization"
	apierrors "code.cloudfoundry.org/korifi/api/errors"
	"code.cloudfoundry.org/korifi/api/repositories/compare"

	korifiv1alpha1 "code.cloudfoundry.org/korifi/controllers/api/v1alpha1"
	"code.cloudfoundry.org/korifi/tools"
	"github.com/BooleanCat/go-functional/v2/it"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/types"
)

const (
	StackResourceType = "Stack"
)

type StackRepository struct {
	builderNames      []string
	userClientFactory authorization.UserK8sClientFactory
	rootNamespace     string
	sorter            StackSorter
}

type StackSorter interface {
	Sort(records []StackRecord, order string) []StackRecord
}

type stackSorter struct {
	sorter *compare.Sorter[StackRecord]
}

func NewStackSorter() *stackSorter {
	return &stackSorter{
		sorter: compare.NewSorter(StackComparator),
	}
}

func (s *stackSorter) Sort(records []StackRecord, order string) []StackRecord {
	return s.sorter.Sort(records, order)
}

func StackComparator(fieldName string) func(StackRecord, StackRecord) int {
	return func(s1, s2 StackRecord) int {
		switch fieldName {
		case "", "name":
			return strings.Compare(s1.Name, s2.Name)
		case "-name":
			return strings.Compare(s2.Name, s1.Name)
		case "created_at":
			return tools.CompareTimePtr(&s1.CreatedAt, &s2.CreatedAt)
		case "-created_at":
			return tools.CompareTimePtr(&s2.CreatedAt, &s1.CreatedAt)
		case "updated_at":
			return tools.CompareTimePtr(s1.UpdatedAt, s2.UpdatedAt)
		case "-updated_at":
			return tools.CompareTimePtr(s2.UpdatedAt, s1.UpdatedAt)
		}
		return 0
	}
}

type StackRecord struct {
	GUID        string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	Name        string
	Description string
}

func NewStackRepository(
	builderNames []string,
	userClientFactory authorization.UserK8sClientFactory,
	rootNamespace string,
	sorter StackSorter,
) *StackRepository {
	return &StackRepository{
		builderNames:      builderNames,
		userClientFactory: userClientFactory,
		rootNamespace:     rootNamespace,
		sorter:            sorter,
	}
}

type ListStacksMessage struct {
	Names         []string
	LabelSelector string
	OrderBy       string
}

func (r *StackRepository) ListStacks(ctx context.Context, authInfo authorization.Info, message ListStacksMessage) ([]StackRecord, error) {
	var builderInfo korifiv1alpha1.BuilderInfo

	userClient, err := r.userClientFactory.BuildClient(authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to build user client: %w", err)
	}

	for _, b := range r.builderNames {
		err = userClient.Get(
			ctx,
			types.NamespacedName{
				Namespace: r.rootNamespace,
				Name:      b,
			},
			&builderInfo,
		)
		if err != nil {
			return nil, apierrors.FromK8sError(err, StackResourceType)
		}

		if !meta.IsStatusConditionTrue(builderInfo.Status.Conditions, korifiv1alpha1.StatusConditionReady) {
			return nil, apierrors.NewResourceNotReadyError(fmt.Errorf("BuilderInfo %q not ready", b))
		}
	}
	return r.sorter.Sort(builderInfoToStackRecords(builderInfo), message.OrderBy), nil
}

func builderInfoToStackRecords(info korifiv1alpha1.BuilderInfo) []StackRecord {
	return slices.Collect(it.Map(slices.Values(info.Status.Stacks), func(s korifiv1alpha1.BuilderInfoStatusStack) StackRecord {
		return StackRecord{
			Name:        s.Name,
			Description: s.Description,
			CreatedAt:   s.CreationTimestamp.Time,
			UpdatedAt:   &s.UpdatedTimestamp.Time,
		}
	}))
}
