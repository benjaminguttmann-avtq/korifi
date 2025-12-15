package repositories

import (
	"context"
	"slices"
	"time"

	"code.cloudfoundry.org/korifi/api/authorization"
	apierrors "code.cloudfoundry.org/korifi/api/errors"
	korifiv1alpha1 "code.cloudfoundry.org/korifi/controllers/api/v1alpha1"
	"code.cloudfoundry.org/korifi/controllers/webhooks/validation"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const SpaceQuotaResourceType = "Space Quota"

type Apps struct {
	PerProcessInMemoryInMB       int
	TotalMemoryInMB              int
	TotalInstances               int
	LogRateLimitInBytesPerSecond int
	PerAppTasks                  int
}

type Services struct {
	PaidServicesAllowed   bool
	TotalServiceInstances int
	TotalServiceKeys      int
}

type Routes struct {
	TotalRoutes        int
	TotalReservedPorts int
}

type SpaceQuotaRecord struct {
	GUID      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Name      string
	Apps      Apps
	Services  Services
	Routes    Routes
}
