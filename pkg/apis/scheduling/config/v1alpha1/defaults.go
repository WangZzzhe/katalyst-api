package v1alpha1

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/kube-scheduler/config/v1"

	"github.com/kubewharf/katalyst-api/pkg/apis/scheduling/config"
	"github.com/kubewharf/katalyst-api/pkg/consts"
)

var defaultResourceSpec = []v1.ResourceSpec{
	{Name: string(corev1.ResourceCPU), Weight: 1},
	{Name: string(corev1.ResourceMemory), Weight: 1},
}

var defaultReclaimedResourceSpec = []v1.ResourceSpec{
	{Name: fmt.Sprintf("%s", consts.ReclaimedResourceMilliCPU), Weight: 1},
	{Name: fmt.Sprintf("%s", consts.ReclaimedResourceMemory), Weight: 1},
}

// SetDefaults_QoSAwareNodeResourcesFitArgs sets the default parameters for QoSAwareNodeResourcesFit plugin.
func SetDefaults_QoSAwareNodeResourcesFitArgs(obj *QoSAwareNodeResourcesFitArgs) {
	if obj.ScoringStrategy == nil {
		obj.ScoringStrategy = &ScoringStrategy{
			Type:               config.LeastAllocated,
			Resources:          defaultResourceSpec,
			ReclaimedResources: defaultReclaimedResourceSpec,
		}
	}
	if len(obj.ScoringStrategy.Resources) == 0 {
		// If no resources specified, use the default set.
		obj.ScoringStrategy.Resources = append(obj.ScoringStrategy.Resources, defaultResourceSpec...)
	}
	if len(obj.ScoringStrategy.ReclaimedResources) == 0 {
		obj.ScoringStrategy.ReclaimedResources = append(obj.ScoringStrategy.ReclaimedResources, defaultReclaimedResourceSpec...)
	}
	for i := range obj.ScoringStrategy.Resources {
		if obj.ScoringStrategy.Resources[i].Weight == 0 {
			obj.ScoringStrategy.Resources[i].Weight = 1
		}
	}
	for i := range obj.ScoringStrategy.ReclaimedResources {
		if obj.ScoringStrategy.ReclaimedResources[i].Weight == 0 {
			obj.ScoringStrategy.ReclaimedResources[i].Weight = 1
		}
	}
}

// SetDefaults_QoSAwareNodeResourcesBalancedAllocationArgs sets the default parameters for QoSAwareNodeResourcesBalancedAllocation plugin.
func SetDefaults_QoSAwareNodeResourcesBalancedAllocationArgs(obj *QoSAwareNodeResourcesBalancedAllocationArgs) {
	if len(obj.Resources) == 0 {
		obj.Resources = append(obj.Resources, defaultResourceSpec...)
	}
	if len(obj.ReclaimedResources) == 0 {
		obj.ReclaimedResources = append(obj.ReclaimedResources, defaultReclaimedResourceSpec...)
	}
	// If the weight is not set or it is explicitly set to 0, then apply the default weight(1) instead.
	for i := range obj.Resources {
		if obj.Resources[i].Weight == 0 {
			obj.Resources[i].Weight = 1
		}
	}
	for i := range obj.ReclaimedResources {
		if obj.ReclaimedResources[i].Weight == 0 {
			obj.ReclaimedResources[i].Weight = 1
		}
	}
}
