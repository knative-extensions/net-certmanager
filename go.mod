module knative.dev/net-certmanager

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.5.6
	github.com/jetstack/cert-manager v1.3.1
	go.uber.org/zap v1.19.1
	k8s.io/api v0.21.4
	k8s.io/apimachinery v0.21.4
	k8s.io/client-go v0.21.4
	knative.dev/hack v0.0.0-20211122162614-813559cefdda
	knative.dev/networking v0.0.0-20211130182409-2731debe4463
	knative.dev/pkg v0.0.0-20211129195804-438776b3c87c
)
