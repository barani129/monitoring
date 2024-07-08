//go:build !ignore_autogenerated

/*
Copyright 2024 baranitharan.chittharanjan@spark.co.nz.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerScan) DeepCopyInto(out *ContainerScan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerScan.
func (in *ContainerScan) DeepCopy() *ContainerScan {
	if in == nil {
		return nil
	}
	out := new(ContainerScan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerScan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerScanCondition) DeepCopyInto(out *ContainerScanCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerScanCondition.
func (in *ContainerScanCondition) DeepCopy() *ContainerScanCondition {
	if in == nil {
		return nil
	}
	out := new(ContainerScanCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerScanList) DeepCopyInto(out *ContainerScanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContainerScan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerScanList.
func (in *ContainerScanList) DeepCopy() *ContainerScanList {
	if in == nil {
		return nil
	}
	out := new(ContainerScanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerScanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerScanSpec) DeepCopyInto(out *ContainerScanSpec) {
	*out = *in
	if in.TargetNamespace != nil {
		in, out := &in.TargetNamespace, &out.TargetNamespace
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SuspendEmailAlert != nil {
		in, out := &in.SuspendEmailAlert, &out.SuspendEmailAlert
		*out = new(bool)
		**out = **in
	}
	if in.AggregateAlerts != nil {
		in, out := &in.AggregateAlerts, &out.AggregateAlerts
		*out = new(bool)
		**out = **in
	}
	if in.NotifyExtenal != nil {
		in, out := &in.NotifyExtenal, &out.NotifyExtenal
		*out = new(bool)
		**out = **in
	}
	if in.CheckInterval != nil {
		in, out := &in.CheckInterval, &out.CheckInterval
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerScanSpec.
func (in *ContainerScanSpec) DeepCopy() *ContainerScanSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerScanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerScanStatus) DeepCopyInto(out *ContainerScanStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ContainerScanCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastRunTime != nil {
		in, out := &in.LastRunTime, &out.LastRunTime
		*out = (*in).DeepCopy()
	}
	if in.ExternalNotifiedTime != nil {
		in, out := &in.ExternalNotifiedTime, &out.ExternalNotifiedTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerScanStatus.
func (in *ContainerScanStatus) DeepCopy() *ContainerScanStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerScanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortScan) DeepCopyInto(out *PortScan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortScan.
func (in *PortScan) DeepCopy() *PortScan {
	if in == nil {
		return nil
	}
	out := new(PortScan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PortScan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortScanCondition) DeepCopyInto(out *PortScanCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortScanCondition.
func (in *PortScanCondition) DeepCopy() *PortScanCondition {
	if in == nil {
		return nil
	}
	out := new(PortScanCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortScanList) DeepCopyInto(out *PortScanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PortScan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortScanList.
func (in *PortScanList) DeepCopy() *PortScanList {
	if in == nil {
		return nil
	}
	out := new(PortScanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PortScanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortScanSpec) DeepCopyInto(out *PortScanSpec) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SuspendEmailAlert != nil {
		in, out := &in.SuspendEmailAlert, &out.SuspendEmailAlert
		*out = new(bool)
		**out = **in
	}
	if in.NotifyExtenal != nil {
		in, out := &in.NotifyExtenal, &out.NotifyExtenal
		*out = new(bool)
		**out = **in
	}
	if in.CheckInterval != nil {
		in, out := &in.CheckInterval, &out.CheckInterval
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortScanSpec.
func (in *PortScanSpec) DeepCopy() *PortScanSpec {
	if in == nil {
		return nil
	}
	out := new(PortScanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortScanStatus) DeepCopyInto(out *PortScanStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PortScanCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastPollTime != nil {
		in, out := &in.LastPollTime, &out.LastPollTime
		*out = (*in).DeepCopy()
	}
	if in.ExternalNotifiedTime != nil {
		in, out := &in.ExternalNotifiedTime, &out.ExternalNotifiedTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortScanStatus.
func (in *PortScanStatus) DeepCopy() *PortScanStatus {
	if in == nil {
		return nil
	}
	out := new(PortScanStatus)
	in.DeepCopyInto(out)
	return out
}
