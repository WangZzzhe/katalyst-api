/*
Copyright 2022 The Katalyst Authors.

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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=kcc
// +kubebuilder:printcolumn:name="GROUP",type="string",JSONPath=".spec.targetType.group"
// +kubebuilder:printcolumn:name="RESOURCE",type="string",JSONPath=".spec.targetType.resource"
// +kubebuilder:printcolumn:name="VERSION",type="string",JSONPath=".spec.targetType.version"
// +kubebuilder:printcolumn:name="GENERATION",type="integer",JSONPath=".metadata.generation"
// +kubebuilder:printcolumn:name="OBSERVED",type="integer",JSONPath=".status.observedGeneration"
// +kubebuilder:printcolumn:name="INVALID",type="string",JSONPath=".status.invalidTargetConfigList"
// +kubebuilder:printcolumn:name="VALID",type="string",JSONPath=".status.conditions[?(@.type==\"Valid\")].status"
// +kubebuilder:printcolumn:name="REASON",type=string,JSONPath=".status.conditions[?(@.type==\"Valid\")].reason"
// +kubebuilder:printcolumn:name="MESSAGE",type=string,JSONPath=".status.conditions[?(@.type==\"Valid\")].message"

// HaloCustomConfig is the Schema for the custom configuration API in katalyst
type HaloCustomConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HaloCustomConfigSpec   `json:"spec,omitempty"`
	Status HaloCustomConfigStatus `json:"status,omitempty"`
}

// HaloCustomConfigSpec defines the desired state of HaloCustomConfig
type HaloCustomConfigSpec struct {
	// the GVR of target config type
	TargetType metav1.GroupVersionResource `json:"targetType"`
	// whether disable revisionHistory for the HaloCustomConfig resource
	// +kubebuilder:default:=true
	// +optional
	DisableRevisionHistory bool `json:"disableRevisionHistory,omitempty"`
	// the keys list allowed in node selector to select which nodes will be effected by the HaloCustomConfig resource,
	// and the priority will be used when one node match two HaloCustomConfig resource at the same time, the higher
	// priority one will be considered. If not set, node label selector is not allowed to use.
	// +patchMergeKey=priority
	// +patchStrategy=merge
	// +optional
	NodeLabelSelectorAllowedKeyList []PriorityNodeLabelSelectorAllowedKeyList `json:"nodeLabelSelectorAllowedKeyList,omitempty"`
}

// PriorityNodeLabelSelectorAllowedKeyList defines the priority and its allowed key list
type PriorityNodeLabelSelectorAllowedKeyList struct {
	// Priority is the priority of configurations
	Priority int32 `json:"priority"`
	// KeyList is allowed to use in node selector in the Priority
	KeyList []string `json:"keyList"`
}

// HaloCustomConfigStatus defines the observed state of HaloCustomConfig
type HaloCustomConfigStatus struct {
	// ObservedGeneration is the generation as observed by the controller consuming the HaloCustomConfig.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// InvalidTargetConfigList is a name list of invalid Config
	InvalidTargetConfigList []string `json:"invalidTargetConfigList,omitempty"`

	// Represents the latest available observations of a HaloCustomConfig's current state.
	Conditions []HaloCustomConfigCondition `json:"conditions,omitempty"`
}

type HaloCustomConfigCondition struct {
	// Type of HaloCustomConfig condition
	Type HaloCustomConfigConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// reason is the reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// message is a human-readable explanation containing details about the transition
	// +optional
	Message string `json:"message,omitempty"`
}

type HaloCustomConfigConditionType string

const (
	HaloCustomConfigConditionTypeValid HaloCustomConfigConditionType = "Valid"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HaloCustomConfigList contains a list of HaloCustomConfig
type HaloCustomConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HaloCustomConfig `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=cnc,scope=Cluster

// CustomNodeConfig is the Schema for the customnodeconfigs API
type CustomNodeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomNodeConfigSpec   `json:"spec,omitempty"`
	Status CustomNodeConfigStatus `json:"status,omitempty"`
}

type CustomNodeConfigSpec struct {
}

// CustomNodeConfigStatus defines the desired state of HaloCustomConfig
type CustomNodeConfigStatus struct {
	// HaloCustomConfigList contains a list of target custom config
	HaloCustomConfigList []TargetConfig `json:"haloCustomConfigList,omitempty"`

	// ServiceProfileConfigList contains a list of target service Profile config
	ServiceProfileConfigList []TargetConfig `json:"serviceProfileConfigList,omitempty"`
}

// TargetConfig current hash for specific gvk config object
type TargetConfig struct {
	// ConfigType gvr of target config
	ConfigType metav1.GroupVersionResource `json:"configType"`
	// ConfigName name of target config
	ConfigName string `json:"configName"`
	// ConfigNamespace namespace of target config
	ConfigNamespace string `json:"configNamespace"`
	// Hash is current hash value of target config.
	// The agent will first check whether the local config hash and
	// the target config hash are equal, only if not, it will try to
	// update the local config from the remote.
	Hash string `json:"hash"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomNodeConfigList contains a list of CustomNodeConfig
type CustomNodeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomNodeConfig `json:"items"`
}
