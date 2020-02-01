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
	"fmt"
	"strings"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	cfnencoder "go.awsctrl.io/manager/encoding/cloudformation"

	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/lambda"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *Permission) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *Permission) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - lambda.Permission (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("Permission"),
			"Export": map[string]interface{}{
				"Name": in.Name + "Ref",
			},
		},
	}

	lambdaPermission := &lambda.Permission{}

	if in.Spec.FunctionName != "" {
		lambdaPermission.FunctionName = in.Spec.FunctionName
	}

	if in.Spec.Principal != "" {
		lambdaPermission.Principal = in.Spec.Principal
	}

	if in.Spec.SourceAccount != "" {
		lambdaPermission.SourceAccount = in.Spec.SourceAccount
	}

	// TODO(christopherhein) move these to a defaulter
	lambdaPermissionSourceRefItem := in.Spec.SourceRef.DeepCopy()

	if lambdaPermissionSourceRefItem.ObjectRef.Namespace == "" {
		lambdaPermissionSourceRefItem.ObjectRef.Namespace = in.Namespace
	}

	in.Spec.SourceRef = *lambdaPermissionSourceRefItem
	sourceArn, err := in.Spec.SourceRef.String(client)
	if err != nil {
		return "", err
	}

	if sourceArn != "" {
		lambdaPermission.SourceArn = sourceArn
	}

	if in.Spec.Action != "" {
		lambdaPermission.Action = in.Spec.Action
	}

	if in.Spec.EventSourceToken != "" {
		lambdaPermission.EventSourceToken = in.Spec.EventSourceToken
	}

	template.Resources = map[string]cloudformation.Resource{
		"Permission": lambdaPermission,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *Permission) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *Permission) GenerateStackName() string {
	return strings.Join([]string{"lambda", "permission", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *Permission) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *Permission) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *Permission) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *Permission) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *Permission) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *Permission) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *Permission) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *Permission) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *Permission) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *Permission) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
