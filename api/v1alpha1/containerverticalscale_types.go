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

// ResourceConstraints defines the resize constraint for resource like CPU or Memory
type ResourceConstraint struct {
	Max               string `json:"max,omitempty"`
	Min               string `json:"min,omitempty"`
	RecommendAboveMax bool   `json:"recommendAboveMax,omitempty"`
	RecommendBelowMin bool   `json:"recommendBelowMin,omitempty"`
}

// resourceConstraints defines the resource constraints for CPU and Memory
type resourceConstraints struct {
	CPU    ResourceConstraint `json:"cpu,omitempty"`
	Memory ResourceConstraint `json:"memory,omitempty"`
}

// ContainerVerticalScaleSpec defines the desired state of ContainerVerticalScale
type ContainerVerticalScaleSpec struct {
	Settings struct {
		Limits     resourceConstraints `json:"limits,omitempty"`
		Requests   resourceConstraints `json:"requests,omitempty"`
		Increments struct {
			CPU    string `json:"cpu,omitempty"`
			Memory string `json:"memory,omitempty"`
		} `json:"increments,omitempty"`
		ObservationPeriod struct {
			Min string `json:"min,omitempty"`
			Max string `json:"max,omitempty"`
		} `json:"observationPeriod,omitempty"`
		RateOfResize   string `json:"rateOfResize,omitempty"`
		Aggressiveness string `json:"aggressiveness,omitempty"`
	} `json:"settings,omitempty"`

	// The behavior of vertical resize actions
	// +kubebuilder:default:={resize:Manual}
	// +optional
	Behavior ActionBehavior `json:"behavior,omitempty"`
}

// ContainerVerticalScaleStatus defines the observed state of ContainerVerticalScale
type ContainerVerticalScaleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ContainerVerticalScale is the Schema for the containerverticalscales API
type ContainerVerticalScale struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerVerticalScaleSpec   `json:"spec,omitempty"`
	Status ContainerVerticalScaleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ContainerVerticalScaleList contains a list of ContainerVerticalScale
type ContainerVerticalScaleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerVerticalScale `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerVerticalScale{}, &ContainerVerticalScaleList{})
}
