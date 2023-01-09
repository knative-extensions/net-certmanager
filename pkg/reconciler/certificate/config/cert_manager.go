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
	"bytes"
	"fmt"
	"text/template"

	"github.com/ghodss/yaml"

	cmeta "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/utils/lru"
)

const (
	issuerRefKey          = "issuerRef"
	commonNameTemplateKey = "commonNameTemplate"

	// CertManagerConfigName is the name of the configmap containing all
	// configuration related to Cert-Manager.
	CertManagerConfigName = "config-certmanager"

	DefaultCommonNameTemplate = "k.{{.Domain}}"
)

var (
	// Store ~10 latest templates per template type.
	templateCache *lru.Cache
	// Verify the default templates are valid.
	_ = template.Must(template.New("common-name-template").Parse(DefaultCommonNameTemplate))
)

func init() {
	// The only failure is due to negative size.
	// Store ~10 latest templates per template type.
	templateCache = lru.New(10 * 2)
}

// CommonNametemplateValues are the available properties people can choose from
// in their Route's "CommonNameTemplate" golang template sting.
type CommonNameTemplateValues struct {
	Domain      string
	Annotations map[string]string
	Labels      map[string]string
}

// CertManagerConfig contains Cert-Manager related configuration defined in the
// `config-certmanager` config map.
type CertManagerConfig struct {
	IssuerRef *cmeta.ObjectReference

	//CommonNameTemplate is the gloang text template to use to generate the
	//CommonName for a Route's Certificate
	//Optional
	CommonNameTemplate string
}

// NewCertManagerConfigFromConfigMap creates an CertManagerConfig from the supplied ConfigMap
func NewCertManagerConfigFromConfigMap(configMap *corev1.ConfigMap) (*CertManagerConfig, error) {
	// TODO(zhiminx): do we need to provide the default values here?
	// TODO: validation check.

	config := &CertManagerConfig{
		IssuerRef:          &cmeta.ObjectReference{},
		CommonNameTemplate: DefaultCommonNameTemplate,
	}

	if v, ok := configMap.Data[issuerRefKey]; ok {
		if err := yaml.Unmarshal([]byte(v), config.IssuerRef); err != nil {
			return nil, err
		}
	}

	if v, ok := configMap.Data[commonNameTemplateKey]; ok && v != "" {
		// Verify common-name-template and add to the cache.
		t, err := template.New("common-name-template").Parse(v)
		if err != nil {
			return nil, err
		}
		if err := checkDomainTemplate(t); err != nil {
			return nil, err
		}
		templateCache.Add(config.CommonNameTemplate, t)
		config.CommonNameTemplate = v
	}

	return config, nil
}

// GetDomainTemplate returns the golang Template from the config map
// or panics (the value is validated during CM validation and at
// this point guaranteed to be parseable).
func (c *CertManagerConfig) GetCommonNameTemplate() *template.Template {
	if tt, ok := templateCache.Get(c.CommonNameTemplate); ok {
		return tt.(*template.Template)
	}
	fmt.Println("no template found")
	// Should not really happen outside of route/ingress unit tests.
	nt := template.Must(template.New("common-name-template").Parse(
		c.CommonNameTemplate))
	templateCache.Add(c.CommonNameTemplate, nt)
	return nt
}

func checkDomainTemplate(t *template.Template) error {
	// To a test run of applying the template, and see if the
	// result is a valid URL.
	data := CommonNameTemplateValues{
		Domain:      "baz.com",
		Annotations: nil,
		Labels:      nil,
	}
	buf := bytes.Buffer{}
	if err := t.Execute(&buf, data); err != nil {
		return err
	}
	if errs := validation.IsDNS1123Subdomain(buf.String()); len(errs) != 0 {
		return fmt.Errorf(fmt.Sprintf("Invalid DNS1123 Subdomain %63s\n\n Errors: %v", buf.String(), errs))
	}

	return nil
}
