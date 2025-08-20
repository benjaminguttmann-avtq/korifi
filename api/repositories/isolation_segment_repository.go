package repositories

import (
	"context"
	"time"

	//	"fmt"
	//	"maps"
	//	"slices"
	"code.cloudfoundry.org/korifi/api/authorization"
	//	apierrors "code.cloudfoundry.org/korifi/api/errors"
	//	"code.cloudfoundry.org/korifi/api/repositories/k8sklient/descriptors"
	//	korifiv1alpha1 "code.cloudfoundry.org/korifi/controllers/api/v1alpha1"
	//	"code.cloudfoundry.org/korifi/tools"
	//
	//	"github.com/BooleanCat/go-functional/v2/it"
	//	"github.com/google/uuid"
	//	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IsolationSegmentRecord struct {
	Name      string
	GUID      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (r *IsolationSegmentRepo) ListIsolationSegments(ctx context.Context, authInfo authorization.Info) (ListResult[IsolationSegmentRecord], error) {
	return ListResult[IsolationSegmentRecord]{}, nil
}

type IsolationSegmentRepo struct {
	klient Klient
}
