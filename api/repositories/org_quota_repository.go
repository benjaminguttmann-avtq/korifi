package repositories

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"time"

	"code.cloudfoundry.org/korifi/api/authorization"
	apierrors "code.cloudfoundry.org/korifi/api/errors"
	korifiv1alpha1 "code.cloudfoundry.org/korifi/controllers/api/v1alpha1"
	"code.cloudfoundry.org/korifi/tools"

	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OrgQuotaRepo struct {
	klient Klient
}

func NewOrgQuotaRepo(
	klient Klient,
) *OrgRepo {
	return &OrgQuotaRepo{
		klient: klient,
	}
}
