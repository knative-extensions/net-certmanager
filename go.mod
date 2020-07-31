module knative.dev/net-certmanager

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.5.1
	github.com/jetstack/cert-manager v0.12.0
	github.com/json-iterator/go v1.1.9 // indirect
	go.uber.org/zap v1.14.1
	k8s.io/api v0.18.1
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/utils v0.0.0-20200327001022-6496210b90e8 // indirect
	knative.dev/networking v0.0.0-20200731163501-cc631d26a89e
	knative.dev/pkg v0.0.0-20200731005101-694087017879
	knative.dev/serving v0.16.1-0.20200731162400-c58ef770af82
	knative.dev/test-infra v0.0.0-20200731141600-8bb2015c65e2
)

replace (
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2

	k8s.io/api => k8s.io/api v0.17.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6
	k8s.io/client-go => k8s.io/client-go v0.17.6
	k8s.io/code-generator => k8s.io/code-generator v0.17.6
)
