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
	"reflect"
	"strconv"
	"strings"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	cfnencoder "go.awsctrl.io/manager/encoding/cloudformation"

	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/lambda"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *Alias) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *Alias) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - lambda.Alias (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("Alias"),
		},
	}

	lambdaAlias := &lambda.Alias{}

	if !reflect.DeepEqual(in.Spec.ProvisionedConcurrencyConfig, Alias_ProvisionedConcurrencyConfiguration{}) {
		lambdaAliasProvisionedConcurrencyConfiguration := lambda.Alias_ProvisionedConcurrencyConfiguration{}

		if in.Spec.ProvisionedConcurrencyConfig.ProvisionedConcurrentExecutions != lambdaAliasProvisionedConcurrencyConfiguration.ProvisionedConcurrentExecutions {
			lambdaAliasProvisionedConcurrencyConfiguration.ProvisionedConcurrentExecutions = in.Spec.ProvisionedConcurrencyConfig.ProvisionedConcurrentExecutions
		}

		lambdaAlias.ProvisionedConcurrencyConfig = &lambdaAliasProvisionedConcurrencyConfiguration
	}

	if !reflect.DeepEqual(in.Spec.RoutingConfig, Alias_AliasRoutingConfiguration{}) {
		lambdaAliasAliasRoutingConfiguration := lambda.Alias_AliasRoutingConfiguration{}

		lambdaAliasAliasRoutingConfigurationAdditionalVersionWeights := []lambda.Alias_VersionWeight{}

		for _, item := range in.Spec.RoutingConfig.AdditionalVersionWeights {
			lambdaAliasAliasRoutingConfigurationVersionWeight := lambda.Alias_VersionWeight{}

			if item.FunctionVersion != "" {
				lambdaAliasAliasRoutingConfigurationVersionWeight.FunctionVersion = item.FunctionVersion
			}

			if f, _ := strconv.ParseFloat(item.FunctionWeight, 64); f != lambdaAliasAliasRoutingConfigurationVersionWeight.FunctionWeight {
				f, _ := strconv.ParseFloat(item.FunctionWeight, 64)
				lambdaAliasAliasRoutingConfigurationVersionWeight.FunctionWeight = f
			}

		}

		if len(lambdaAliasAliasRoutingConfigurationAdditionalVersionWeights) > 0 {
			lambdaAliasAliasRoutingConfiguration.AdditionalVersionWeights = lambdaAliasAliasRoutingConfigurationAdditionalVersionWeights
		}

		lambdaAlias.RoutingConfig = &lambdaAliasAliasRoutingConfiguration
	}

	if in.Spec.Description != "" {
		lambdaAlias.Description = in.Spec.Description
	}

	if in.Spec.FunctionName != "" {
		lambdaAlias.FunctionName = in.Spec.FunctionName
	}

	if in.Spec.FunctionVersion != "" {
		lambdaAlias.FunctionVersion = in.Spec.FunctionVersion
	}

	if in.Spec.Name != "" {
		lambdaAlias.Name = in.Spec.Name
	}

	template.Resources = map[string]cloudformation.Resource{
		"Alias": lambdaAlias,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *Alias) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *Alias) GenerateStackName() string {
	return strings.Join([]string{"lambda", "alias", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *Alias) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *Alias) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *Alias) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *Alias) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *Alias) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *Alias) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *Alias) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *Alias) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *Alias) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *Alias) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}