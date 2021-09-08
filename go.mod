module knative.dev/net-certmanager

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.5.6
	github.com/jetstack/cert-manager v1.3.1
	go.uber.org/zap v1.19.0
	k8s.io/api v0.21.4
	k8s.io/apimachinery v0.21.4
	k8s.io/client-go v0.21.4
	knative.dev/hack v0.0.0-20210806075220-815cd312d65c
	knative.dev/networking v0.0.0-20210908132645-c94e114d7fed
	knative.dev/pkg v0.0.0-20210908025933-71508fc69a57
)
