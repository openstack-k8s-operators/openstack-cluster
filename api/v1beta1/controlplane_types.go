/*


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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KeystoneSpec defines the desired state of KeystoneAPI
type KeystoneSpec struct {
	// number of Keystone API replicas
	Replicas int `json:"replicas,omitempty"`
}

// GlanceSpec defines the desired state of GlanceAPI
type GlanceSpec struct {
	// number of Glance API replicas
	Replicas int `json:"replicas,omitempty"`
}

// PlacementSpec defines the desired state of PlacementAPI
type PlacementSpec struct {
	// number of Placement API replicas
	Replicas int `json:"replicas,omitempty"`
}

// InterconnectSpec defines the desired state of Interconnect
type InterconnectSpec struct {
	// number of Interconnect
	Replicas int `json:"replicas,omitempty"`
}

// NovaSpec defines the desired state of Nova Control Plane
type NovaSpec struct {
	// number of Nova API replicas
	NovaAPIReplicas int `json:"novaAPIReplicas,omitempty"`
	// number of Nova Scheduler replicas
	NovaSchedulerReplicas int `json:"novaSchedulerReplicas,omitempty"`
	// number of Nova Conductor replicas
	NovaConductorReplicas int `json:"novaConductorReplicas,omitempty"`
	// number of Nova Metadata replicas
	NovaMetadataReplicas int `json:"novaMetadataReplicas,omitempty"`
	// number of Nova NoVNCProxy replicas
	NovaNoVNCProxyReplicas int `json:"novaNoVNCProxyReplicas,omitempty"`
}

// CinderSpec defines the desired state of Cinder Control Plane
type CinderSpec struct {
	// number of Cinder API replicas
	CinderAPIReplicas int `json:"cinderAPIReplicas,omitempty"`
	// number of Cinder Scheduler replicas
	CinderSchedulerReplicas int `json:"cinderSchedulerReplicas,omitempty"`
	// number of Cinder Backup replicas
	CinderBackupReplicas int `json:"cinderBackupReplicas,omitempty"`
	// number of Cinder Volume replicas
	// Todo: how to handle different cinder volume services
	CinderVolumeReplicas int `json:"cinderVolumeReplicas,omitempty"`
}

// NeutronSpec defines the desired state of NeutronAPI
type NeutronSpec struct {
	// number of Neutron API replicas
	Replicas int `json:"replicas,omitempty"`
}

// ControlPlaneSpec defines the desired state of ControlPlane
type ControlPlaneSpec struct {
	// storage class to use for storage claims
	StorageClass string `json:"storageClass,omitempty"`
	// Keystone API settings
	Keystone KeystoneSpec `json:"keystone,omitempty"`
	// Glance API settings
	Glance GlanceSpec `json:"glance,omitempty"`
	// Placement API settings
	Placement PlacementSpec `json:"placement,omitempty"`
	// AMQ Interconnect settings
	Interconnect InterconnectSpec `json:"interconnect,omitempty"`
	// Nova settings
	Nova NovaSpec `json:"nova,omitempty"`
	// Cinder settings
	Cinder CinderSpec `json:"cinder,omitempty"`
	// Neutron settings
	Neutron NeutronSpec `json:"neutron,omitempty"`
}

// ControlPlaneStatus defines the observed state of ControlPlane
type ControlPlaneStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ControlPlane is the Schema for the controlplanes API
type ControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ControlPlaneSpec   `json:"spec,omitempty"`
	Status ControlPlaneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ControlPlaneList contains a list of ControlPlane
type ControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControlPlane `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ControlPlane{}, &ControlPlaneList{})
}
