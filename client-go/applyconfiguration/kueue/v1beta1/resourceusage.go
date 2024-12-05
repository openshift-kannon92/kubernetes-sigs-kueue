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
	v1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// ResourceUsageApplyConfiguration represents a declarative configuration of the ResourceUsage type for use
// with apply.
type ResourceUsageApplyConfiguration struct {
	Name     *v1.ResourceName   `json:"name,omitempty"`
	Total    *resource.Quantity `json:"total,omitempty"`
	Borrowed *resource.Quantity `json:"borrowed,omitempty"`
}

// ResourceUsageApplyConfiguration constructs a declarative configuration of the ResourceUsage type for use with
// apply.
func ResourceUsage() *ResourceUsageApplyConfiguration {
	return &ResourceUsageApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceUsageApplyConfiguration) WithName(value v1.ResourceName) *ResourceUsageApplyConfiguration {
	b.Name = &value
	return b
}

// WithTotal sets the Total field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Total field is set to the value of the last call.
func (b *ResourceUsageApplyConfiguration) WithTotal(value resource.Quantity) *ResourceUsageApplyConfiguration {
	b.Total = &value
	return b
}

// WithBorrowed sets the Borrowed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Borrowed field is set to the value of the last call.
func (b *ResourceUsageApplyConfiguration) WithBorrowed(value resource.Quantity) *ResourceUsageApplyConfiguration {
	b.Borrowed = &value
	return b
}
