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

package config

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	apisconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QoSAwareNodeResourcesBalancedAllocationArgs) DeepCopyInto(out *QoSAwareNodeResourcesBalancedAllocationArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]apisconfig.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	if in.ReclaimedResources != nil {
		in, out := &in.ReclaimedResources, &out.ReclaimedResources
		*out = make([]apisconfig.ResourceSpec, len(*in))
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
		*out = make([]apisconfig.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	if in.ReclaimedResources != nil {
		in, out := &in.ReclaimedResources, &out.ReclaimedResources
		*out = make([]apisconfig.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	if in.RequestedToCapacityRatio != nil {
		in, out := &in.RequestedToCapacityRatio, &out.RequestedToCapacityRatio
		*out = new(apisconfig.RequestedToCapacityRatioParam)
		(*in).DeepCopyInto(*out)
	}
	if in.ReclaimedRequestedToCapacityRatio != nil {
		in, out := &in.ReclaimedRequestedToCapacityRatio, &out.ReclaimedRequestedToCapacityRatio
		*out = new(apisconfig.RequestedToCapacityRatioParam)
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
