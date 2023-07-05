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

	cmv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	cmeta "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	"github.com/google/uuid"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/net-certmanager/pkg/reconciler/certificate/config"
	"knative.dev/networking/pkg/apis/networking"
	"knative.dev/networking/pkg/apis/networking/v1alpha1"
	"knative.dev/pkg/apis"
	"knative.dev/pkg/kmeta"

	"github.com/martinlindhe/base36"
)

const (
	longest                               = 63
	base36Len                             = 25
	CreateCertManagerCertificateCondition = "CreateCertManagerCertificate"
	IssuerNotSetCondition                 = "IssuerNotSet"
	VisibilityClusterLocal                = "cluster-local"
)

// MakeCertManagerCertificate creates a Cert-Manager `Certificate` for requesting a SSL certificate.
func MakeCertManagerCertificate(cmConfig *config.CertManagerConfig, knCert *v1alpha1.Certificate) (*cmv1.Certificate, *apis.Condition) {
	var commonName string
	var dnsNames []string
	attemptedToShorten := false

	if len(knCert.Spec.DNSNames) > 0 {
		commonName = knCert.Spec.DNSNames[0]
	}

	// https://github.com/knative-sandbox/net-certmanager/issues/214
	// Only use the domain template if the commonName is too big.
	// This is to make the upgrade path easier and reduce churn on certificates.
	// The Route controller adds spec.domain to existing KCerts
	// The KCert controller requests new certs with same domain names, but a different CN if spec.domain is set and the other domain name would be too long
	// cert-manager Certificates are updated only if the existing domain name kept them from being issued.
	if len(commonName) > longest {
		if knCert.Spec.Domain != "" && knCert.Spec.Domain != commonName {
			//Split out the domain, and create a hash of the remaining part
			domainSuffix := "." + knCert.Spec.Domain
			prefix := strings.TrimSuffix(commonName, domainSuffix)
			if len(prefix) > base36Len {
				attemptedToShorten = true

				parsedUUID, err := uuid.Parse(string(knCert.UID))
				if err != nil {
					return nil, &apis.Condition{
						Type:   CreateCertManagerCertificateCondition,
						Status: corev1.ConditionFalse,
						Reason: "Failed To Parse UID",
						Message: fmt.Sprintf(
							"error creating cert-manager certificate: failed to parse UID (%s) on KCert (%s): %s",
							knCert.UID,
							knCert.Name,
							err,
						),
					}
				}
				parsedUUIDbytes := [16]byte(parsedUUID)
				prefix = strings.ToLower(base36.EncodeBytes(parsedUUIDbytes[:]))
			}
			commonName = prefix + domainSuffix

			//If the new name is still too long, then error
			if len(commonName) > longest {
				if attemptedToShorten {
					return nil, &apis.Condition{
						Type:   CreateCertManagerCertificateCondition,
						Status: corev1.ConditionFalse,
						Reason: "CommonName Too Long After Shortening",
						Message: fmt.Sprintf(
							"error creating cert-manager certificate: cannot create valid length CommonName: (%s) still longer than 63 characters after shortening",
							commonName,
						),
					}
				} else {
					return nil, &apis.Condition{
						Type:   CreateCertManagerCertificateCondition,
						Status: corev1.ConditionFalse,
						Reason: "CommonName Too Long",
						Message: fmt.Sprintf(
							"error creating cert-manager certificate: cannot create valid length CommonName: (%s) still longer than 63 characters, cannot shorten",
							commonName,
						),
					}
				}
			}
			dnsNames = append(dnsNames, commonName)
		} else {
			if knCert.Spec.Domain == commonName {
				return nil, &apis.Condition{
					Type:   CreateCertManagerCertificateCondition,
					Status: corev1.ConditionFalse,
					Reason: "DomainMapping Name Too Long",
					Message: fmt.Sprintf(
						"error creating cert-manager certificate: DomainMapping name (%s) longer than 63 characters",
						commonName,
					),
				}
			} else {
				return nil, &apis.Condition{
					Type:   CreateCertManagerCertificateCondition,
					Status: corev1.ConditionFalse,
					Reason: "CommonName Too Long",
					Message: fmt.Sprintf(
						"error creating cert-manager certificate: CommonName (%s) too long and no Domain available",
						commonName,
					),
				}
			}
		}

	}

	dnsNames = append(dnsNames, knCert.Spec.DNSNames...)

	var issuerRef cmeta.ObjectReference
	if knCert.Labels[networking.VisibilityLabelKey] == VisibilityClusterLocal {
		if cmConfig.ClusterInternalIssuerRef == nil {
			return nil, &apis.Condition{
				Type:    IssuerNotSetCondition,
				Status:  corev1.ConditionFalse,
				Reason:  "clusterInternalIssuerRef not set",
				Message: "error creating cert-manager certificate: clusterInternalIssuerRef was not set in config-certmanager",
			}
		}
		issuerRef = *cmConfig.ClusterInternalIssuerRef
	} else {
		if cmConfig.IssuerRef == nil {
			return nil, &apis.Condition{
				Type:    IssuerNotSetCondition,
				Status:  corev1.ConditionFalse,
				Reason:  "issuerRef not set",
				Message: "error creating cert-manager certificate: issuerRef was not set in config-certmanager",
			}
		}
		issuerRef = *cmConfig.IssuerRef
	}

	cert := &cmv1.Certificate{
		ObjectMeta: metav1.ObjectMeta{
			Name:            knCert.Name,
			Namespace:       knCert.Namespace,
			OwnerReferences: []metav1.OwnerReference{*kmeta.NewControllerRef(knCert)},
			Annotations:     knCert.GetAnnotations(),
			Labels:          knCert.GetLabels(),
		},
		Spec: cmv1.CertificateSpec{
			CommonName: commonName,
			SecretName: knCert.Spec.SecretName,
			DNSNames:   dnsNames,
			IssuerRef:  issuerRef,
			SecretTemplate: &cmv1.CertificateSecretTemplate{
				Labels: map[string]string{
					networking.CertificateUIDLabelKey: string(knCert.GetUID()),
				}},
		},
	}
	return cert, nil
}

// GetReadyCondition gets the ready condition of a Cert-Manager `Certificate`.
func GetReadyCondition(cmCert *cmv1.Certificate) *cmv1.CertificateCondition {
	for _, cond := range cmCert.Status.Conditions {
		if cond.Type == cmv1.CertificateConditionReady {
			return &cond
		}
	}
	return nil
}
