/*
Copyright © 2019 AWS Controller authors

Licensed under the Apache License, Version 2.0 (the &#34;License&#34;);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an &#34;AS IS&#34; BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LayerVersionPermissionSpec defines the desired state of LayerVersionPermission
type LayerVersionPermissionSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Action http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-action
	Action string `json:"action,omitempty" cloudformation:"Action,Parameter"`

	// LayerVersionRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-layerversionarn
	LayerVersionRef metav1alpha1.ObjectReference `json:"layerVersionRef,omitempty" cloudformation:"LayerVersionArn,Parameter"`

	// OrganizationRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-organizationid
	OrganizationRef metav1alpha1.ObjectReference `json:"organizationRef,omitempty" cloudformation:"OrganizationId,Parameter"`

	// Principal http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-principal
	Principal string `json:"principal,omitempty" cloudformation:"Principal,Parameter"`
}

// LayerVersionPermissionStatus defines the observed state of LayerVersionPermission
type LayerVersionPermissionStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// LayerVersionPermissionOutput defines the stack outputs
type LayerVersionPermissionOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;lambda
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// LayerVersionPermission is the Schema for the lambda LayerVersionPermission API
type LayerVersionPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LayerVersionPermissionSpec   `json:"spec,omitempty"`
	Status LayerVersionPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LayerVersionPermissionList contains a list of Account
type LayerVersionPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LayerVersionPermission `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LayerVersionPermission{}, &LayerVersionPermissionList{})
}
