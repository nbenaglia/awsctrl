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
func (in *Alias) DeepCopyInto(out *Alias) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias.
func (in *Alias) DeepCopy() *Alias {
	if in == nil {
		return nil
	}
	out := new(Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alias) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasList) DeepCopyInto(out *AliasList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasList.
func (in *AliasList) DeepCopy() *AliasList {
	if in == nil {
		return nil
	}
	out := new(AliasList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AliasList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasOutput) DeepCopyInto(out *AliasOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasOutput.
func (in *AliasOutput) DeepCopy() *AliasOutput {
	if in == nil {
		return nil
	}
	out := new(AliasOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasSpec) DeepCopyInto(out *AliasSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.ProvisionedConcurrencyConfig = in.ProvisionedConcurrencyConfig
	in.RoutingConfig.DeepCopyInto(&out.RoutingConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasSpec.
func (in *AliasSpec) DeepCopy() *AliasSpec {
	if in == nil {
		return nil
	}
	out := new(AliasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasStatus) DeepCopyInto(out *AliasStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasStatus.
func (in *AliasStatus) DeepCopy() *AliasStatus {
	if in == nil {
		return nil
	}
	out := new(AliasStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alias_AliasRoutingConfiguration) DeepCopyInto(out *Alias_AliasRoutingConfiguration) {
	*out = *in
	if in.AdditionalVersionWeights != nil {
		in, out := &in.AdditionalVersionWeights, &out.AdditionalVersionWeights
		*out = make([]Alias_VersionWeight, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias_AliasRoutingConfiguration.
func (in *Alias_AliasRoutingConfiguration) DeepCopy() *Alias_AliasRoutingConfiguration {
	if in == nil {
		return nil
	}
	out := new(Alias_AliasRoutingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alias_ProvisionedConcurrencyConfiguration) DeepCopyInto(out *Alias_ProvisionedConcurrencyConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias_ProvisionedConcurrencyConfiguration.
func (in *Alias_ProvisionedConcurrencyConfiguration) DeepCopy() *Alias_ProvisionedConcurrencyConfiguration {
	if in == nil {
		return nil
	}
	out := new(Alias_ProvisionedConcurrencyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alias_VersionWeight) DeepCopyInto(out *Alias_VersionWeight) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias_VersionWeight.
func (in *Alias_VersionWeight) DeepCopy() *Alias_VersionWeight {
	if in == nil {
		return nil
	}
	out := new(Alias_VersionWeight)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventInvokeConfig) DeepCopyInto(out *EventInvokeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventInvokeConfig.
func (in *EventInvokeConfig) DeepCopy() *EventInvokeConfig {
	if in == nil {
		return nil
	}
	out := new(EventInvokeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventInvokeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventInvokeConfigList) DeepCopyInto(out *EventInvokeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventInvokeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventInvokeConfigList.
func (in *EventInvokeConfigList) DeepCopy() *EventInvokeConfigList {
	if in == nil {
		return nil
	}
	out := new(EventInvokeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventInvokeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventInvokeConfigOutput) DeepCopyInto(out *EventInvokeConfigOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventInvokeConfigOutput.
func (in *EventInvokeConfigOutput) DeepCopy() *EventInvokeConfigOutput {
	if in == nil {
		return nil
	}
	out := new(EventInvokeConfigOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventInvokeConfigSpec) DeepCopyInto(out *EventInvokeConfigSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.DestinationConfig = in.DestinationConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventInvokeConfigSpec.
func (in *EventInvokeConfigSpec) DeepCopy() *EventInvokeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(EventInvokeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventInvokeConfigStatus) DeepCopyInto(out *EventInvokeConfigStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventInvokeConfigStatus.
func (in *EventInvokeConfigStatus) DeepCopy() *EventInvokeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(EventInvokeConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventInvokeConfig_DestinationConfig) DeepCopyInto(out *EventInvokeConfig_DestinationConfig) {
	*out = *in
	out.OnSuccess = in.OnSuccess
	out.OnFailure = in.OnFailure
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventInvokeConfig_DestinationConfig.
func (in *EventInvokeConfig_DestinationConfig) DeepCopy() *EventInvokeConfig_DestinationConfig {
	if in == nil {
		return nil
	}
	out := new(EventInvokeConfig_DestinationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventInvokeConfig_OnFailure) DeepCopyInto(out *EventInvokeConfig_OnFailure) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventInvokeConfig_OnFailure.
func (in *EventInvokeConfig_OnFailure) DeepCopy() *EventInvokeConfig_OnFailure {
	if in == nil {
		return nil
	}
	out := new(EventInvokeConfig_OnFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventInvokeConfig_OnSuccess) DeepCopyInto(out *EventInvokeConfig_OnSuccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventInvokeConfig_OnSuccess.
func (in *EventInvokeConfig_OnSuccess) DeepCopy() *EventInvokeConfig_OnSuccess {
	if in == nil {
		return nil
	}
	out := new(EventInvokeConfig_OnSuccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMapping) DeepCopyInto(out *EventSourceMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMapping.
func (in *EventSourceMapping) DeepCopy() *EventSourceMapping {
	if in == nil {
		return nil
	}
	out := new(EventSourceMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSourceMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMappingList) DeepCopyInto(out *EventSourceMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventSourceMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMappingList.
func (in *EventSourceMappingList) DeepCopy() *EventSourceMappingList {
	if in == nil {
		return nil
	}
	out := new(EventSourceMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSourceMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMappingOutput) DeepCopyInto(out *EventSourceMappingOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMappingOutput.
func (in *EventSourceMappingOutput) DeepCopy() *EventSourceMappingOutput {
	if in == nil {
		return nil
	}
	out := new(EventSourceMappingOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMappingSpec) DeepCopyInto(out *EventSourceMappingSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.DestinationConfig = in.DestinationConfig
	out.EventSource = in.EventSource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMappingSpec.
func (in *EventSourceMappingSpec) DeepCopy() *EventSourceMappingSpec {
	if in == nil {
		return nil
	}
	out := new(EventSourceMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMappingStatus) DeepCopyInto(out *EventSourceMappingStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMappingStatus.
func (in *EventSourceMappingStatus) DeepCopy() *EventSourceMappingStatus {
	if in == nil {
		return nil
	}
	out := new(EventSourceMappingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMapping_DestinationConfig) DeepCopyInto(out *EventSourceMapping_DestinationConfig) {
	*out = *in
	out.OnFailure = in.OnFailure
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMapping_DestinationConfig.
func (in *EventSourceMapping_DestinationConfig) DeepCopy() *EventSourceMapping_DestinationConfig {
	if in == nil {
		return nil
	}
	out := new(EventSourceMapping_DestinationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceMapping_OnFailure) DeepCopyInto(out *EventSourceMapping_OnFailure) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceMapping_OnFailure.
func (in *EventSourceMapping_OnFailure) DeepCopy() *EventSourceMapping_OnFailure {
	if in == nil {
		return nil
	}
	out := new(EventSourceMapping_OnFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionOutput) DeepCopyInto(out *FunctionOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionOutput.
func (in *FunctionOutput) DeepCopy() *FunctionOutput {
	if in == nil {
		return nil
	}
	out := new(FunctionOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.DeadLetterConfig = in.DeadLetterConfig
	in.VpcConfig.DeepCopyInto(&out.VpcConfig)
	out.Code = in.Code
	in.Environment.DeepCopyInto(&out.Environment)
	if in.Layers != nil {
		in, out := &in.Layers, &out.Layers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.KmsKey = in.KmsKey
	out.TracingConfig = in.TracingConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function_Code) DeepCopyInto(out *Function_Code) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function_Code.
func (in *Function_Code) DeepCopy() *Function_Code {
	if in == nil {
		return nil
	}
	out := new(Function_Code)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function_DeadLetterConfig) DeepCopyInto(out *Function_DeadLetterConfig) {
	*out = *in
	out.Target = in.Target
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function_DeadLetterConfig.
func (in *Function_DeadLetterConfig) DeepCopy() *Function_DeadLetterConfig {
	if in == nil {
		return nil
	}
	out := new(Function_DeadLetterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function_Environment) DeepCopyInto(out *Function_Environment) {
	*out = *in
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function_Environment.
func (in *Function_Environment) DeepCopy() *Function_Environment {
	if in == nil {
		return nil
	}
	out := new(Function_Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function_TracingConfig) DeepCopyInto(out *Function_TracingConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function_TracingConfig.
func (in *Function_TracingConfig) DeepCopy() *Function_TracingConfig {
	if in == nil {
		return nil
	}
	out := new(Function_TracingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function_VpcConfig) DeepCopyInto(out *Function_VpcConfig) {
	*out = *in
	if in.SecurityGroup != nil {
		in, out := &in.SecurityGroup, &out.SecurityGroup
		*out = make([]metav1alpha1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = make([]metav1alpha1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function_VpcConfig.
func (in *Function_VpcConfig) DeepCopy() *Function_VpcConfig {
	if in == nil {
		return nil
	}
	out := new(Function_VpcConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersion) DeepCopyInto(out *LayerVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersion.
func (in *LayerVersion) DeepCopy() *LayerVersion {
	if in == nil {
		return nil
	}
	out := new(LayerVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LayerVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersionList) DeepCopyInto(out *LayerVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LayerVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersionList.
func (in *LayerVersionList) DeepCopy() *LayerVersionList {
	if in == nil {
		return nil
	}
	out := new(LayerVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LayerVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersionOutput) DeepCopyInto(out *LayerVersionOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersionOutput.
func (in *LayerVersionOutput) DeepCopy() *LayerVersionOutput {
	if in == nil {
		return nil
	}
	out := new(LayerVersionOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersionPermission) DeepCopyInto(out *LayerVersionPermission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersionPermission.
func (in *LayerVersionPermission) DeepCopy() *LayerVersionPermission {
	if in == nil {
		return nil
	}
	out := new(LayerVersionPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LayerVersionPermission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersionPermissionList) DeepCopyInto(out *LayerVersionPermissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LayerVersionPermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersionPermissionList.
func (in *LayerVersionPermissionList) DeepCopy() *LayerVersionPermissionList {
	if in == nil {
		return nil
	}
	out := new(LayerVersionPermissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LayerVersionPermissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersionPermissionOutput) DeepCopyInto(out *LayerVersionPermissionOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersionPermissionOutput.
func (in *LayerVersionPermissionOutput) DeepCopy() *LayerVersionPermissionOutput {
	if in == nil {
		return nil
	}
	out := new(LayerVersionPermissionOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersionPermissionSpec) DeepCopyInto(out *LayerVersionPermissionSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.LayerVersion = in.LayerVersion
	out.Organization = in.Organization
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersionPermissionSpec.
func (in *LayerVersionPermissionSpec) DeepCopy() *LayerVersionPermissionSpec {
	if in == nil {
		return nil
	}
	out := new(LayerVersionPermissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersionPermissionStatus) DeepCopyInto(out *LayerVersionPermissionStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersionPermissionStatus.
func (in *LayerVersionPermissionStatus) DeepCopy() *LayerVersionPermissionStatus {
	if in == nil {
		return nil
	}
	out := new(LayerVersionPermissionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersionSpec) DeepCopyInto(out *LayerVersionSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.CompatibleRuntimes != nil {
		in, out := &in.CompatibleRuntimes, &out.CompatibleRuntimes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Content = in.Content
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersionSpec.
func (in *LayerVersionSpec) DeepCopy() *LayerVersionSpec {
	if in == nil {
		return nil
	}
	out := new(LayerVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersionStatus) DeepCopyInto(out *LayerVersionStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersionStatus.
func (in *LayerVersionStatus) DeepCopy() *LayerVersionStatus {
	if in == nil {
		return nil
	}
	out := new(LayerVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LayerVersion_Content) DeepCopyInto(out *LayerVersion_Content) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LayerVersion_Content.
func (in *LayerVersion_Content) DeepCopy() *LayerVersion_Content {
	if in == nil {
		return nil
	}
	out := new(LayerVersion_Content)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Permission) DeepCopyInto(out *Permission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Permission.
func (in *Permission) DeepCopy() *Permission {
	if in == nil {
		return nil
	}
	out := new(Permission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Permission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionList) DeepCopyInto(out *PermissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Permission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionList.
func (in *PermissionList) DeepCopy() *PermissionList {
	if in == nil {
		return nil
	}
	out := new(PermissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PermissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionOutput) DeepCopyInto(out *PermissionOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionOutput.
func (in *PermissionOutput) DeepCopy() *PermissionOutput {
	if in == nil {
		return nil
	}
	out := new(PermissionOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSpec) DeepCopyInto(out *PermissionSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.Source = in.Source
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSpec.
func (in *PermissionSpec) DeepCopy() *PermissionSpec {
	if in == nil {
		return nil
	}
	out := new(PermissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionStatus) DeepCopyInto(out *PermissionStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionStatus.
func (in *PermissionStatus) DeepCopy() *PermissionStatus {
	if in == nil {
		return nil
	}
	out := new(PermissionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Version) DeepCopyInto(out *Version) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Version.
func (in *Version) DeepCopy() *Version {
	if in == nil {
		return nil
	}
	out := new(Version)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Version) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionList) DeepCopyInto(out *VersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Version, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionList.
func (in *VersionList) DeepCopy() *VersionList {
	if in == nil {
		return nil
	}
	out := new(VersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionOutput) DeepCopyInto(out *VersionOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionOutput.
func (in *VersionOutput) DeepCopy() *VersionOutput {
	if in == nil {
		return nil
	}
	out := new(VersionOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionSpec) DeepCopyInto(out *VersionSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.ProvisionedConcurrencyConfig = in.ProvisionedConcurrencyConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionSpec.
func (in *VersionSpec) DeepCopy() *VersionSpec {
	if in == nil {
		return nil
	}
	out := new(VersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionStatus) DeepCopyInto(out *VersionStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionStatus.
func (in *VersionStatus) DeepCopy() *VersionStatus {
	if in == nil {
		return nil
	}
	out := new(VersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Version_ProvisionedConcurrencyConfiguration) DeepCopyInto(out *Version_ProvisionedConcurrencyConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Version_ProvisionedConcurrencyConfiguration.
func (in *Version_ProvisionedConcurrencyConfiguration) DeepCopy() *Version_ProvisionedConcurrencyConfiguration {
	if in == nil {
		return nil
	}
	out := new(Version_ProvisionedConcurrencyConfiguration)
	in.DeepCopyInto(out)
	return out
}