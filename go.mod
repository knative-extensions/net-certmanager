module knative.dev/net-certmanager

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.5.2
	github.com/jetstack/cert-manager v0.12.0
	go.uber.org/zap v1.16.0
	k8s.io/api v0.18.12
	k8s.io/apimachinery v0.18.12
	k8s.io/client-go v0.18.12
	knative.dev/networking v0.0.0-20201118013152-4fcad21135a2
	knative.dev/pkg v0.0.0-20201117221452-0fccc54273ed
	knative.dev/test-infra/scripts v0.0.0-20200610004422-8b4a5283a123
)
