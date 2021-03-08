module knative.dev/net-certmanager

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.5.4
	github.com/jetstack/cert-manager v1.0.0
	go.uber.org/zap v1.16.0
	k8s.io/api v0.19.7
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/hack v0.0.0-20210305150220-f99a25560134
	knative.dev/networking v0.0.0-20210308052021-1b3d24e5be44
	knative.dev/pkg v0.0.0-20210308052421-737401c38b22
	sigs.k8s.io/structured-merge-diff/v3 v3.0.1-0.20200706213357-43c19bbb7fba // indirect
)

replace (
	github.com/go-logr/logr => github.com/go-logr/logr v0.1.0
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.1.0
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
	k8s.io/api => k8s.io/api v0.18.12
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.12
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.12
	k8s.io/apiserver => k8s.io/apiserver v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.12
	k8s.io/code-generator => k8s.io/code-generator v0.18.12
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.0.0
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
)
