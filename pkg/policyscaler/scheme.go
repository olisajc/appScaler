package policyscaler

import "k8s.io/apimachinery/pkg/runtime/schema"

var policyScalerSchemaV1 = schema.GroupVersionResource{
	Group:    "extensions.example.com",
	Version:  "v1",
	Resource: "policyscalers",
}

func GetPolicyScalerSchema() schema.GroupVersionResource {
	return policyScalerSchemaV1
}
