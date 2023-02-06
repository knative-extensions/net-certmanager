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
	"crypto/md5"
	"fmt"
	"strings"

	cmv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/net-certmanager/pkg/reconciler/certificate/config"
	"knative.dev/networking/pkg/apis/networking"
	"knative.dev/networking/pkg/apis/networking/v1alpha1"
	"knative.dev/pkg/kmeta"
)

const (
	longest = 63
	md5Len  = 32
)

// MakeCertManagerCertificate creates a Cert-Manager `Certificate` for requesting a SSL certificate.
func MakeCertManagerCertificate(cmConfig *config.CertManagerConfig, knCert *v1alpha1.Certificate) (*cmv1.Certificate, error) {
	var commonName string
	var dnsNames []string

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
		if knCert.Spec.Domain != "" {
			//Split out the domain, and create a hash of the remaining part
			domainSuffix := "." + knCert.Spec.Domain
			prefix := strings.SplitN(commonName, domainSuffix, 2)[0]
			if len(prefix) > md5Len {
				prefix = fmt.Sprintf("%x", md5.Sum([]byte(prefix)))
			}
			commonName = prefix + domainSuffix

			//If the new name is still too long, then error
			if len(commonName) > longest {
				return nil, fmt.Errorf("error creating Certmanager Certificate: cannot create valid length CommonName: %s", "Domain plus md5 hash of name+namespace data still longer than 63 characters.")
			}
			dnsNames = append(dnsNames, commonName)
		} else {
			return nil, fmt.Errorf("error creating Certmanager Certificate: %s", "commonName too long and no Domain available")
		}

	}

	dnsNames = append(dnsNames, knCert.Spec.DNSNames...)

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
			IssuerRef:  *cmConfig.IssuerRef,
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
