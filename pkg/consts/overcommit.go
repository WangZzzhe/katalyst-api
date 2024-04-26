/*
Copyright 2022 The Katalyst Authors.

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

package consts

// const variables for node annotations about overcommit ratio
const (
	NodeAnnotationCPUOvercommitRatioKey            = "halo.io/cpu_overcommit_ratio"
	NodeAnnotationMemoryOvercommitRatioKey         = "halo.io/memory_overcommit_ratio"
	NodeAnnotationRealtimeCPUOvercommitRatioKey    = "halo.io/realtime_cpu_overcommit_ratio"
	NodeAnnotationRealtimeMemoryOvercommitRatioKey = "halo.io/realtime_memory_overcommit_ratio"

	NodeAnnotationPredictCPUOvercommitRatioKey    = "halo.io/predict_cpu_overcommit_ratio"
	NodeAnnotationPredictMemoryOvercommitRatioKey = "halo.io/predict_memory_overcommit_ratio"

	NodeAnnotationOriginalCapacityCPUKey       = "halo.io/original_capacity_cpu"
	NodeAnnotationOriginalCapacityMemoryKey    = "halo.io/original_capacity_memory"
	NodeAnnotationOriginalAllocatableCPUKey    = "halo.io/original_allocatable_cpu"
	NodeAnnotationOriginalAllocatableMemoryKey = "halo.io/original_allocatable_memory"

	NodeAnnotationOvercommitCapacityCPUKey       = "halo.io/overcommit_capacity_cpu"
	NodeAnnotationOvercommitAllocatableCPUKey    = "halo.io/overcommit_allocatable_cpu"
	NodeAnnotationOvercommitCapacityMemoryKey    = "halo.io/overcommit_capacity_memory"
	NodeAnnotationOvercommitAllocatableMemoryKey = "halo.io/overcommit_allocatable_memory"
)

// const variables for matching up with node labels about overcommit
const (
	NodeOvercommitSelectorKey = "halo.io/overcommit_node_pool"

	DefaultNodeCPUOvercommitRatio    = "1"
	DefaultNodeMemoryOvercommitRatio = "1"
)

type KCNRAnnotationCPUManagerPolicy string
type KCNRAnnotationMemoryManagerPolicy string

const (
	// KCNRAnnotationGuaranteedCPUs sum of pod guaranteed cpus in node
	KCNRAnnotationGuaranteedCPUs = "halo.io/guaranteed_cpus"

	KCNRAnnotationCPUManager    = "halo.io/overcommit_cpu_manager"
	KCNRAnnotationMemoryManager = "halo.io/overcommit_memory_manager"

	CPUManagerOff        KCNRAnnotationCPUManagerPolicy = "none"
	CPUManagerPolicyNone KCNRAnnotationCPUManagerPolicy = "none"

	MemoryManagerOff        KCNRAnnotationMemoryManagerPolicy = "None"
	MemoryManagerPolicyNone KCNRAnnotationMemoryManagerPolicy = "None"
)
