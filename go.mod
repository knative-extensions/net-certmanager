module knative.dev/net-certmanager

go 1.13

require (
	contrib.go.opencensus.io/exporter/ocagent v0.6.0 // indirect
	contrib.go.opencensus.io/exporter/prometheus v0.1.0 // indirect
	contrib.go.opencensus.io/exporter/stackdriver v0.12.9-0.20191108183826-59d068f8d8ff // indirect
	github.com/aws/aws-sdk-go v1.30.3 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/go-cmp v0.3.1
	github.com/google/go-containerregistry v0.0.0-20200331213917-3d03ed9b1ca2 // indirect
	github.com/googleapis/gnostic v0.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.3 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jetstack/cert-manager v0.12.0
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/openzipkin/zipkin-go v0.2.2 // indirect
	github.com/prometheus/client_golang v1.1.0 // indirect
	github.com/prometheus/common v0.9.1 // indirect
	github.com/prometheus/procfs v0.0.11 // indirect
	go.opencensus.io v0.22.3 // indirect
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59 // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/api v0.13.1-0.20191104181305-997ecb25550d // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	istio.io/client-go v0.0.0-20200428154323-0ed2dc14724c // indirect
	k8s.io/api v0.18.1
	k8s.io/apiextensions-apiserver v0.18.0 // indirect
	k8s.io/apimachinery v0.18.1
	k8s.io/client-go v0.18.1
	k8s.io/code-generator v0.18.0 // indirect
	k8s.io/kube-openapi v0.0.0-20200121204235-bf4fb3bd569c // indirect
	k8s.io/utils v0.0.0-20200327001022-6496210b90e8 // indirect
	knative.dev/pkg v0.0.0-20200427220451-5ba6c1d101a0
	knative.dev/serving v0.14.1-0.20200426043050-7ad5cc721f86
	knative.dev/test-infra v0.0.0-20200317133231-01c075fbeae4
)

replace (
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2

	k8s.io/api => k8s.io/api v0.16.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.4
	k8s.io/client-go => k8s.io/client-go v0.16.4
	k8s.io/code-generator => k8s.io/code-generator v0.16.4
)
