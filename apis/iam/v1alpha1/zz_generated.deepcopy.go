// +build !ignore_autogenerated

/*
Copyright © 2019 AWS Controller authors

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
	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessKey) DeepCopyInto(out *AccessKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessKey.
func (in *AccessKey) DeepCopy() *AccessKey {
	if in == nil {
		return nil
	}
	out := new(AccessKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessKeyList) DeepCopyInto(out *AccessKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessKeyList.
func (in *AccessKeyList) DeepCopy() *AccessKeyList {
	if in == nil {
		return nil
	}
	out := new(AccessKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessKeyOutput) DeepCopyInto(out *AccessKeyOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessKeyOutput.
func (in *AccessKeyOutput) DeepCopy() *AccessKeyOutput {
	if in == nil {
		return nil
	}
	out := new(AccessKeyOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessKeySpec) DeepCopyInto(out *AccessKeySpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessKeySpec.
func (in *AccessKeySpec) DeepCopy() *AccessKeySpec {
	if in == nil {
		return nil
	}
	out := new(AccessKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessKeyStatus) DeepCopyInto(out *AccessKeyStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessKeyStatus.
func (in *AccessKeyStatus) DeepCopy() *AccessKeyStatus {
	if in == nil {
		return nil
	}
	out := new(AccessKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Group) DeepCopyInto(out *Group) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Group.
func (in *Group) DeepCopy() *Group {
	if in == nil {
		return nil
	}
	out := new(Group)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Group) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupList) DeepCopyInto(out *GroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Group, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupList.
func (in *GroupList) DeepCopy() *GroupList {
	if in == nil {
		return nil
	}
	out := new(GroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupOutput) DeepCopyInto(out *GroupOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupOutput.
func (in *GroupOutput) DeepCopy() *GroupOutput {
	if in == nil {
		return nil
	}
	out := new(GroupOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSpec) DeepCopyInto(out *GroupSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.ManagedPolicyRefs != nil {
		in, out := &in.ManagedPolicyRefs, &out.ManagedPolicyRefs
		*out = make([]metav1alpha1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]Group_Policy, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSpec.
func (in *GroupSpec) DeepCopy() *GroupSpec {
	if in == nil {
		return nil
	}
	out := new(GroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupStatus) DeepCopyInto(out *GroupStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupStatus.
func (in *GroupStatus) DeepCopy() *GroupStatus {
	if in == nil {
		return nil
	}
	out := new(GroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Group_Policy) DeepCopyInto(out *Group_Policy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Group_Policy.
func (in *Group_Policy) DeepCopy() *Group_Policy {
	if in == nil {
		return nil
	}
	out := new(Group_Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceProfile) DeepCopyInto(out *InstanceProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceProfile.
func (in *InstanceProfile) DeepCopy() *InstanceProfile {
	if in == nil {
		return nil
	}
	out := new(InstanceProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceProfileList) DeepCopyInto(out *InstanceProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstanceProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceProfileList.
func (in *InstanceProfileList) DeepCopy() *InstanceProfileList {
	if in == nil {
		return nil
	}
	out := new(InstanceProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceProfileOutput) DeepCopyInto(out *InstanceProfileOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceProfileOutput.
func (in *InstanceProfileOutput) DeepCopy() *InstanceProfileOutput {
	if in == nil {
		return nil
	}
	out := new(InstanceProfileOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceProfileSpec) DeepCopyInto(out *InstanceProfileSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceProfileSpec.
func (in *InstanceProfileSpec) DeepCopy() *InstanceProfileSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceProfileStatus) DeepCopyInto(out *InstanceProfileStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceProfileStatus.
func (in *InstanceProfileStatus) DeepCopy() *InstanceProfileStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPolicy) DeepCopyInto(out *ManagedPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPolicy.
func (in *ManagedPolicy) DeepCopy() *ManagedPolicy {
	if in == nil {
		return nil
	}
	out := new(ManagedPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPolicyList) DeepCopyInto(out *ManagedPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPolicyList.
func (in *ManagedPolicyList) DeepCopy() *ManagedPolicyList {
	if in == nil {
		return nil
	}
	out := new(ManagedPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPolicyOutput) DeepCopyInto(out *ManagedPolicyOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPolicyOutput.
func (in *ManagedPolicyOutput) DeepCopy() *ManagedPolicyOutput {
	if in == nil {
		return nil
	}
	out := new(ManagedPolicyOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPolicySpec) DeepCopyInto(out *ManagedPolicySpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPolicySpec.
func (in *ManagedPolicySpec) DeepCopy() *ManagedPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ManagedPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPolicyStatus) DeepCopyInto(out *ManagedPolicyStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPolicyStatus.
func (in *ManagedPolicyStatus) DeepCopy() *ManagedPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyOutput) DeepCopyInto(out *PolicyOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyOutput.
func (in *PolicyOutput) DeepCopy() *PolicyOutput {
	if in == nil {
		return nil
	}
	out := new(PolicyOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Role) DeepCopyInto(out *Role) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Role.
func (in *Role) DeepCopy() *Role {
	if in == nil {
		return nil
	}
	out := new(Role)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Role) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleList) DeepCopyInto(out *RoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Role, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleList.
func (in *RoleList) DeepCopy() *RoleList {
	if in == nil {
		return nil
	}
	out := new(RoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleOutput) DeepCopyInto(out *RoleOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleOutput.
func (in *RoleOutput) DeepCopy() *RoleOutput {
	if in == nil {
		return nil
	}
	out := new(RoleOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleSpec) DeepCopyInto(out *RoleSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.ManagedPolicyRefs != nil {
		in, out := &in.ManagedPolicyRefs, &out.ManagedPolicyRefs
		*out = make([]metav1alpha1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]Role_Policy, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleSpec.
func (in *RoleSpec) DeepCopy() *RoleSpec {
	if in == nil {
		return nil
	}
	out := new(RoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleStatus) DeepCopyInto(out *RoleStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleStatus.
func (in *RoleStatus) DeepCopy() *RoleStatus {
	if in == nil {
		return nil
	}
	out := new(RoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Role_Policy) DeepCopyInto(out *Role_Policy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Role_Policy.
func (in *Role_Policy) DeepCopy() *Role_Policy {
	if in == nil {
		return nil
	}
	out := new(Role_Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLinkedRole) DeepCopyInto(out *ServiceLinkedRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLinkedRole.
func (in *ServiceLinkedRole) DeepCopy() *ServiceLinkedRole {
	if in == nil {
		return nil
	}
	out := new(ServiceLinkedRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceLinkedRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLinkedRoleList) DeepCopyInto(out *ServiceLinkedRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceLinkedRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLinkedRoleList.
func (in *ServiceLinkedRoleList) DeepCopy() *ServiceLinkedRoleList {
	if in == nil {
		return nil
	}
	out := new(ServiceLinkedRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceLinkedRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLinkedRoleOutput) DeepCopyInto(out *ServiceLinkedRoleOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLinkedRoleOutput.
func (in *ServiceLinkedRoleOutput) DeepCopy() *ServiceLinkedRoleOutput {
	if in == nil {
		return nil
	}
	out := new(ServiceLinkedRoleOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLinkedRoleSpec) DeepCopyInto(out *ServiceLinkedRoleSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLinkedRoleSpec.
func (in *ServiceLinkedRoleSpec) DeepCopy() *ServiceLinkedRoleSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceLinkedRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLinkedRoleStatus) DeepCopyInto(out *ServiceLinkedRoleStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLinkedRoleStatus.
func (in *ServiceLinkedRoleStatus) DeepCopy() *ServiceLinkedRoleStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceLinkedRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *User) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserList) DeepCopyInto(out *UserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserList.
func (in *UserList) DeepCopy() *UserList {
	if in == nil {
		return nil
	}
	out := new(UserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserOutput) DeepCopyInto(out *UserOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserOutput.
func (in *UserOutput) DeepCopy() *UserOutput {
	if in == nil {
		return nil
	}
	out := new(UserOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.LoginProfile = in.LoginProfile
	if in.ManagedPolicyRefs != nil {
		in, out := &in.ManagedPolicyRefs, &out.ManagedPolicyRefs
		*out = make([]metav1alpha1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]User_Policy, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserStatus) DeepCopyInto(out *UserStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserStatus.
func (in *UserStatus) DeepCopy() *UserStatus {
	if in == nil {
		return nil
	}
	out := new(UserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserToGroupAddition) DeepCopyInto(out *UserToGroupAddition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserToGroupAddition.
func (in *UserToGroupAddition) DeepCopy() *UserToGroupAddition {
	if in == nil {
		return nil
	}
	out := new(UserToGroupAddition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserToGroupAddition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserToGroupAdditionList) DeepCopyInto(out *UserToGroupAdditionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserToGroupAddition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserToGroupAdditionList.
func (in *UserToGroupAdditionList) DeepCopy() *UserToGroupAdditionList {
	if in == nil {
		return nil
	}
	out := new(UserToGroupAdditionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserToGroupAdditionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserToGroupAdditionOutput) DeepCopyInto(out *UserToGroupAdditionOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserToGroupAdditionOutput.
func (in *UserToGroupAdditionOutput) DeepCopy() *UserToGroupAdditionOutput {
	if in == nil {
		return nil
	}
	out := new(UserToGroupAdditionOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserToGroupAdditionSpec) DeepCopyInto(out *UserToGroupAdditionSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserToGroupAdditionSpec.
func (in *UserToGroupAdditionSpec) DeepCopy() *UserToGroupAdditionSpec {
	if in == nil {
		return nil
	}
	out := new(UserToGroupAdditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserToGroupAdditionStatus) DeepCopyInto(out *UserToGroupAdditionStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserToGroupAdditionStatus.
func (in *UserToGroupAdditionStatus) DeepCopy() *UserToGroupAdditionStatus {
	if in == nil {
		return nil
	}
	out := new(UserToGroupAdditionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User_LoginProfile) DeepCopyInto(out *User_LoginProfile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User_LoginProfile.
func (in *User_LoginProfile) DeepCopy() *User_LoginProfile {
	if in == nil {
		return nil
	}
	out := new(User_LoginProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User_Policy) DeepCopyInto(out *User_Policy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User_Policy.
func (in *User_Policy) DeepCopy() *User_Policy {
	if in == nil {
		return nil
	}
	out := new(User_Policy)
	in.DeepCopyInto(out)
	return out
}
