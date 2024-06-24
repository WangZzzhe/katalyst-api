//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta3

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1beta3 "k8s.io/kube-scheduler/config/v1beta3"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadAwareArgs) DeepCopyInto(out *LoadAwareArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.PodAnnotationLoadAwareEnable != nil {
		in, out := &in.PodAnnotationLoadAwareEnable, &out.PodAnnotationLoadAwareEnable
		*out = new(string)
		**out = **in
	}
	if in.FilterExpiredNodeMetrics != nil {
		in, out := &in.FilterExpiredNodeMetrics, &out.FilterExpiredNodeMetrics
		*out = new(bool)
		**out = **in
	}
	if in.NodeMetricsExpiredSeconds != nil {
		in, out := &in.NodeMetricsExpiredSeconds, &out.NodeMetricsExpiredSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ResourceToWeightMap != nil {
		in, out := &in.ResourceToWeightMap, &out.ResourceToWeightMap
		*out = make(map[v1.ResourceName]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ResourceToThresholdMap != nil {
		in, out := &in.ResourceToThresholdMap, &out.ResourceToThresholdMap
		*out = make(map[v1.ResourceName]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ResourceToScalingFactorMap != nil {
		in, out := &in.ResourceToScalingFactorMap, &out.ResourceToScalingFactorMap
		*out = make(map[v1.ResourceName]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CalculateIndicatorWeight != nil {
		in, out := &in.CalculateIndicatorWeight, &out.CalculateIndicatorWeight
		*out = make(map[IndicatorType]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadAwareArgs.
func (in *LoadAwareArgs) DeepCopy() *LoadAwareArgs {
	if in == nil {
		return nil
	}
	out := new(LoadAwareArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadAwareArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourceTopologyArgs) DeepCopyInto(out *NodeResourceTopologyArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ScoringStrategy != nil {
		in, out := &in.ScoringStrategy, &out.ScoringStrategy
		*out = new(ScoringStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.AlignedResources != nil {
		in, out := &in.AlignedResources, &out.AlignedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourceTopologyArgs.
func (in *NodeResourceTopologyArgs) DeepCopy() *NodeResourceTopologyArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourceTopologyArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourceTopologyArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QoSAwareNodeResourcesBalancedAllocationArgs) DeepCopyInto(out *QoSAwareNodeResourcesBalancedAllocationArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]configv1beta3.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	if in.ReclaimedResources != nil {
		in, out := &in.ReclaimedResources, &out.ReclaimedResources
		*out = make([]configv1beta3.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QoSAwareNodeResourcesBalancedAllocationArgs.
func (in *QoSAwareNodeResourcesBalancedAllocationArgs) DeepCopy() *QoSAwareNodeResourcesBalancedAllocationArgs {
	if in == nil {
		return nil
	}
	out := new(QoSAwareNodeResourcesBalancedAllocationArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QoSAwareNodeResourcesBalancedAllocationArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QoSAwareNodeResourcesFitArgs) DeepCopyInto(out *QoSAwareNodeResourcesFitArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ScoringStrategy != nil {
		in, out := &in.ScoringStrategy, &out.ScoringStrategy
		*out = new(ScoringStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QoSAwareNodeResourcesFitArgs.
func (in *QoSAwareNodeResourcesFitArgs) DeepCopy() *QoSAwareNodeResourcesFitArgs {
	if in == nil {
		return nil
	}
	out := new(QoSAwareNodeResourcesFitArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QoSAwareNodeResourcesFitArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScoringStrategy) DeepCopyInto(out *ScoringStrategy) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]configv1beta3.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	if in.ReclaimedResources != nil {
		in, out := &in.ReclaimedResources, &out.ReclaimedResources
		*out = make([]configv1beta3.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	if in.RequestedToCapacityRatio != nil {
		in, out := &in.RequestedToCapacityRatio, &out.RequestedToCapacityRatio
		*out = new(configv1beta3.RequestedToCapacityRatioParam)
		(*in).DeepCopyInto(*out)
	}
	if in.ReclaimedRequestedToCapacityRatio != nil {
		in, out := &in.ReclaimedRequestedToCapacityRatio, &out.ReclaimedRequestedToCapacityRatio
		*out = new(configv1beta3.RequestedToCapacityRatioParam)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoringStrategy.
func (in *ScoringStrategy) DeepCopy() *ScoringStrategy {
	if in == nil {
		return nil
	}
	out := new(ScoringStrategy)
	in.DeepCopyInto(out)
	return out
}
