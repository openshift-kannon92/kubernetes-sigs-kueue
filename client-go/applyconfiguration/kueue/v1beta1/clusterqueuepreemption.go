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

// ClusterQueuePreemptionApplyConfiguration represents a declarative configuration of the ClusterQueuePreemption type for use
// with apply.
type ClusterQueuePreemptionApplyConfiguration struct {
	ReclaimWithinCohort *v1beta1.PreemptionPolicy             `json:"reclaimWithinCohort,omitempty"`
	BorrowWithinCohort  *BorrowWithinCohortApplyConfiguration `json:"borrowWithinCohort,omitempty"`
	WithinClusterQueue  *v1beta1.PreemptionPolicy             `json:"withinClusterQueue,omitempty"`
}

// ClusterQueuePreemptionApplyConfiguration constructs a declarative configuration of the ClusterQueuePreemption type for use with
// apply.
func ClusterQueuePreemption() *ClusterQueuePreemptionApplyConfiguration {
	return &ClusterQueuePreemptionApplyConfiguration{}
}

// WithReclaimWithinCohort sets the ReclaimWithinCohort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReclaimWithinCohort field is set to the value of the last call.
func (b *ClusterQueuePreemptionApplyConfiguration) WithReclaimWithinCohort(value v1beta1.PreemptionPolicy) *ClusterQueuePreemptionApplyConfiguration {
	b.ReclaimWithinCohort = &value
	return b
}

// WithBorrowWithinCohort sets the BorrowWithinCohort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BorrowWithinCohort field is set to the value of the last call.
func (b *ClusterQueuePreemptionApplyConfiguration) WithBorrowWithinCohort(value *BorrowWithinCohortApplyConfiguration) *ClusterQueuePreemptionApplyConfiguration {
	b.BorrowWithinCohort = value
	return b
}

// WithWithinClusterQueue sets the WithinClusterQueue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WithinClusterQueue field is set to the value of the last call.
func (b *ClusterQueuePreemptionApplyConfiguration) WithWithinClusterQueue(value v1beta1.PreemptionPolicy) *ClusterQueuePreemptionApplyConfiguration {
	b.WithinClusterQueue = &value
	return b
}
