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

// BorrowWithinCohortApplyConfiguration represents a declarative configuration of the BorrowWithinCohort type for use
// with apply.
type BorrowWithinCohortApplyConfiguration struct {
	Policy               *v1beta1.BorrowWithinCohortPolicy `json:"policy,omitempty"`
	MaxPriorityThreshold *int32                            `json:"maxPriorityThreshold,omitempty"`
}

// BorrowWithinCohortApplyConfiguration constructs a declarative configuration of the BorrowWithinCohort type for use with
// apply.
func BorrowWithinCohort() *BorrowWithinCohortApplyConfiguration {
	return &BorrowWithinCohortApplyConfiguration{}
}

// WithPolicy sets the Policy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Policy field is set to the value of the last call.
func (b *BorrowWithinCohortApplyConfiguration) WithPolicy(value v1beta1.BorrowWithinCohortPolicy) *BorrowWithinCohortApplyConfiguration {
	b.Policy = &value
	return b
}

// WithMaxPriorityThreshold sets the MaxPriorityThreshold field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxPriorityThreshold field is set to the value of the last call.
func (b *BorrowWithinCohortApplyConfiguration) WithMaxPriorityThreshold(value int32) *BorrowWithinCohortApplyConfiguration {
	b.MaxPriorityThreshold = &value
	return b
}
