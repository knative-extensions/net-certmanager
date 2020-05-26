module knative.dev/net-certmanager

go 1.13

require (
	github.com/aws/aws-sdk-go v1.30.3 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.4.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.3 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jetstack/cert-manager v0.12.0
	github.com/json-iterator/go v1.1.9 // indirect
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59 // indirect
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
	istio.io/client-go v0.0.0-20200428154323-0ed2dc14724c // indirect
	k8s.io/api v0.18.1
	k8s.io/apiextensions-apiserver v0.18.0 // indirect
	k8s.io/apimachinery v0.18.1
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/kube-openapi v0.0.0-20200121204235-bf4fb3bd569c // indirect
	k8s.io/utils v0.0.0-20200327001022-6496210b90e8 // indirect
	knative.dev/pkg v0.0.0-20200519155757-14eb3ae3a5a7
	knative.dev/serving v0.15.0
	knative.dev/test-infra v0.0.0-20200519161858-554a95a37986
)

replace (
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2

	k8s.io/api => k8s.io/api v0.16.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.4
	k8s.io/client-go => k8s.io/client-go v0.16.4
	k8s.io/code-generator => k8s.io/code-generator v0.16.4
)
