module knative.dev/net-certmanager

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.4.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.3 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jetstack/cert-manager v0.12.0
	github.com/json-iterator/go v1.1.9 // indirect
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59 // indirect
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
	k8s.io/api v0.18.1
	k8s.io/apimachinery v0.18.1
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/utils v0.0.0-20200327001022-6496210b90e8 // indirect
	knative.dev/networking v0.0.0-20200630191330-5080f859c17d
	knative.dev/pkg v0.0.0-20200630170034-2c1a029eb97f
	knative.dev/serving v0.15.1-0.20200630173329-4066b8c74b37
	knative.dev/test-infra v0.0.0-20200630141629-15f40fe97047
)

replace (
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2

	k8s.io/api => k8s.io/api v0.17.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6
	k8s.io/client-go => k8s.io/client-go v0.17.6
	k8s.io/code-generator => k8s.io/code-generator v0.17.6
)
