/*
Copyright 2024 The Kubernetes Authors.

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

package wrappers

import (
	rayv1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/kueue/cmd/experimental/kjobctl/pkg/constants"
	kueueconstants "sigs.k8s.io/kueue/pkg/controller/constants"
)

// RayJobWrapper wraps a RayJob.
type RayJobWrapper struct{ rayv1.RayJob }

// MakeRayJob creates a wrapper for a RayJob
func MakeRayJob(name, ns string) *RayJobWrapper {
	return &RayJobWrapper{
		rayv1.RayJob{
			TypeMeta: metav1.TypeMeta{Kind: "RayJob", APIVersion: "ray.io/v1"},
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: ns,
			},
		},
	}
}

// Obj returns the inner RayJob.
func (j *RayJobWrapper) Obj() *rayv1.RayJob {
	return &j.RayJob
}

// GenerateName updates generateName.
func (j *RayJobWrapper) GenerateName(v string) *RayJobWrapper {
	j.ObjectMeta.GenerateName = v
	return j
}

// Profile sets the profile label.
func (j *RayJobWrapper) Profile(v string) *RayJobWrapper {
	return j.Label(constants.ProfileLabel, v)
}

// LocalQueue sets the localqueue label.
func (j *RayJobWrapper) LocalQueue(v string) *RayJobWrapper {
	return j.Label(kueueconstants.QueueLabel, v)
}

// Label sets the label key and value.
func (j *RayJobWrapper) Label(key, value string) *RayJobWrapper {
	if j.Labels == nil {
		j.Labels = make(map[string]string)
	}
	j.ObjectMeta.Labels[key] = value
	return j
}

// WithWorkerGroupSpec add worker group to the ray cluster template.
func (j *RayJobWrapper) WithWorkerGroupSpec(spec rayv1.WorkerGroupSpec) *RayJobWrapper {
	if j.RayJob.Spec.RayClusterSpec == nil {
		j.RayJob.Spec.RayClusterSpec = &rayv1.RayClusterSpec{}
	}

	j.RayJob.Spec.RayClusterSpec.WorkerGroupSpecs = append(j.RayJob.Spec.RayClusterSpec.WorkerGroupSpecs, spec)

	return j
}

// Spec set job spec.
func (j *RayJobWrapper) Spec(spec rayv1.RayJobSpec) *RayJobWrapper {
	j.RayJob.Spec = spec
	return j
}

// Entrypoint set entrypoint.
func (j *RayJobWrapper) Entrypoint(entrypoint string) *RayJobWrapper {
	j.RayJob.Spec.Entrypoint = entrypoint
	return j
}
