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
	"github.com/awslabs/goformation/v4/cloudformation/route53"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *HostedZone) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *HostedZone) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - route53.HostedZone (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("HostedZone"),
			"Export": map[string]interface{}{
				"Name": in.Name + "Ref",
			},
		},
	}

	route53HostedZone := &route53.HostedZone{}

	route53HostedZoneHostedZoneTags := []route53.HostedZone_HostedZoneTag{}

	for _, item := range in.Spec.HostedZoneTags {
		route53HostedZoneHostedZoneTag := route53.HostedZone_HostedZoneTag{}

		if item.Key != "" {
			route53HostedZoneHostedZoneTag.Key = item.Key
		}

		if item.Value != "" {
			route53HostedZoneHostedZoneTag.Value = item.Value
		}

	}

	if len(route53HostedZoneHostedZoneTags) > 0 {
		route53HostedZone.HostedZoneTags = route53HostedZoneHostedZoneTags
	}
	if in.Spec.Name != "" {
		route53HostedZone.Name = in.Spec.Name
	}

	if !reflect.DeepEqual(in.Spec.QueryLoggingConfig, HostedZone_QueryLoggingConfig{}) {
		route53HostedZoneQueryLoggingConfig := route53.HostedZone_QueryLoggingConfig{}

		// TODO(christopherhein) move these to a defaulter
		route53HostedZoneQueryLoggingConfigCloudWatchLogsLogGroupRefItem := in.Spec.QueryLoggingConfig.CloudWatchLogsLogGroupRef.DeepCopy()

		if route53HostedZoneQueryLoggingConfigCloudWatchLogsLogGroupRefItem.ObjectRef.Namespace == "" {
			route53HostedZoneQueryLoggingConfigCloudWatchLogsLogGroupRefItem.ObjectRef.Namespace = in.Namespace
		}

		in.Spec.QueryLoggingConfig.CloudWatchLogsLogGroupRef = *route53HostedZoneQueryLoggingConfigCloudWatchLogsLogGroupRefItem
		cloudWatchLogsLogGroupArn, err := in.Spec.QueryLoggingConfig.CloudWatchLogsLogGroupRef.String(client)
		if err != nil {
			return "", err
		}

		if cloudWatchLogsLogGroupArn != "" {
			route53HostedZoneQueryLoggingConfig.CloudWatchLogsLogGroupArn = cloudWatchLogsLogGroupArn
		}

		route53HostedZone.QueryLoggingConfig = &route53HostedZoneQueryLoggingConfig
	}

	route53HostedZoneVPCs := []route53.HostedZone_VPC{}

	for _, item := range in.Spec.VPCs {
		route53HostedZoneVPC := route53.HostedZone_VPC{}

		// TODO(christopherhein) move these to a defaulter
		route53HostedZoneVPCVPCRefItem := item.VPCRef.DeepCopy()

		if route53HostedZoneVPCVPCRefItem.ObjectRef.Namespace == "" {
			route53HostedZoneVPCVPCRefItem.ObjectRef.Namespace = in.Namespace
		}

		item.VPCRef = *route53HostedZoneVPCVPCRefItem
		vPCId, err := item.VPCRef.String(client)
		if err != nil {
			return "", err
		}

		if vPCId != "" {
			route53HostedZoneVPC.VPCId = vPCId
		}

		if item.VPCRegion != "" {
			route53HostedZoneVPC.VPCRegion = item.VPCRegion
		}

	}

	if len(route53HostedZoneVPCs) > 0 {
		route53HostedZone.VPCs = route53HostedZoneVPCs
	}
	if !reflect.DeepEqual(in.Spec.HostedZoneConfig, HostedZone_HostedZoneConfig{}) {
		route53HostedZoneHostedZoneConfig := route53.HostedZone_HostedZoneConfig{}

		if in.Spec.HostedZoneConfig.Comment != "" {
			route53HostedZoneHostedZoneConfig.Comment = in.Spec.HostedZoneConfig.Comment
		}

		route53HostedZone.HostedZoneConfig = &route53HostedZoneHostedZoneConfig
	}

	template.Resources = map[string]cloudformation.Resource{
		"HostedZone": route53HostedZone,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *HostedZone) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *HostedZone) GenerateStackName() string {
	return strings.Join([]string{"route53", "hostedzone", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *HostedZone) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *HostedZone) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *HostedZone) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *HostedZone) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *HostedZone) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *HostedZone) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *HostedZone) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *HostedZone) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *HostedZone) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *HostedZone) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
