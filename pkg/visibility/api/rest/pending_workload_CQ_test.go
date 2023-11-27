// Copyright 2023 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	kueue "sigs.k8s.io/kueue/apis/kueue/v1beta1"
	visibility "sigs.k8s.io/kueue/apis/visibility/v1alpha1"
	"sigs.k8s.io/kueue/pkg/constants"
	"sigs.k8s.io/kueue/pkg/queue"
	utiltesting "sigs.k8s.io/kueue/pkg/util/testing"
)

func TestPendingWorkloads(t *testing.T) {
	const (
		nsName   = "foo"
		cqNameA  = "cqA"
		cqNameB  = "cqB"
		lqNameA  = "lqA"
		lqNameB  = "lqB"
		lowPrio  = 50
		highPrio = 100
	)

	var (
		defaultQueryParams = &visibility.PendingWorkloadOptions{
			Offset: 0,
			Limit:  constants.DefaultPendingWorkloadsLimit,
		}
	)

	scheme := runtime.NewScheme()
	if err := kueue.AddToScheme(scheme); err != nil {
		t.Fatalf("Failed adding kueue scheme: %s", err)
	}
	if err := visibility.AddToScheme(scheme); err != nil {
		t.Fatalf("Failed adding kueue scheme: %s", err)
	}

	now := time.Now()
	cases := map[string]struct {
		clusterQueues        []*kueue.ClusterQueue
		queues               []*kueue.LocalQueue
		workloads            []*kueue.Workload
		queryParams          *visibility.PendingWorkloadOptions
		wantPendingWorkloads []visibility.PendingWorkload
	}{
		"single ClusterQueue and single LocalQueue setup with two workloads and default query parameters": {
			clusterQueues: []*kueue.ClusterQueue{
				utiltesting.MakeClusterQueue(cqNameA).Obj(),
			},
			queues: []*kueue.LocalQueue{
				utiltesting.MakeLocalQueue(lqNameA, nsName).ClusterQueue(cqNameA).Obj(),
			},
			workloads: []*kueue.Workload{
				utiltesting.MakeWorkload("a", nsName).Queue(lqNameA).Priority(highPrio).Obj(),
				utiltesting.MakeWorkload("b", nsName).Queue(lqNameA).Priority(lowPrio).Obj(),
			},
			queryParams: defaultQueryParams,
			wantPendingWorkloads: []visibility.PendingWorkload{
				{
					ObjectMeta: v1.ObjectMeta{
						Name:      "a",
						Namespace: nsName,
					},
					LocalQueueName:         lqNameA,
					Priority:               highPrio,
					PositionInClusterQueue: 0,
					PositionInLocalQueue:   0,
				},
				{
					ObjectMeta: v1.ObjectMeta{
						Name:      "b",
						Namespace: nsName,
					},
					LocalQueueName:         lqNameA,
					Priority:               lowPrio,
					PositionInClusterQueue: 1,
					PositionInLocalQueue:   1,
				},
			},
		},
		"single ClusterQueue and two LocalQueue setup with four workloads and default query parameters": {
			clusterQueues: []*kueue.ClusterQueue{
				utiltesting.MakeClusterQueue(cqNameA).Obj(),
			},
			queues: []*kueue.LocalQueue{
				utiltesting.MakeLocalQueue(lqNameA, nsName).ClusterQueue(cqNameA).Obj(),
				utiltesting.MakeLocalQueue(lqNameB, nsName).ClusterQueue(cqNameA).Obj(),
			},
			workloads: []*kueue.Workload{
				utiltesting.MakeWorkload("lqA-high-prio", nsName).Queue(lqNameA).Priority(highPrio).Creation(now).Obj(),
				utiltesting.MakeWorkload("lqA-low-prio", nsName).Queue(lqNameA).Priority(lowPrio).Creation(now).Obj(),
				utiltesting.MakeWorkload("lqB-high-prio", nsName).Queue(lqNameB).Priority(highPrio).Creation(now.Add(time.Second)).Obj(),
				utiltesting.MakeWorkload("lqB-low-prio", nsName).Queue(lqNameB).Priority(lowPrio).Creation(now.Add(time.Second)).Obj(),
			},
			queryParams: defaultQueryParams,
			wantPendingWorkloads: []visibility.PendingWorkload{
				{
					ObjectMeta: v1.ObjectMeta{
						Name:      "lqA-high-prio",
						Namespace: nsName,
					},
					LocalQueueName:         lqNameA,
					Priority:               highPrio,
					PositionInClusterQueue: 0,
					PositionInLocalQueue:   0,
				},
				{
					ObjectMeta: v1.ObjectMeta{
						Name:      "lqB-high-prio",
						Namespace: nsName,
					},
					LocalQueueName:         lqNameB,
					Priority:               highPrio,
					PositionInClusterQueue: 1,
					PositionInLocalQueue:   0,
				},
				{
					ObjectMeta: v1.ObjectMeta{
						Name:      "lqA-low-prio",
						Namespace: nsName,
					},
					LocalQueueName:         lqNameA,
					Priority:               lowPrio,
					PositionInClusterQueue: 2,
					PositionInLocalQueue:   1,
				},
				{
					ObjectMeta: v1.ObjectMeta{
						Name:      "lqB-low-prio",
						Namespace: nsName,
					},
					LocalQueueName:         lqNameB,
					Priority:               lowPrio,
					PositionInClusterQueue: 3,
					PositionInLocalQueue:   1,
				},
			},
		},
		"query parameters limit set": {
			clusterQueues: []*kueue.ClusterQueue{
				utiltesting.MakeClusterQueue(cqNameA).Obj(),
			},
			queues: []*kueue.LocalQueue{
				utiltesting.MakeLocalQueue(lqNameA, nsName).ClusterQueue(cqNameA).Obj(),
			},
			workloads: []*kueue.Workload{
				utiltesting.MakeWorkload("a", nsName).Queue(lqNameA).Priority(highPrio).Creation(now).Obj(),
				utiltesting.MakeWorkload("b", nsName).Queue(lqNameA).Priority(highPrio).Creation(now.Add(time.Second)).Obj(),
				utiltesting.MakeWorkload("c", nsName).Queue(lqNameA).Priority(highPrio).Creation(now.Add(time.Second * 2)).Obj(),
			},
			queryParams: &visibility.PendingWorkloadOptions{
				Limit: 2,
			},
			wantPendingWorkloads: []visibility.PendingWorkload{
				{
					ObjectMeta: v1.ObjectMeta{
						Name:      "a",
						Namespace: nsName,
					},
					LocalQueueName:         lqNameA,
					Priority:               highPrio,
					PositionInClusterQueue: 0,
					PositionInLocalQueue:   0,
				},
				{
					ObjectMeta: v1.ObjectMeta{
						Name:      "b",
						Namespace: nsName,
					},
					LocalQueueName:         lqNameA,
					Priority:               highPrio,
					PositionInClusterQueue: 1,
					PositionInLocalQueue:   1,
				},
			},
		},
		"query parameters offset set": {
			clusterQueues: []*kueue.ClusterQueue{
				utiltesting.MakeClusterQueue(cqNameA).Obj(),
			},
			queues: []*kueue.LocalQueue{
				utiltesting.MakeLocalQueue(lqNameA, nsName).ClusterQueue(cqNameA).Obj(),
			},
			workloads: []*kueue.Workload{
				utiltesting.MakeWorkload("a", nsName).Queue(lqNameA).Priority(highPrio).Creation(now).Obj(),
				utiltesting.MakeWorkload("b", nsName).Queue(lqNameA).Priority(highPrio).Creation(now.Add(time.Second)).Obj(),
				utiltesting.MakeWorkload("c", nsName).Queue(lqNameA).Priority(highPrio).Creation(now.Add(time.Second * 2)).Obj(),
			},
			queryParams: &visibility.PendingWorkloadOptions{
				Offset: 1,
				Limit:  constants.DefaultPendingWorkloadsLimit,
			},
			wantPendingWorkloads: []visibility.PendingWorkload{
				{
					ObjectMeta: v1.ObjectMeta{
						Name:      "b",
						Namespace: nsName,
					},
					LocalQueueName:         lqNameA,
					Priority:               highPrio,
					PositionInClusterQueue: 1,
					PositionInLocalQueue:   1,
				},
				{
					ObjectMeta: v1.ObjectMeta{
						Name:      "c",
						Namespace: nsName,
					},
					LocalQueueName:         lqNameA,
					Priority:               highPrio,
					PositionInClusterQueue: 2,
					PositionInLocalQueue:   2,
				},
			},
		},
		"query parameters offset and limit set": {
			clusterQueues: []*kueue.ClusterQueue{
				utiltesting.MakeClusterQueue(cqNameA).Obj(),
			},
			queues: []*kueue.LocalQueue{
				utiltesting.MakeLocalQueue(lqNameA, nsName).ClusterQueue(cqNameA).Obj(),
			},
			workloads: []*kueue.Workload{
				utiltesting.MakeWorkload("a", nsName).Queue(lqNameA).Priority(highPrio).Creation(now).Obj(),
				utiltesting.MakeWorkload("b", nsName).Queue(lqNameA).Priority(highPrio).Creation(now.Add(time.Second)).Obj(),
				utiltesting.MakeWorkload("c", nsName).Queue(lqNameA).Priority(highPrio).Creation(now.Add(time.Second * 2)).Obj(),
			},
			queryParams: &visibility.PendingWorkloadOptions{
				Offset: 1,
				Limit:  1,
			},
			wantPendingWorkloads: []visibility.PendingWorkload{
				{
					ObjectMeta: v1.ObjectMeta{
						Name:      "b",
						Namespace: nsName,
					},
					LocalQueueName:         lqNameA,
					Priority:               highPrio,
					PositionInClusterQueue: 1,
					PositionInLocalQueue:   1,
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			manager := queue.NewManager(utiltesting.NewFakeClient(), nil)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			go manager.CleanUpOnContext(ctx)

			pendingWorkloadsInCqRest := NewPendingWorkloadsInCqREST(manager)
			for _, cq := range tc.clusterQueues {
				if err := manager.AddClusterQueue(ctx, cq); err != nil {
					t.Fatalf("Adding cluster queue %s: %v", cq.Name, err)
				}
			}
			for _, q := range tc.queues {
				if err := manager.AddLocalQueue(ctx, q); err != nil {
					t.Fatalf("Adding queue %q: %v", q.Name, err)
				}
			}
			for _, w := range tc.workloads {
				manager.AddOrUpdateWorkload(w)
			}

			for _, cq := range tc.clusterQueues {
				info, err := pendingWorkloadsInCqRest.Get(ctx, cq.Name, tc.queryParams)
				if err != nil {
					t.Fatal(err)
				}
				pendingWorkloadsInfo := info.(*visibility.PendingWorkloadsSummary)

				if diff := cmp.Diff(tc.wantPendingWorkloads, pendingWorkloadsInfo.Items, cmpopts.IgnoreTypes(metav1.TypeMeta{})); diff != "" {
					t.Errorf("Pending workloads differ: (-want,+got):\n%s", diff)
				}
			}
		})
	}
}
