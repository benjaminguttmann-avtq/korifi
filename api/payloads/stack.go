package payloads

import (
	"net/url"

	"code.cloudfoundry.org/korifi/api/payloads/parse"
	"code.cloudfoundry.org/korifi/api/payloads/validation"
	"code.cloudfoundry.org/korifi/api/repositories"
	jellidation "github.com/jellydator/validation"
)

type StackList struct {
	Names         string
	OrderBy       string
	LabelSelector string
}

func (s StackList) Validate() error {
	return jellidation.ValidateStruct(&s,
		jellidation.Field(&s.OrderBy, validation.OneOfOrderBy("created_at", "updated_at", "name")),
	)
}

func (s *StackList) ToMessage() repositories.ListStacksMessage {
	return repositories.ListStacksMessage{
		Names:         parse.ArrayParam(s.Names),
		LabelSelector: s.LabelSelector,
		OrderBy:       s.OrderBy,
	}
}

func (s *StackList) SupportedKeys() []string {
	return []string{"names", "order_by", "per_page", "page", "label_selector"}
}

func (s *StackList) DecodeFromURLValues(values url.Values) error {
	s.Names = values.Get("names")
	s.OrderBy = values.Get("order_by")
	s.LabelSelector = values.Get("label_selector")
	return nil
}
