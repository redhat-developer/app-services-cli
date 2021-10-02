package utils

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
	"github.com/redhat-developer/service-binding-operator/apis/binding/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
)

// TODO This should be gone/ Not needed.

// UseOperatorForBinding performs binding using ServiceBinding operator
func UseOperatorForBinding(ctx context.Context, opts v1alpha.CommandEnvironment, sb *v1alpha1.ServiceBinding, dynamicClient dynamic.Interface, ns string) error {
	opts.Logger.Info(opts.Localizer.MustLocalize("cluster.serviceBinding.usingOperator"))
	sbData, err := runtime.DefaultUnstructuredConverter.ToUnstructured(sb)
	if err != nil {
		return err
	}

	unstructuredSB := unstructured.Unstructured{Object: sbData}
	_, err = dynamicClient.Resource(v1alpha1.GroupVersionResource).Namespace(ns).
		Create(ctx, &unstructuredSB, metav1.CreateOptions{})

	return err
}

// CheckIfOperatorIsInstalled checks if ServiceBindingOperator is installed on the cluster
func CheckIfOperatorIsInstalled(ctx context.Context, dynamicClient dynamic.Interface, ns string) error {
	_, err := dynamicClient.Resource(v1alpha1.GroupVersionResource).Namespace(ns).
		List(ctx, metav1.ListOptions{Limit: 1})
	return err
}
