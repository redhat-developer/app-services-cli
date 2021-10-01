package cluster

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO this is not needed.

// CustomResourceOptions object contains the data required to create a custom connection
type CustomResourceOptions struct {
	Path        string
	CRName      string
	ServiceName string
	CRJSON      []byte
	Resource    schema.GroupVersionResource
}
