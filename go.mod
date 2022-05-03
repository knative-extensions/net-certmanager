module knative.dev/net-certmanager

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.5.6
	github.com/jetstack/cert-manager v1.3.1
	go.uber.org/zap v1.19.1
	k8s.io/api v0.23.5
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.23.5
	knative.dev/hack v0.0.0-20220427014036-5f473869d377
	knative.dev/networking v0.0.0-20220429044653-591d2bb63aae
	knative.dev/pkg v0.0.0-20220502225657-4fced0164c9a
)
