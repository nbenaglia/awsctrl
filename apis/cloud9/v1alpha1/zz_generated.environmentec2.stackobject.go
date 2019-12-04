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
	"strings"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	cfnencoder "go.awsctrl.io/manager/encoding/cloudformation"

	goformation "github.com/awslabs/goformation/v3/cloudformation"
	"github.com/awslabs/goformation/v3/cloudformation/cloud9"
	"github.com/awslabs/goformation/v3/intrinsics"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *EnvironmentEC2) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *EnvironmentEC2) GetTemplate() (string, error) {
	template := goformation.NewTemplate()

	template.Description = "AWS Controller - cloud9.EnvironmentEC2 (ac-{TODO})"

	// GenerateTemplateFunctions
	cloud9EnvironmentEC2 := &cloud9.EnvironmentEC2{}

	if in.Spec.Name != "" {
		cloud9EnvironmentEC2.Name = in.Spec.Name
	}

	cloud9EnvironmentEC2Repositories := []cloud9.EnvironmentEC2_Repository{}

	for _, item := range in.Spec.Repositories {
		cloud9EnvironmentEC2Repository := cloud9.EnvironmentEC2_Repository{}

		if item.PathComponent != "" {
			cloud9EnvironmentEC2Repository.PathComponent = item.PathComponent
		}

		if item.RepositoryUrl != "" {
			cloud9EnvironmentEC2Repository.RepositoryUrl = item.RepositoryUrl
		}

	}

	if len(cloud9EnvironmentEC2Repositories) > 0 {
		cloud9EnvironmentEC2.Repositories = cloud9EnvironmentEC2Repositories
	}
	if in.Spec.OwnerArn != "" {
		cloud9EnvironmentEC2.OwnerArn = in.Spec.OwnerArn
	}

	if in.Spec.Description != "" {
		cloud9EnvironmentEC2.Description = in.Spec.Description
	}

	if in.Spec.SubnetId != "" {
		cloud9EnvironmentEC2.SubnetId = in.Spec.SubnetId
	}

	if in.Spec.InstanceType != "" {
		cloud9EnvironmentEC2.InstanceType = in.Spec.InstanceType
	}

	template.Resources = map[string]goformation.Resource{
		"EnvironmentEC2": cloud9EnvironmentEC2,
	}

	json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *EnvironmentEC2) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *EnvironmentEC2) GenerateStackName() string {
	return strings.Join([]string{"cloud9", "environmentec2", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *EnvironmentEC2) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *EnvironmentEC2) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *EnvironmentEC2) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *EnvironmentEC2) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *EnvironmentEC2) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *EnvironmentEC2) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *EnvironmentEC2) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *EnvironmentEC2) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *EnvironmentEC2) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *EnvironmentEC2) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}