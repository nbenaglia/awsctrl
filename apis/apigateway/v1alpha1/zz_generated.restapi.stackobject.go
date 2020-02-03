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
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	cfnencoder "go.awsctrl.io/manager/encoding/cloudformation"

	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/apigateway"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *RestApi) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *RestApi) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - apigateway.RestApi (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("RestApi"),
			"Export": map[string]interface{}{
				"Name": in.Name + "Ref",
			},
		},
		"RootResourceId": map[string]interface{}{
			"Value":  cloudformation.GetAtt("RestApi", "RootResourceId"),
			"Export": map[string]interface{}{"Name": in.Name + "RootResourceId"},
		},
	}

	apigatewayRestApi := &apigateway.RestApi{}

	if in.Spec.ApiKeySourceType != "" {
		apigatewayRestApi.ApiKeySourceType = in.Spec.ApiKeySourceType
	}

	if len(in.Spec.BinaryMediaTypes) > 0 {
		apigatewayRestApi.BinaryMediaTypes = in.Spec.BinaryMediaTypes
	}

	if in.Spec.Body != "" {
		apigatewayRestApiJSON := make(map[string]interface{})
		err := json.Unmarshal([]byte(in.Spec.Body), &apigatewayRestApiJSON)
		if err != nil {
			return "", err
		}
		apigatewayRestApi.Body = apigatewayRestApiJSON
	}

	if !reflect.DeepEqual(in.Spec.BodyS3Location, RestApi_S3Location{}) {
		apigatewayRestApiS3Location := apigateway.RestApi_S3Location{}

		if in.Spec.BodyS3Location.Bucket != "" {
			apigatewayRestApiS3Location.Bucket = in.Spec.BodyS3Location.Bucket
		}

		if in.Spec.BodyS3Location.ETag != "" {
			apigatewayRestApiS3Location.ETag = in.Spec.BodyS3Location.ETag
		}

		if in.Spec.BodyS3Location.Key != "" {
			apigatewayRestApiS3Location.Key = in.Spec.BodyS3Location.Key
		}

		if in.Spec.BodyS3Location.Version != "" {
			apigatewayRestApiS3Location.Version = in.Spec.BodyS3Location.Version
		}

		apigatewayRestApi.BodyS3Location = &apigatewayRestApiS3Location
	}

	if in.Spec.CloneFrom != "" {
		apigatewayRestApi.CloneFrom = in.Spec.CloneFrom
	}

	if in.Spec.Description != "" {
		apigatewayRestApi.Description = in.Spec.Description
	}

	if !reflect.DeepEqual(in.Spec.EndpointConfiguration, RestApi_EndpointConfiguration{}) {
		apigatewayRestApiEndpointConfiguration := apigateway.RestApi_EndpointConfiguration{}

		if len(in.Spec.EndpointConfiguration.Types) > 0 {
			apigatewayRestApiEndpointConfiguration.Types = in.Spec.EndpointConfiguration.Types
		}

		if len(in.Spec.EndpointConfiguration.VpcEndpointRefs) > 0 {
			apigatewayRestApiEndpointConfigurationVpcEndpointRefs := []string{}

			for _, item := range in.Spec.EndpointConfiguration.VpcEndpointRefs {
				apigatewayRestApiEndpointConfigurationVpcEndpointRefsItem := item.DeepCopy()

				if apigatewayRestApiEndpointConfigurationVpcEndpointRefsItem.ObjectRef.Namespace == "" {
					apigatewayRestApiEndpointConfigurationVpcEndpointRefsItem.ObjectRef.Namespace = in.Namespace
				}

				vpcEndpointIds, err := apigatewayRestApiEndpointConfigurationVpcEndpointRefsItem.String(client)
				if err != nil {
					return "", err
				}

				if vpcEndpointIds != "" {
					apigatewayRestApiEndpointConfigurationVpcEndpointRefs = append(apigatewayRestApiEndpointConfigurationVpcEndpointRefs, vpcEndpointIds)
				}
			}

			apigatewayRestApiEndpointConfiguration.VpcEndpointIds = apigatewayRestApiEndpointConfigurationVpcEndpointRefs
		}

		apigatewayRestApi.EndpointConfiguration = &apigatewayRestApiEndpointConfiguration
	}

	if in.Spec.FailOnWarnings || !in.Spec.FailOnWarnings {
		apigatewayRestApi.FailOnWarnings = in.Spec.FailOnWarnings
	}

	if in.Spec.MinimumCompressionSize != apigatewayRestApi.MinimumCompressionSize {
		apigatewayRestApi.MinimumCompressionSize = in.Spec.MinimumCompressionSize
	}

	if in.Spec.Name != "" {
		apigatewayRestApi.Name = in.Spec.Name
	}

	if !reflect.DeepEqual(in.Spec.Parameters, map[string]string{}) {
		apigatewayRestApi.Parameters = in.Spec.Parameters
	}

	if in.Spec.Policy != "" {
		apigatewayRestApiJSON := make(map[string]interface{})
		err := json.Unmarshal([]byte(in.Spec.Policy), &apigatewayRestApiJSON)
		if err != nil {
			return "", err
		}
		apigatewayRestApi.Policy = apigatewayRestApiJSON
	}

	// TODO(christopherhein): implement tags this could be easy now that I have the mechanims of nested objects

	template.Resources = map[string]cloudformation.Resource{
		"RestApi": apigatewayRestApi,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *RestApi) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *RestApi) GenerateStackName() string {
	return strings.Join([]string{"apigateway", "restapi", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *RestApi) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *RestApi) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *RestApi) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *RestApi) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *RestApi) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *RestApi) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *RestApi) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *RestApi) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *RestApi) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *RestApi) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
