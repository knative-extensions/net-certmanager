module knative.dev/net-certmanager

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.5.6
	github.com/jetstack/cert-manager v1.3.1
	go.uber.org/zap v1.18.1
	k8s.io/api v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v0.20.7
	knative.dev/hack v0.0.0-20210806075220-815cd312d65c
	knative.dev/networking v0.0.0-20210824140523-51512a042e23
	knative.dev/pkg v0.0.0-20210825070025-a70bb26767b8
)
