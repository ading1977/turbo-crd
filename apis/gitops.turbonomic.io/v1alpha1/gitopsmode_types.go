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

// GitOpsModeSpec defines the desired state of GitOpsMode
type GitOpsModeSpec struct {
	// Overrides the default GitOps configuration with custom configuration for the specified app(s).
	Configutaion []GitOpsConfiguration `json:"configuration"`
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

type GitOpsConfiguration struct {
	// Specifies the GitOps commit mode.
	// Valid values are:
	// - "direct": actions will produce commit directly within the underlying repository without creating a pull/merge request;
	// - "pr": actions will result in a pull/merge request being creating within the underlying repository
	CommitMode CommitMode `json:"commitMode"`
	// Specifies the credentials for the underlying repository (CURRENTLY UNSUPPORTED)
	// +optional
	Credentials GitOpsCredentials `json:"credentials,omitempty"`
	// Specifies the applications that the commit mode should apply to
	Applications []string `json:"applications"`
}

type GitOpsCredentials struct {
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

// GitOpsModeStatus defines the observed state of GitOpsMode
type GitOpsModeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GitOpsMode is the Schema for the gitopsmodes API
type GitOpsMode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitOpsModeSpec   `json:"spec,omitempty"`
	Status GitOpsModeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GitOpsModeList contains a list of GitOpsMode
type GitOpsModeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitOpsMode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitOpsMode{}, &GitOpsModeList{})
}
