// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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
	corev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/provider-nop/apis/container/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorConfig) DeepCopyInto(out *AcceleratorConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorConfig.
func (in *AcceleratorConfig) DeepCopy() *AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := new(AcceleratorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoUpgradeOptions) DeepCopyInto(out *AutoUpgradeOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoUpgradeOptions.
func (in *AutoUpgradeOptions) DeepCopy() *AutoUpgradeOptions {
	if in == nil {
		return nil
	}
	out := new(AutoUpgradeOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfig) DeepCopyInto(out *NodeConfig) {
	*out = *in
	if in.Accelerators != nil {
		in, out := &in.Accelerators, &out.Accelerators
		*out = make([]*AcceleratorConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AcceleratorConfig)
				**out = **in
			}
		}
	}
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.ImageType != nil {
		in, out := &in.ImageType, &out.ImageType
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LocalSsdCount != nil {
		in, out := &in.LocalSsdCount, &out.LocalSsdCount
		*out = new(int64)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MinCPUPlatform != nil {
		in, out := &in.MinCPUPlatform, &out.MinCPUPlatform
		*out = new(string)
		**out = **in
	}
	if in.OauthScopes != nil {
		in, out := &in.OauthScopes, &out.OauthScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Preemptible != nil {
		in, out := &in.Preemptible, &out.Preemptible
		*out = new(bool)
		**out = **in
	}
	if in.SandboxConfig != nil {
		in, out := &in.SandboxConfig, &out.SandboxConfig
		*out = new(SandboxConfig)
		**out = **in
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.ShieldedInstanceConfig != nil {
		in, out := &in.ShieldedInstanceConfig, &out.ShieldedInstanceConfig
		*out = new(ShieldedInstanceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]*NodeTaint, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NodeTaint)
				**out = **in
			}
		}
	}
	if in.WorkloadMetadataConfig != nil {
		in, out := &in.WorkloadMetadataConfig, &out.WorkloadMetadataConfig
		*out = new(WorkloadMetadataConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfig.
func (in *NodeConfig) DeepCopy() *NodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeManagementSpec) DeepCopyInto(out *NodeManagementSpec) {
	*out = *in
	if in.AutoRepair != nil {
		in, out := &in.AutoRepair, &out.AutoRepair
		*out = new(bool)
		**out = **in
	}
	if in.AutoUpgrade != nil {
		in, out := &in.AutoUpgrade, &out.AutoUpgrade
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeManagementSpec.
func (in *NodeManagementSpec) DeepCopy() *NodeManagementSpec {
	if in == nil {
		return nil
	}
	out := new(NodeManagementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeManagementStatus) DeepCopyInto(out *NodeManagementStatus) {
	*out = *in
	if in.UpgradeOptions != nil {
		in, out := &in.UpgradeOptions, &out.UpgradeOptions
		*out = new(AutoUpgradeOptions)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeManagementStatus.
func (in *NodeManagementStatus) DeepCopy() *NodeManagementStatus {
	if in == nil {
		return nil
	}
	out := new(NodeManagementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePool) DeepCopyInto(out *NodePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePool.
func (in *NodePool) DeepCopy() *NodePool {
	if in == nil {
		return nil
	}
	out := new(NodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolAutoscaling) DeepCopyInto(out *NodePoolAutoscaling) {
	*out = *in
	if in.Autoprovisioned != nil {
		in, out := &in.Autoprovisioned, &out.Autoprovisioned
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.MaxNodeCount != nil {
		in, out := &in.MaxNodeCount, &out.MaxNodeCount
		*out = new(int64)
		**out = **in
	}
	if in.MinNodeCount != nil {
		in, out := &in.MinNodeCount, &out.MinNodeCount
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolAutoscaling.
func (in *NodePoolAutoscaling) DeepCopy() *NodePoolAutoscaling {
	if in == nil {
		return nil
	}
	out := new(NodePoolAutoscaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolClass) DeepCopyInto(out *NodePoolClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SpecTemplate.DeepCopyInto(&out.SpecTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolClass.
func (in *NodePoolClass) DeepCopy() *NodePoolClass {
	if in == nil {
		return nil
	}
	out := new(NodePoolClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePoolClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolClassList) DeepCopyInto(out *NodePoolClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodePoolClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolClassList.
func (in *NodePoolClassList) DeepCopy() *NodePoolClassList {
	if in == nil {
		return nil
	}
	out := new(NodePoolClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePoolClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolClassSpecTemplate) DeepCopyInto(out *NodePoolClassSpecTemplate) {
	*out = *in
	in.ClassSpecTemplate.DeepCopyInto(&out.ClassSpecTemplate)
	in.NodePoolParameters.DeepCopyInto(&out.NodePoolParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolClassSpecTemplate.
func (in *NodePoolClassSpecTemplate) DeepCopy() *NodePoolClassSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(NodePoolClassSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolList) DeepCopyInto(out *NodePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolList.
func (in *NodePoolList) DeepCopy() *NodePoolList {
	if in == nil {
		return nil
	}
	out := new(NodePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolObservation) DeepCopyInto(out *NodePoolObservation) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*v1beta1.StatusCondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1beta1.StatusCondition)
				**out = **in
			}
		}
	}
	if in.InstanceGroupUrls != nil {
		in, out := &in.InstanceGroupUrls, &out.InstanceGroupUrls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Management != nil {
		in, out := &in.Management, &out.Management
		*out = new(NodeManagementStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolObservation.
func (in *NodePoolObservation) DeepCopy() *NodePoolObservation {
	if in == nil {
		return nil
	}
	out := new(NodePoolObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolParameters) DeepCopyInto(out *NodePoolParameters) {
	*out = *in
	if in.ClusterRef != nil {
		in, out := &in.ClusterRef, &out.ClusterRef
		*out = new(corev1alpha1.Reference)
		**out = **in
	}
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = new(corev1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Autoscaling != nil {
		in, out := &in.Autoscaling, &out.Autoscaling
		*out = new(NodePoolAutoscaling)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(NodeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.InitialNodeCount != nil {
		in, out := &in.InitialNodeCount, &out.InitialNodeCount
		*out = new(int64)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Management != nil {
		in, out := &in.Management, &out.Management
		*out = new(NodeManagementSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxPodsConstraint != nil {
		in, out := &in.MaxPodsConstraint, &out.MaxPodsConstraint
		*out = new(v1beta1.MaxPodsConstraint)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolParameters.
func (in *NodePoolParameters) DeepCopy() *NodePoolParameters {
	if in == nil {
		return nil
	}
	out := new(NodePoolParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpec) DeepCopyInto(out *NodePoolSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpec.
func (in *NodePoolSpec) DeepCopy() *NodePoolSpec {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolStatus) DeepCopyInto(out *NodePoolStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolStatus.
func (in *NodePoolStatus) DeepCopy() *NodePoolStatus {
	if in == nil {
		return nil
	}
	out := new(NodePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTaint) DeepCopyInto(out *NodeTaint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTaint.
func (in *NodeTaint) DeepCopy() *NodeTaint {
	if in == nil {
		return nil
	}
	out := new(NodeTaint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SandboxConfig) DeepCopyInto(out *SandboxConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SandboxConfig.
func (in *SandboxConfig) DeepCopy() *SandboxConfig {
	if in == nil {
		return nil
	}
	out := new(SandboxConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShieldedInstanceConfig) DeepCopyInto(out *ShieldedInstanceConfig) {
	*out = *in
	if in.EnableIntegrityMonitoring != nil {
		in, out := &in.EnableIntegrityMonitoring, &out.EnableIntegrityMonitoring
		*out = new(bool)
		**out = **in
	}
	if in.EnableSecureBoot != nil {
		in, out := &in.EnableSecureBoot, &out.EnableSecureBoot
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShieldedInstanceConfig.
func (in *ShieldedInstanceConfig) DeepCopy() *ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := new(ShieldedInstanceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadMetadataConfig) DeepCopyInto(out *WorkloadMetadataConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadMetadataConfig.
func (in *WorkloadMetadataConfig) DeepCopy() *WorkloadMetadataConfig {
	if in == nil {
		return nil
	}
	out := new(WorkloadMetadataConfig)
	in.DeepCopyInto(out)
	return out
}
