package timepolicy

import "github.com/olisajc/appScaler/pkg/policy/resourcetargets"

type TimePolicy struct {
	Name     string `json:"name,omitempty"`
	Schedule string `json:"schedule,omitempty"`
	*resourcetargets.ResourceTargets
}

type TimePolicies []TimePolicy
