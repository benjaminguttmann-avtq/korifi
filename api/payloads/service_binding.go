package payloads

import (
	"errors"
	"net/url"

	"code.cloudfoundry.org/korifi/api/payloads/parse"
	"code.cloudfoundry.org/korifi/api/payloads/validation"
	"code.cloudfoundry.org/korifi/api/repositories"
	jellidation "github.com/jellydator/validation"
)

type ServiceBindingCreate struct {
	Relationships *ServiceBindingRelationships `json:"relationships"`
	Type          string                       `json:"type"`
	Parameters    map[string]any               `json:"parameters"`
	Name          *string                      `json:"name"`
}

func (p ServiceBindingCreate) ToMessage(spaceGUID string) repositories.CreateServiceBindingMessage {
	var appGUID string
	if p.Relationships.App != nil {
		appGUID = p.Relationships.App.Data.GUID
	}

	return repositories.CreateServiceBindingMessage{
		Name:                p.Name,
		ServiceInstanceGUID: p.Relationships.ServiceInstance.Data.GUID,
		AppGUID:             appGUID,
		SpaceGUID:           spaceGUID,
		Parameters:          p.Parameters,
		Type:                p.Type,
	}
}

func (p ServiceBindingCreate) Validate() error {
	return jellidation.ValidateStruct(&p,
		jellidation.Field(&p.Type, validation.OneOf("app", "key")),
		jellidation.Field(&p.Name, jellidation.Required.When(p.Type == "key")),
		jellidation.Field(&p.Relationships, jellidation.Required),

		jellidation.Field(&p.Relationships, jellidation.By(func(value any) error {
			relationships, ok := value.(*ServiceBindingRelationships)
			if !ok || relationships == nil {
				return errors.New("relationships is required")
			}

			if p.Type == "app" {
				if relationships.App == nil {
					return jellidation.NewError("validation_required", "relationships.app is required")
				}
				if relationships.App.Data.GUID == "" {
					return jellidation.NewError("validation_required", "relationships.app.data.guid cannot be blank")
				}
			}

			return nil
		})),
	)
}

type ServiceBindingRelationships struct {
	App             *Relationship `json:"app"`
	ServiceInstance *Relationship `json:"service_instance"`
}

func (r ServiceBindingRelationships) Validate() error {
	return jellidation.ValidateStruct(&r,
		jellidation.Field(&r.ServiceInstance, jellidation.NotNil),
	)
}

type ServiceBindingList struct {
	Type                 string
	AppGUIDs             string
	ServiceInstanceGUIDs string
	Include              string
	LabelSelector        string
	PlanGUIDs            string
}

func (l ServiceBindingList) Validate() error {
	return jellidation.ValidateStruct(&l,
		jellidation.Field(&l.Type, validation.OneOf("app", "key")),
		jellidation.Field(&l.Include, validation.OneOf("app", "service_instance")),
	)
}

func (l *ServiceBindingList) ToMessage() repositories.ListServiceBindingsMessage {
	return repositories.ListServiceBindingsMessage{
		ServiceInstanceGUIDs: parse.ArrayParam(l.ServiceInstanceGUIDs),
		AppGUIDs:             parse.ArrayParam(l.AppGUIDs),
		LabelSelector:        l.LabelSelector,
		PlanGUIDs:            parse.ArrayParam(l.PlanGUIDs),
		Type:                 l.Type,
	}
}

func (l *ServiceBindingList) SupportedKeys() []string {
	return []string{"app_guids", "service_instance_guids", "include", "type", "per_page", "page", "label_selector", "service_plan_guids"}
}

func (l *ServiceBindingList) DecodeFromURLValues(values url.Values) error {
	l.Type = values.Get("type")
	l.AppGUIDs = values.Get("app_guids")
	l.ServiceInstanceGUIDs = values.Get("service_instance_guids")
	l.Include = values.Get("include")
	l.LabelSelector = values.Get("label_selector")
	l.PlanGUIDs = values.Get("service_plan_guids")
	return nil
}

type ServiceBindingUpdate struct {
	Metadata MetadataPatch `json:"metadata"`
}

func (u ServiceBindingUpdate) Validate() error {
	return jellidation.ValidateStruct(&u,
		jellidation.Field(&u.Metadata),
	)
}

func (c *ServiceBindingUpdate) ToMessage(serviceBindingGUID string) repositories.UpdateServiceBindingMessage {
	return repositories.UpdateServiceBindingMessage{
		GUID: serviceBindingGUID,
		MetadataPatch: repositories.MetadataPatch{
			Labels:      c.Metadata.Labels,
			Annotations: c.Metadata.Annotations,
		},
	}
}
