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
	"strings"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	cfnencoder "go.awsctrl.io/manager/encoding/cloudformation"

	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/lambda"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *EventInvokeConfig) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *EventInvokeConfig) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - lambda.EventInvokeConfig (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("EventInvokeConfig"),
			"Export": map[string]interface{}{
				"Name": in.Name + "Ref",
			},
		},
	}

	lambdaEventInvokeConfig := &lambda.EventInvokeConfig{}

	if !reflect.DeepEqual(in.Spec.DestinationConfig, EventInvokeConfig_DestinationConfig{}) {
		lambdaEventInvokeConfigDestinationConfig := lambda.EventInvokeConfig_DestinationConfig{}

		if !reflect.DeepEqual(in.Spec.DestinationConfig.OnFailure, EventInvokeConfig_OnFailure{}) {
			lambdaEventInvokeConfigDestinationConfigOnFailure := lambda.EventInvokeConfig_OnFailure{}

			if in.Spec.DestinationConfig.OnFailure.Destination != "" {
				lambdaEventInvokeConfigDestinationConfigOnFailure.Destination = in.Spec.DestinationConfig.OnFailure.Destination
			}

			lambdaEventInvokeConfigDestinationConfig.OnFailure = &lambdaEventInvokeConfigDestinationConfigOnFailure
		}

		if !reflect.DeepEqual(in.Spec.DestinationConfig.OnSuccess, EventInvokeConfig_OnSuccess{}) {
			lambdaEventInvokeConfigDestinationConfigOnSuccess := lambda.EventInvokeConfig_OnSuccess{}

			if in.Spec.DestinationConfig.OnSuccess.Destination != "" {
				lambdaEventInvokeConfigDestinationConfigOnSuccess.Destination = in.Spec.DestinationConfig.OnSuccess.Destination
			}

			lambdaEventInvokeConfigDestinationConfig.OnSuccess = &lambdaEventInvokeConfigDestinationConfigOnSuccess
		}

		lambdaEventInvokeConfig.DestinationConfig = &lambdaEventInvokeConfigDestinationConfig
	}

	if in.Spec.FunctionName != "" {
		lambdaEventInvokeConfig.FunctionName = in.Spec.FunctionName
	}

	if in.Spec.MaximumEventAgeInSeconds != lambdaEventInvokeConfig.MaximumEventAgeInSeconds {
		lambdaEventInvokeConfig.MaximumEventAgeInSeconds = in.Spec.MaximumEventAgeInSeconds
	}

	if in.Spec.MaximumRetryAttempts != lambdaEventInvokeConfig.MaximumRetryAttempts {
		lambdaEventInvokeConfig.MaximumRetryAttempts = in.Spec.MaximumRetryAttempts
	}

	if in.Spec.Qualifier != "" {
		lambdaEventInvokeConfig.Qualifier = in.Spec.Qualifier
	}

	template.Resources = map[string]cloudformation.Resource{
		"EventInvokeConfig": lambdaEventInvokeConfig,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *EventInvokeConfig) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *EventInvokeConfig) GenerateStackName() string {
	return strings.Join([]string{"lambda", "eventinvokeconfig", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *EventInvokeConfig) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *EventInvokeConfig) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *EventInvokeConfig) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *EventInvokeConfig) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *EventInvokeConfig) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *EventInvokeConfig) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *EventInvokeConfig) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *EventInvokeConfig) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *EventInvokeConfig) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *EventInvokeConfig) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
