/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"github.com/ghodss/yaml"

	cmv1alpha2 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	cmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	issuerRefKey    = "issuerRef"
	keyAlgorithmKey = "keyAlgorithm"
	keySizeKey      = "keySize"

	// CertManagerConfigName is the name of the configmap containing all
	// configuration related to Cert-Manager.
	CertManagerConfigName = "config-certmanager"
)

// CertManagerConfig contains Cert-Manager related configuration defined in the
// `config-certmanager` config map.
type CertManagerConfig struct {
	IssuerRef    *cmeta.ObjectReference
	KeyAlgorithm cmv1alpha2.KeyAlgorithm
	KeySize      int
}

// NewCertManagerConfigFromConfigMap creates an CertManagerConfig from the supplied ConfigMap
func NewCertManagerConfigFromConfigMap(configMap *corev1.ConfigMap) (*CertManagerConfig, error) {
	// TODO(zhiminx): do we need to provide the default values here?
	// TODO: validation check.

	config := &CertManagerConfig{
		IssuerRef: &cmeta.ObjectReference{},
	}

	if v, ok := configMap.Data[issuerRefKey]; ok {
		if err := yaml.Unmarshal([]byte(v), config.IssuerRef); err != nil {
			return nil, err
		}
	}

	if v, ok := configMap.Data[keyAlgorithmKey]; ok {
		if err := yaml.Unmarshal([]byte(v), &config.KeyAlgorithm); err != nil {
			return nil, err
		}
	}

	if v, ok := configMap.Data[keySizeKey]; ok {
		if err := yaml.Unmarshal([]byte(v), &config.KeySize); err != nil {
			return nil, err
		}
	}
	return config, nil
}
