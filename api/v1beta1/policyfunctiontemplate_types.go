/*
Copyright 2020 Stark & Wayne LLC.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PolicyFunctionTemplateSpec defines the desired state of PolicyFunctionTemplate
type PolicyFunctionTemplateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// AdmissionType indicates it should create a ValidationWebhookConfiguration ("validation")
	// rather than the default MutatingWebhookConfiguration ("mutating")
	// +kubebuilder:validation:Enum=mutating;validating
	AdmissionType string `json:"admission_type,omitempty"`

	// Language helps the webhook runner to decide how to interpret/run the script
	Language string `json:"language"`

	// LanguageWebhookService indicates which service in current namespace operates webhook
	// that implements the Language requested.
	LanguageWebhookService string `json:"language_webhook_service"`

	// HandlerScript contains a handler(admission_review, admission_resp) function
	// in the Language.
	HandlerScript string `json:"handler_script"`
}

// PolicyFunctionTemplateStatus defines the observed state of PolicyFunctionTemplate
type PolicyFunctionTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// PolicyFunctionTemplate is the Schema for the policyfunctiontemplates API
type PolicyFunctionTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PolicyFunctionTemplateSpec   `json:"spec,omitempty"`
	Status PolicyFunctionTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyFunctionTemplateList contains a list of PolicyFunctionTemplate
type PolicyFunctionTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyFunctionTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PolicyFunctionTemplate{}, &PolicyFunctionTemplateList{})
}
