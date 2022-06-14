//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAllocationDetail) DeepCopyInto(out *IPAllocationDetail) {
	*out = *in
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(string)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(string)
		**out = **in
	}
	if in.IPv4Pool != nil {
		in, out := &in.IPv4Pool, &out.IPv4Pool
		*out = new(string)
		**out = **in
	}
	if in.IPv6Pool != nil {
		in, out := &in.IPv6Pool, &out.IPv6Pool
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(Vlan)
		**out = **in
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		copy(*out, *in)
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAllocationDetail.
func (in *IPAllocationDetail) DeepCopy() *IPAllocationDetail {
	if in == nil {
		return nil
	}
	out := new(IPAllocationDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPool) DeepCopyInto(out *IPPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPool.
func (in *IPPool) DeepCopy() *IPPool {
	if in == nil {
		return nil
	}
	out := new(IPPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolList) DeepCopyInto(out *IPPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolList.
func (in *IPPoolList) DeepCopy() *IPPoolList {
	if in == nil {
		return nil
	}
	out := new(IPPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolSpec) DeepCopyInto(out *IPPoolSpec) {
	*out = *in
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(IPVersion)
		**out = **in
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Disable != nil {
		in, out := &in.Disable, &out.Disable
		*out = new(bool)
		**out = **in
	}
	if in.ExcludeIPs != nil {
		in, out := &in.ExcludeIPs, &out.ExcludeIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		copy(*out, *in)
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(Vlan)
		**out = **in
	}
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamesapceSelector != nil {
		in, out := &in.NamesapceSelector, &out.NamesapceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolSpec.
func (in *IPPoolSpec) DeepCopy() *IPPoolSpec {
	if in == nil {
		return nil
	}
	out := new(IPPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolStatus) DeepCopyInto(out *IPPoolStatus) {
	*out = *in
	if in.AllocatedIPs != nil {
		in, out := &in.AllocatedIPs, &out.AllocatedIPs
		*out = make(PoolIPAllocations, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AllocateCount != nil {
		in, out := &in.AllocateCount, &out.AllocateCount
		*out = new(int64)
		**out = **in
	}
	if in.DeallocateCount != nil {
		in, out := &in.DeallocateCount, &out.DeallocateCount
		*out = new(int64)
		**out = **in
	}
	if in.TotalIPCount != nil {
		in, out := &in.TotalIPCount, &out.TotalIPCount
		*out = new(int32)
		**out = **in
	}
	if in.AllocatedIPCount != nil {
		in, out := &in.AllocatedIPCount, &out.AllocatedIPCount
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolStatus.
func (in *IPPoolStatus) DeepCopy() *IPPoolStatus {
	if in == nil {
		return nil
	}
	out := new(IPPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodIPAllocation) DeepCopyInto(out *PodIPAllocation) {
	*out = *in
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]IPAllocationDetail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CreationTime.DeepCopyInto(&out.CreationTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodIPAllocation.
func (in *PodIPAllocation) DeepCopy() *PodIPAllocation {
	if in == nil {
		return nil
	}
	out := new(PodIPAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolIPAllocation) DeepCopyInto(out *PoolIPAllocation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolIPAllocation.
func (in *PoolIPAllocation) DeepCopy() *PoolIPAllocation {
	if in == nil {
		return nil
	}
	out := new(PoolIPAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PoolIPAllocations) DeepCopyInto(out *PoolIPAllocations) {
	{
		in := &in
		*out = make(PoolIPAllocations, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolIPAllocations.
func (in PoolIPAllocations) DeepCopy() PoolIPAllocations {
	if in == nil {
		return nil
	}
	out := new(PoolIPAllocations)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservedIP) DeepCopyInto(out *ReservedIP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservedIP.
func (in *ReservedIP) DeepCopy() *ReservedIP {
	if in == nil {
		return nil
	}
	out := new(ReservedIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReservedIP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservedIPList) DeepCopyInto(out *ReservedIPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReservedIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservedIPList.
func (in *ReservedIPList) DeepCopy() *ReservedIPList {
	if in == nil {
		return nil
	}
	out := new(ReservedIPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReservedIPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservedIPSpec) DeepCopyInto(out *ReservedIPSpec) {
	*out = *in
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(IPVersion)
		**out = **in
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservedIPSpec.
func (in *ReservedIPSpec) DeepCopy() *ReservedIPSpec {
	if in == nil {
		return nil
	}
	out := new(ReservedIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadEndpoint) DeepCopyInto(out *WorkloadEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadEndpoint.
func (in *WorkloadEndpoint) DeepCopy() *WorkloadEndpoint {
	if in == nil {
		return nil
	}
	out := new(WorkloadEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadEndpointList) DeepCopyInto(out *WorkloadEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkloadEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadEndpointList.
func (in *WorkloadEndpointList) DeepCopy() *WorkloadEndpointList {
	if in == nil {
		return nil
	}
	out := new(WorkloadEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadEndpointStatus) DeepCopyInto(out *WorkloadEndpointStatus) {
	*out = *in
	if in.Current != nil {
		in, out := &in.Current, &out.Current
		*out = new(PodIPAllocation)
		(*in).DeepCopyInto(*out)
	}
	if in.History != nil {
		in, out := &in.History, &out.History
		*out = make([]PodIPAllocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadEndpointStatus.
func (in *WorkloadEndpointStatus) DeepCopy() *WorkloadEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadEndpointStatus)
	in.DeepCopyInto(out)
	return out
}
