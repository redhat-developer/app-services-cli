module github.com/bf2fc6cc711aee1a0c2a/cli/operator

go 1.13

replace github.com/bf2fc6cc711aee1a0c2a/cli/client/mas => ./client/mas

require (
	github.com/go-logr/logr v0.1.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	sigs.k8s.io/controller-runtime v0.6.2
)
