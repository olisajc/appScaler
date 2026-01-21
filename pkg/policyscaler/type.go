package policyscaler

import (
	"github.com/olisajc/appScaler/pkg/policy"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PolicyScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyScalerSpec `json:"spec,omitempty"`
}

type PolicyScalerSpec struct {
	Policies policy.Policy `json:"policies,omitempty"`
}

type PolicyScalerList []*PolicyScaler
