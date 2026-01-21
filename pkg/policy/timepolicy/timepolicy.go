package timepolicy

import (
	"github.com/olisajc/appScaler/pkg/policy/cpuconfig"
	"github.com/olisajc/appScaler/pkg/policy/memconfig"
	"github.com/olisajc/appScaler/pkg/policy/storageconfig"
)

type TimePolicy struct {
	Name     string                       `json:"name,omitempty"`
	Schedule string                       `json:"schedule,omitempty"`
	Replicas *int32                       `json:"replicas,omitempty"`
	CPU      *cpuconfig.CpuConfig         `json:"cpu,omitempty"`
	Memory   *memconfig.MemConfig         `json:"memory,omitempty"`
	Storage  *storageconfig.StorageConfig `json:"storage,omitempty"`
}

type TimePolicies []TimePolicy
