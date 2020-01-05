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

package e2e_test

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"
	iamv1alpha1 "go.awsctrl.io/manager/apis/iam/v1alpha1"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
)

// RunAccountSpecs allows all usertogroupaddition E2E tests to run
var _ = Describe("Run IAM UserToGroupAddition Controller", func() {

	Context("Without UserToGroupAddition{} existing", func() {

		It("Should create iam.UserToGroupAddition{}", func() {
			var stackID string
			var stackName string
			var stack *cloudformationv1alpha1.Stack
			k8sclient := k8smanager.GetClient()
			Expect(k8sclient).ToNot(BeNil())

			group := &iamv1alpha1.Group{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "sample-group-",
					Namespace:    podnamespace,
				},
				Spec: iamv1alpha1.GroupSpec{
					CloudFormationMeta: metav1alpha1.CloudFormationMeta{
						Capabilities: []*string{
							&capabilityIAM,
						},
					},
					GroupName: "sample-group",
					ManagedPolicyArns: []string{
						"arn:aws:iam::aws:policy/AdministratorAccess",
					},
				},
			}
			By("Creating new IAM Group")
			Expect(k8sclient.Create(context.Background(), group)).Should(Succeed())

			user := &iamv1alpha1.User{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "sample-user-",
					Namespace:    podnamespace,
				},
				Spec: iamv1alpha1.UserSpec{
					CloudFormationMeta: metav1alpha1.CloudFormationMeta{
						Capabilities: []*string{
							&capabilityIAM,
						},
					},
					UserName: "christopherhein",
					ManagedPolicyArns: []string{
						"arn:aws:iam::aws:policy/AdministratorAccess",
					},
				},
			}
			By("Creating new IAM User")
			Expect(k8sclient.Create(context.Background(), user)).Should(Succeed())

			instance := &iamv1alpha1.UserToGroupAddition{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "sample-usertogroupaddition-",
					Namespace:    podnamespace,
				},
				Spec: iamv1alpha1.UserToGroupAdditionSpec{
					GroupName: "sample-group",
					Users: []string{
						"sample-user",
					},
				},
			}
			By("Creating new IAM UserToGroupAddition")
			Expect(k8sclient.Create(context.Background(), instance)).Should(Succeed())

			key := types.NamespacedName{
				Name:      instance.GetName(),
				Namespace: podnamespace,
			}

			By("Expecting CreateComplete")
			Eventually(func() bool {
				By("Getting latest IAM UserToGroupAddition")
				instance = &iamv1alpha1.UserToGroupAddition{}
				err := k8sclient.Get(context.Background(), key, instance)
				if err != nil {
					return false
				}

				stackID = instance.GetStackID()
				stackName = instance.GetStackName()

				return instance.Status.Status == metav1alpha1.CreateCompleteStatus ||
					(os.Getenv("USE_AWS_CLIENT") != "true" && instance.Status.Status != "")
			}, timeout, interval).Should(BeTrue())

			By("Checking object OwnerShip")
			Eventually(func() bool {
				stackkey := types.NamespacedName{
					Name:      stackName,
					Namespace: key.Namespace,
				}

				stack = &cloudformationv1alpha1.Stack{}
				Expect(k8sclient.Get(context.Background(), stackkey, stack)).Should(Succeed())

				expectedOwnerReference := v1.OwnerReference{
					Kind:       instance.Kind,
					APIVersion: instance.APIVersion,
					UID:        instance.UID,
					Name:       instance.Name,
				}

				ownerrefs := stack.GetOwnerReferences()
				Expect(len(ownerrefs)).To(Equal(1))

				return ownerrefs[0].Name == expectedOwnerReference.Name
			}, timeout, interval).Should(BeTrue())

			By("Deleting IAM UserToGroupAddition")
			Expect(k8sclient.Delete(context.Background(), instance)).Should(Succeed())

			By("Deleting IAM UserToGroupAddition Stack")
			Expect(k8sclient.Delete(context.Background(), stack)).Should(Succeed())

			By("Expecting metav1alpha1.DeleteCompleteStatus")
			Eventually(func() bool {
				if os.Getenv("USE_AWS_CLIENT") != "true" {
					return true
				}

				output, err := awsclient.GetClient("us-west-2").DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackID)})
				Expect(err).To(BeNil())
				stackoutput := output.Stacks[0].StackStatus
				return *stackoutput == "DELETE_COMPLETE"
			}, timeout, interval).Should(BeTrue())
		})
	})
})
