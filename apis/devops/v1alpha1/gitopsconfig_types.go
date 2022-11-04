/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GitopsConfigSpec defines the desired state of GitopsConfig
type GitopsConfigSpec struct {
	// Overrides the default GitOps configuration with custom configuration for the specified app(s).
	Configutaion []Configuration `json:"config"`
}

// CommitMode describes how Turbonomic events will be processed.
// +kubebuilder:validation:Enum=direct;pr
type CommitMode string

const (
	// Actions will produce commit directly within the underlying repository without creating a pull/merge request
	DirectCommit CommitMode = "direct"

	// Actions will result in a pull/merge request being creating within the underlying repository
	PullRequest CommitMode = "pr"
)

type Configuration struct {
	// Specifies the GitOps commit mode.
	// Valid values are:
	// - "direct": actions will produce commit directly within the underlying repository without creating a pull/merge request;
	// - "pr": actions will result in a pull/merge request being creating within the underlying repository
	// +kubebuilder:default:=pr
	CommitMode CommitMode `json:"commitMode,omitempty"`
	// Specifies the credentials for the underlying repository (CURRENTLY UNSUPPORTED)
	// +optional
	Credentials Credentials `json:"credentials,omitempty"`
	// Specifies the applications that the commit mode should apply to
	// +kubebuilder:validation:MinItems:=1
	Applications []string `json:"apps"`
}

type Credentials struct {
	// Specifies the email address of the user from which commits/PRs will be created
	// +kubebuilder:validation:Format=email
	Email string `json:"email"`
	// Specifies the name of the secret containing credentials for the repository
	SecretName string `json:"secretName"`
	// Specifies the namespace in which the secret containing the credentials exists
	SecretNamespace string `json:"secretNamespace"`
	// Specifies the username from which commits/PRs will be created by
	Username string `json:"username"`
}

// GitopsConfigStatus defines the observed state of GitopsConfig
type GitopsConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GitopsConfig is the Schema for the gitopsconfigs API
type GitopsConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitopsConfigSpec   `json:"spec,omitempty"`
	Status GitopsConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GitopsConfigList contains a list of GitopsConfig
type GitopsConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitopsConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitopsConfig{}, &GitopsConfigList{})
}
