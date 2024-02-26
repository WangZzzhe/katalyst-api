package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// ResourcePortraitSpec defines the desired state of ResourcePortrait
type ResourcePortraitSpec struct {
	Target `json:"target"`

	Metrics       []MetricSpec `json:"metrics,omitempty"`
	CustomMetrics []MetricSpec `json:"customMetrics,omitempty"`

	Algorithms []Algorithm `json:"algorithms"`
}

type TargetType string

const (
	ClusterTargetType   TargetType = "cluster"
	NamespaceTargetType TargetType = "namespace"
	NodeTargetType      TargetType = "node"
	WorkloadTargetType  TargetType = "workload"
)

const (
	TagTargetType      = "target_type"
	TagTargetName      = "target_name"
	TagTargetNamespace = "namespace"

	TagMethodIdentify = "method_identify"
)

type OperatorType string

const (
	OperatorTypeAvg   OperatorType = "avg"
	OperatorTypeMax   OperatorType = "max"
	OperatorTypeMin   OperatorType = "min"
	OperatorTypeLast  OperatorType = "last"
	OperatorTypeFirst OperatorType = "first"
)

type TimeWindow struct {
	Input    int          `json:"input,omitempty"`
	Output   int          `json:"output,omitempty"`
	Operator OperatorType `json:"operator,omitempty"`
}

type AuthType string

const (
	AuthTypeTokenKey         = "AuthToken"
	AuthTypeBasicUsernameKey = "Username"
	AuthTypeBasicPasswordKey = "Password"
	AuthTypeAKSKAccessKey    = "AccessKey"
	AuthTypeAKSKSecretKey    = "SecretKey"
)

type Target struct {
	Type TargetType `json:"type"`
	Name string     `json:"name"`
}

type MetricSpec struct {
	TimeWindow time.Duration `json:"timeWindow,omitempty"`

	Metric string `json:"metric"`
	PromQL string `json:"promQL,omitempty"`

	Source *MetricSource `json:"sourceAddress,omitempty"`
}

type MetricSource struct {
	Address string `json:"address"`

	AuthType  string `json:"authType,omitempty"`
	AuthToken string `json:"authToken,omitempty"`

	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Algorithm struct {
	Method       string            `json:"method"`
	Version      string            `json:"version,omitempty"`
	CallInterval time.Duration     `json:"callInterval"`
	Params       map[string]string `json:"params,omitempty"`
}

type AlgorithmRecord struct {
	Algorithm    Algorithm `json:"algorithm"`
	LastCallTime int64     `json:"lastCallTime"`
}

// ResourcePortraitStatus defines the observed state of ResourcePortrait
type ResourcePortraitStatus struct {
	CallTimeRecord []AlgorithmRecord `json:"callTimeRecord,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=rpt

// ResourcePortrait is the Schema for the resourceportraits API
type ResourcePortrait struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourcePortraitSpec   `json:"spec,omitempty"`
	Status ResourcePortraitStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// ResourcePortraitList contains a list of ResourcePortrait
type ResourcePortraitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcePortrait `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=rpr

// Result is the Schema for the results API
type Portrait struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	TimeSeries map[string][]Item `json:"timeSeries,omitempty"`
	Data       map[string]string `json:"data,omitempty"`
}

type Item struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`

	LabelKeys   []int32   `json:"labelKeys,omitempty"`
	LabelValues []float64 `json:"labelValues,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// PortraitList contains a list of Result
type PortraitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Portrait `json:"items"`
}
