package v1beta1

import (
	"github.com/kubewharf/katalyst-api/pkg/apis/scheduling/config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kube-scheduler/config/v1beta1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// QoSAwareNodeResourcesFitArgs holds arguments used to configure the QoSAwareNodeResourcesFit plugin.
type QoSAwareNodeResourcesFitArgs struct {
	metav1.TypeMeta `json:",inline"`

	// ScoringStrategy selects the node resource scoring strategy.
	ScoringStrategy *ScoringStrategy `json:"scoringStrategy,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// QoSAwareNodeResourcesBalancedAllocationArgs holds arguments used to configure QoSAwareNodeResourcesBalancedAllocation plugin.
type QoSAwareNodeResourcesBalancedAllocationArgs struct {
	metav1.TypeMeta `json:",inline"`

	// Resources to be considered when scoring.
	// The default resource set includes "cpu" and "memory", only valid weight is 1.
	Resources []v1beta1.ResourceSpec `json:"resources,omitempty"`

	// ReclaimedResources to be considered when scoring.
	// The default resource set includes "resource.katalyst.kubewharf.io/reclaimed_millicpu"
	// and "resource.katalyst.kubewharf.io/reclaimed_memory", only valid weight is 1.
	ReclaimedResources []v1beta1.ResourceSpec `json:"reclaimedResources,omitempty"`
}

// ScoringStrategy define ScoringStrategyType for node resource plugin
type ScoringStrategy struct {
	// Type selects which strategy to run.
	Type config.ScoringStrategyType `json:"type,omitempty"`

	// Resources to consider when scoring.
	// The default resource set includes "cpu" and "memory" with an equal weight.
	// Allowed weights go from 1 to 100.
	// Weight defaults to 1 if not specified or explicitly set to 0.
	Resources []v1beta1.ResourceSpec `json:"resources,omitempty"`

	// ReclaimedResources to consider when scoring.
	// The default resource set includes "resource.katalyst.kubewharf.io/reclaimed_millicpu"
	// and "resource.katalyst.kubewharf.io/reclaimed_memory", only valid weight is 1.
	ReclaimedResources []v1beta1.ResourceSpec `json:"reclaimedResources,omitempty"`

	// Arguments specific to RequestedToCapacityRatio strategy.
	RequestedToCapacityRatio *v1beta1.RequestedToCapacityRatioArgs `json:"requestedToCapacityRatio,omitempty"`

	// Arguments specific to RequestedToCapacityRatio strategy.
	ReclaimedRequestedToCapacityRatio *v1beta1.RequestedToCapacityRatioArgs `json:"reclaimedRequestedToCapacityRatio,omitempty"`
}
