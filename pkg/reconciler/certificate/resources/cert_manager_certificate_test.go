/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resources

import (
	"fmt"
	"strings"
	"testing"

	cmv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	cmmeta "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	"github.com/google/go-cmp/cmp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"knative.dev/net-certmanager/pkg/reconciler/certificate/config"
	"knative.dev/networking/pkg/apis/networking"
	"knative.dev/networking/pkg/apis/networking/v1alpha1"
	"knative.dev/pkg/kmeta"
)

const (
	servingRouteLabelKey     = "serving.knative.dev/route"
	servingCreatorAnnotation = "serving.knative.dev/creator"
	servingUpdaterAnnotation = "serving.knative.dev/lastModifier"
)

var cert = &v1alpha1.Certificate{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "test-cert",
		Namespace: "test-ns",
		Labels: map[string]string{
			servingRouteLabelKey: "test-route",
		},
		Annotations: map[string]string{
			servingCreatorAnnotation: "someone",
			servingUpdaterAnnotation: "someone",
		},
	},
	Spec: v1alpha1.CertificateSpec{
		DNSNames:   []string{"host1.example.com", "host2.example.com"},
		Domain:     "example.com",
		SecretName: "secret0",
	},
}

var (
	longHost         = "somebighost12345678910.somebignamespacename12345678910"
	domain           = "some.domain.test"
	longHostDNSNames = []string{longHost + "." + domain}
	certWithLongHost = &v1alpha1.Certificate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cert",
			Namespace: "test-ns",
			Labels: map[string]string{
				servingRouteLabelKey: "test-route",
			},
			Annotations: map[string]string{
				servingCreatorAnnotation: "someone",
				servingUpdaterAnnotation: "someone",
			},
		},
		Spec: v1alpha1.CertificateSpec{
			DNSNames:   longHostDNSNames,
			Domain:     domain,
			SecretName: "secret0",
		},
	}

	longDomain         = fmt.Sprintf("%s.%s", strings.Repeat("a", 54), "com")
	longDomainDNSNames = []string{"host1." + longDomain, "host2." + longDomain}
	certWithLongDomain = &v1alpha1.Certificate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cert",
			Namespace: "test-ns",
			Labels: map[string]string{
				servingRouteLabelKey: "test-route",
			},
			Annotations: map[string]string{
				servingCreatorAnnotation: "someone",
				servingUpdaterAnnotation: "someone",
			},
		},
		Spec: v1alpha1.CertificateSpec{
			DNSNames:   longDomainDNSNames,
			Domain:     longDomain,
			SecretName: "secret0",
		},
	}

	cmConfig = &config.CertManagerConfig{
		IssuerRef: &cmmeta.ObjectReference{
			Kind: "ClusterIssuer",
			Name: "Letsencrypt-issuer",
		},
		CommonNameTemplate: config.DefaultCommonNameTemplate,
	}
)

func TestMakeCertManagerCertificate(t *testing.T) {
	want := &cmv1.Certificate{
		ObjectMeta: metav1.ObjectMeta{
			Name:            "test-cert",
			Namespace:       "test-ns",
			OwnerReferences: []metav1.OwnerReference{*kmeta.NewControllerRef(cert)},
			Labels: map[string]string{
				servingRouteLabelKey: "test-route",
			},
			Annotations: map[string]string{
				servingCreatorAnnotation: "someone",
				servingUpdaterAnnotation: "someone",
			},
		},
		Spec: cmv1.CertificateSpec{
			SecretName: "secret0",
			CommonName: "host1.example.com",
			DNSNames:   []string{"host1.example.com", "host2.example.com"},
			IssuerRef: cmmeta.ObjectReference{
				Kind: "ClusterIssuer",
				Name: "Letsencrypt-issuer",
			},
			SecretTemplate: &cmv1.CertificateSecretTemplate{
				Labels: map[string]string{networking.CertificateUIDLabelKey: ""},
			},
		},
	}
	got, err := MakeCertManagerCertificate(cmConfig, cert)
	if err != nil {
		t.Errorf("MakeCertManagerCertificate Error: %s", err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("MakeCertManagerCertificate (-want, +got) = %s", diff)
	}
}

func TestMakeCertManagerCertificateLongCommonName(t *testing.T) {
	want := &cmv1.Certificate{
		ObjectMeta: metav1.ObjectMeta{
			Name:            "test-cert",
			Namespace:       "test-ns",
			OwnerReferences: []metav1.OwnerReference{*kmeta.NewControllerRef(cert)},
			Labels: map[string]string{
				servingRouteLabelKey: "test-route",
			},
			Annotations: map[string]string{
				servingCreatorAnnotation: "someone",
				servingUpdaterAnnotation: "someone",
			},
		},
		Spec: cmv1.CertificateSpec{
			SecretName: "secret0",
			CommonName: "3be3726dee8fe24a3dd3c59842f172e1.some.domain.test",
			DNSNames:   append([]string{"3be3726dee8fe24a3dd3c59842f172e1.some.domain.test"}, longHostDNSNames...),
			IssuerRef: cmmeta.ObjectReference{
				Kind: "ClusterIssuer",
				Name: "Letsencrypt-issuer",
			},
			SecretTemplate: &cmv1.CertificateSecretTemplate{
				Labels: map[string]string{networking.CertificateUIDLabelKey: ""},
			},
		},
	}
	got, err := MakeCertManagerCertificate(cmConfig, certWithLongHost)
	if err != nil {
		t.Errorf("MakeCertManagerCertificate Error: %s", err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("MakeCertManagerCertificate (-want, +got) = %s", diff)
	}
}

func TestMakeCertManagerCertificateDomainIsTooLong(t *testing.T) {
	wantError := fmt.Errorf("error creating Certmanager Certificate: cannot create valid length CommonName: Domain plus md5 hash of name+namespace data still longer than 63 characters.")
	cert, gotError := MakeCertManagerCertificate(cmConfig, certWithLongDomain)

	if cert != nil {
		t.Errorf("Expected no cert, got: %s", cmp.Diff(nil, cert))
	}

	if diff := cmp.Diff(wantError.Error(), gotError.Error()); diff != "" {
		t.Errorf("MakeCertManagerCertificate (-want, +got) = %s", diff)
	}
}

func TestGetReadyCondition(t *testing.T) {
	tests := []struct {
		name          string
		cmCertificate *cmv1.Certificate
		want          *cmv1.CertificateCondition
	}{{
		name:          "ready",
		cmCertificate: makeTestCertificate(cmmeta.ConditionTrue, cmv1.CertificateConditionReady, "ready", "ready"),
		want: &cmv1.CertificateCondition{
			Type:    cmv1.CertificateConditionReady,
			Status:  cmmeta.ConditionTrue,
			Reason:  "ready",
			Message: "ready",
		}}, {
		name:          "not ready",
		cmCertificate: makeTestCertificate(cmmeta.ConditionFalse, cmv1.CertificateConditionReady, "not ready", "not ready"),
		want: &cmv1.CertificateCondition{
			Type:    cmv1.CertificateConditionReady,
			Status:  cmmeta.ConditionFalse,
			Reason:  "not ready",
			Message: "not ready",
		}}, {
		name:          "unknow",
		cmCertificate: makeTestCertificate(cmmeta.ConditionUnknown, cmv1.CertificateConditionReady, "unknown", "unknown"),
		want: &cmv1.CertificateCondition{
			Type:    cmv1.CertificateConditionReady,
			Status:  cmmeta.ConditionUnknown,
			Reason:  "unknown",
			Message: "unknown",
		}}, {
		name:          "condition not ready",
		cmCertificate: makeTestCertificate(cmmeta.ConditionTrue, cmv1.CertificateConditionIssuing, "Renewing", "Renewing certificate"),
		want:          nil,
	},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := GetReadyCondition(test.cmCertificate)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("GetReadyCondition (-want, +got) = %s", diff)
			}
		})
	}
}

func makeTestCertificate(condStatus cmmeta.ConditionStatus, condType cmv1.CertificateConditionType, reason, message string) *cmv1.Certificate {
	cert := &cmv1.Certificate{
		Status: cmv1.CertificateStatus{
			Conditions: []cmv1.CertificateCondition{{
				Type:    condType,
				Status:  condStatus,
				Reason:  reason,
				Message: message,
			}},
		},
	}
	return cert
}
