/*
Copyright The Kubernetes Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "sigs.k8s.io/kueue/apis/kueue/v1beta1"
)

// KubeConfigApplyConfiguration represents a declarative configuration of the KubeConfig type for use
// with apply.
type KubeConfigApplyConfiguration struct {
	Location     *string               `json:"location,omitempty"`
	LocationType *v1beta1.LocationType `json:"locationType,omitempty"`
}

// KubeConfigApplyConfiguration constructs a declarative configuration of the KubeConfig type for use with
// apply.
func KubeConfig() *KubeConfigApplyConfiguration {
	return &KubeConfigApplyConfiguration{}
}

// WithLocation sets the Location field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Location field is set to the value of the last call.
func (b *KubeConfigApplyConfiguration) WithLocation(value string) *KubeConfigApplyConfiguration {
	b.Location = &value
	return b
}

// WithLocationType sets the LocationType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LocationType field is set to the value of the last call.
func (b *KubeConfigApplyConfiguration) WithLocationType(value v1beta1.LocationType) *KubeConfigApplyConfiguration {
	b.LocationType = &value
	return b
}
