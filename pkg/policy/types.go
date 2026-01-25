package policy

import (
	"github.com/olisajc/appScaler/pkg/policy/overridepolicy"
	"github.com/olisajc/appScaler/pkg/policy/timepolicy"
)

type Policy struct {
	Name     string                         `json:"name,omitempty"`
	Time     timepolicy.TimePolicies        `json:"time,omitempty"`
	Override *overridepolicy.OverridePolicy `json:"override,omitempty"`
}
