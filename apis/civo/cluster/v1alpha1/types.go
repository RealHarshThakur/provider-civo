/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/civo/civogo"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// CivoKubernetesParameters are the configurable fields of a CivoKubernetes.
type CivoKubernetesParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// CivoKubernetesObservation are the observable fields of a CivoKubernetes.
type CivoKubernetesObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// CivoKubernetesConnectionDetails is the desired output secret to store connection information
type CivoKubernetesConnectionDetails struct {
	ConnectionSecretNamePrefix string `json:"connectionSecretNamePrefix"`
	ConnectionSecretNamespace  string `json:"connectionSecretNamespace"`
}

// A CivoKubernetesSpec defines the desired state of a CivoKubernetes.
type CivoKubernetesSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	Name              string                               `json:"name"`
	Pools             []civogo.KubernetesClusterPoolConfig `json:"pools"`
	Applications      []string                             `json:"applications,omitempty"`
	ConnectionDetails CivoKubernetesConnectionDetails      `json:"connectionDetails"`
}

// A CivoKubernetesStatus represents the observed state of a CivoKubernetes.
type CivoKubernetesStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CivoKubernetesObservation `json:"atProvider,omitempty"`
	Message             string                    `json:"message"`
}

// +kubebuilder:object:root=true

// A CivoKubernetes is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="MESSAGE",type="string",JSONPath=".status.message"
// +kubebuilder:printcolumn:name="APPLICATIONS",type="string",JSONPath=".spec.applications"
// Please replace `PROVIDER-NAME` with your actual provider name, like `aws`, `azure`, `gcp`, `alibaba`
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,civo}
// +kubebuilder:subresource:status
type CivoKubernetes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CivoKubernetesSpec   `json:"spec"`
	Status CivoKubernetesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CivoKubernetesList contains a list of CivoKubernetes
type CivoKubernetesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CivoKubernetes `json:"items"`
}
