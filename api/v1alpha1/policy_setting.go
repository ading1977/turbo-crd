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
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type PolicySettingName string

const (
	ResponseTime PolicySettingName = "ResponseTime"
	Transaction  PolicySettingName = "Transaction"
)

type PolicySettingValue *v1.JSON

type PolicySetting struct {
	// The name of the policy setting
	// +kubebuilder:validation:Enum=ResponseTime;Transaction
	Name PolicySettingName `json:"name"`

	// The value of the policy setting
	Value PolicySettingValue `json:"value"`
}
