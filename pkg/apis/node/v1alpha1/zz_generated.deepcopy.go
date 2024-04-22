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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Allocation) DeepCopyInto(out *Allocation) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(v1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Allocation.
func (in *Allocation) DeepCopy() *Allocation {
	if in == nil {
		return nil
	}
	out := new(Allocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attribute) DeepCopyInto(out *Attribute) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attribute.
func (in *Attribute) DeepCopy() *Attribute {
	if in == nil {
		return nil
	}
	out := new(Attribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNRCondition) DeepCopyInto(out *CNRCondition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNRCondition.
func (in *CNRCondition) DeepCopy() *CNRCondition {
	if in == nil {
		return nil
	}
	out := new(CNRCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeResource) DeepCopyInto(out *CustomNodeResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeResource.
func (in *CustomNodeResource) DeepCopy() *CustomNodeResource {
	if in == nil {
		return nil
	}
	out := new(CustomNodeResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomNodeResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeResourceList) DeepCopyInto(out *CustomNodeResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomNodeResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeResourceList.
func (in *CustomNodeResourceList) DeepCopy() *CustomNodeResourceList {
	if in == nil {
		return nil
	}
	out := new(CustomNodeResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomNodeResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeResourceSpec) DeepCopyInto(out *CustomNodeResourceSpec) {
	*out = *in
	if in.NodeResourceProperties != nil {
		in, out := &in.NodeResourceProperties, &out.NodeResourceProperties
		*out = make([]*Property, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Property)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]*Taint, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Taint)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeResourceSpec.
func (in *CustomNodeResourceSpec) DeepCopy() *CustomNodeResourceSpec {
	if in == nil {
		return nil
	}
	out := new(CustomNodeResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeResourceStatus) DeepCopyInto(out *CustomNodeResourceStatus) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.TopologyZone != nil {
		in, out := &in.TopologyZone, &out.TopologyZone
		*out = make([]*TopologyZone, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TopologyZone)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CNRCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeMetricStatus != nil {
		in, out := &in.NodeMetricStatus, &out.NodeMetricStatus
		*out = new(NodeMetricStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeResourceStatus.
func (in *CustomNodeResourceStatus) DeepCopy() *CustomNodeResourceStatus {
	if in == nil {
		return nil
	}
	out := new(CustomNodeResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupMetricInfo) DeepCopyInto(out *GroupMetricInfo) {
	*out = *in
	in.ResourceUsage.DeepCopyInto(&out.ResourceUsage)
	if in.PodList != nil {
		in, out := &in.PodList, &out.PodList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupMetricInfo.
func (in *GroupMetricInfo) DeepCopy() *GroupMetricInfo {
	if in == nil {
		return nil
	}
	out := new(GroupMetricInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NUMAMetricInfo) DeepCopyInto(out *NUMAMetricInfo) {
	*out = *in
	if in.Usage != nil {
		in, out := &in.Usage, &out.Usage
		*out = new(ResourceMetric)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NUMAMetricInfo.
func (in *NUMAMetricInfo) DeepCopy() *NUMAMetricInfo {
	if in == nil {
		return nil
	}
	out := new(NUMAMetricInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMetricInfo) DeepCopyInto(out *NodeMetricInfo) {
	*out = *in
	in.ResourceUsage.DeepCopyInto(&out.ResourceUsage)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMetricInfo.
func (in *NodeMetricInfo) DeepCopy() *NodeMetricInfo {
	if in == nil {
		return nil
	}
	out := new(NodeMetricInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMetricStatus) DeepCopyInto(out *NodeMetricStatus) {
	*out = *in
	in.UpdateTime.DeepCopyInto(&out.UpdateTime)
	if in.NodeMetric != nil {
		in, out := &in.NodeMetric, &out.NodeMetric
		*out = new(NodeMetricInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.GroupMetric != nil {
		in, out := &in.GroupMetric, &out.GroupMetric
		*out = make([]GroupMetricInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMetricStatus.
func (in *NodeMetricStatus) DeepCopy() *NodeMetricStatus {
	if in == nil {
		return nil
	}
	out := new(NodeMetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMonitor) DeepCopyInto(out *NodeMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMonitor.
func (in *NodeMonitor) DeepCopy() *NodeMonitor {
	if in == nil {
		return nil
	}
	out := new(NodeMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMonitorList) DeepCopyInto(out *NodeMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMonitorList.
func (in *NodeMonitorList) DeepCopy() *NodeMonitorList {
	if in == nil {
		return nil
	}
	out := new(NodeMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMonitorSpec) DeepCopyInto(out *NodeMonitorSpec) {
	*out = *in
	if in.ReportInterval != nil {
		in, out := &in.ReportInterval, &out.ReportInterval
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMonitorSpec.
func (in *NodeMonitorSpec) DeepCopy() *NodeMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(NodeMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMonitorStatus) DeepCopyInto(out *NodeMonitorStatus) {
	*out = *in
	if in.NodeUsage != nil {
		in, out := &in.NodeUsage, &out.NodeUsage
		*out = make(map[string]v1.ResourceList, len(*in))
		for key, val := range *in {
			var outVal map[v1.ResourceName]resource.Quantity
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(v1.ResourceList, len(*in))
				for key, val := range *in {
					(*out)[key] = val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.PodUsage != nil {
		in, out := &in.PodUsage, &out.PodUsage
		*out = make(map[string]v1.ResourceList, len(*in))
		for key, val := range *in {
			var outVal map[v1.ResourceName]resource.Quantity
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(v1.ResourceList, len(*in))
				for key, val := range *in {
					(*out)[key] = val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMonitorStatus.
func (in *NodeMonitorStatus) DeepCopy() *NodeMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(NodeMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Property) DeepCopyInto(out *Property) {
	*out = *in
	if in.PropertyValues != nil {
		in, out := &in.PropertyValues, &out.PropertyValues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PropertyQuantity != nil {
		in, out := &in.PropertyQuantity, &out.PropertyQuantity
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Property.
func (in *Property) DeepCopy() *Property {
	if in == nil {
		return nil
	}
	out := new(Property)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMetric) DeepCopyInto(out *ResourceMetric) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMetric.
func (in *ResourceMetric) DeepCopy() *ResourceMetric {
	if in == nil {
		return nil
	}
	out := new(ResourceMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceUsage) DeepCopyInto(out *ResourceUsage) {
	*out = *in
	if in.NUMAUsage != nil {
		in, out := &in.NUMAUsage, &out.NUMAUsage
		*out = make([]NUMAMetricInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GenericUsage != nil {
		in, out := &in.GenericUsage, &out.GenericUsage
		*out = new(ResourceMetric)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceUsage.
func (in *ResourceUsage) DeepCopy() *ResourceUsage {
	if in == nil {
		return nil
	}
	out := new(ResourceUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.Allocatable != nil {
		in, out := &in.Allocatable, &out.Allocatable
		*out = new(v1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(v1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sibling) DeepCopyInto(out *Sibling) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make([]Attribute, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sibling.
func (in *Sibling) DeepCopy() *Sibling {
	if in == nil {
		return nil
	}
	out := new(Sibling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Taint) DeepCopyInto(out *Taint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Taint.
func (in *Taint) DeepCopy() *Taint {
	if in == nil {
		return nil
	}
	out := new(Taint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopologyZone) DeepCopyInto(out *TopologyZone) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make([]Attribute, len(*in))
		copy(*out, *in)
	}
	if in.Allocations != nil {
		in, out := &in.Allocations, &out.Allocations
		*out = make([]*Allocation, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Allocation)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Children != nil {
		in, out := &in.Children, &out.Children
		*out = make([]*TopologyZone, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TopologyZone)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Siblings != nil {
		in, out := &in.Siblings, &out.Siblings
		*out = make([]Sibling, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopologyZone.
func (in *TopologyZone) DeepCopy() *TopologyZone {
	if in == nil {
		return nil
	}
	out := new(TopologyZone)
	in.DeepCopyInto(out)
	return out
}
