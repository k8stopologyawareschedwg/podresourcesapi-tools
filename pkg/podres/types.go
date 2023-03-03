/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2020 The Kubernetes Authors.
 */

package podres

// The structs are copied from k/k/pkg/apis/podresources/v1/api.pb.go and should be in sync with it.
// Should not be a big problem because it is v1 API and should not have a lot of changes.

// ListPodResourcesResponse is the response returned by List function
type ListPodResourcesResponse struct {
	PodResources []*PodResources `json:"pod_resources,omitempty"`
}

// PodResources contains information about the node resources assigned to a pod
type PodResources struct {
	Name       string                `json:"name,omitempty"`
	Namespace  string                `json:"namespace,omitempty"`
	Containers []*ContainerResources `json:"containers,omitempty"`
}

// ContainerResources contains information about the resources assigned to a container
type ContainerResources struct {
	Name    string              `json:"name,omitempty"`
	Devices []*ContainerDevices `json:"devices,omitempty"`
	CpuIds  []int64             `json:"cpu_ids,omitempty"`
	Memory  []*ContainerMemory  `json:"memory,omitempty"`
}

// AllocatableResourcesResponse contains information about all the devices known by the kubelet
type AllocatableResourcesResponse struct {
	Devices []*ContainerDevices `json:"devices,omitempty"`
	CpuIds  []int64             `json:"cpu_ids,omitempty"`
	Memory  []*ContainerMemory  `json:"memory,omitempty"`
}

// ContainerDevices contains information about the devices assigned to a container
type ContainerDevices struct {
	ResourceName string        `json:"resource_name,omitempty"`
	DeviceIds    []string      `json:"device_ids,omitempty"`
	Topology     *TopologyInfo `json:"topology,omitempty"`
}

// ContainerMemory contains information about memory and hugepages assigned to a container
type ContainerMemory struct {
	MemoryType string        `json:"memory_type,omitempty"`
	Size_      uint64        `json:"size,omitempty"`
	Topology   *TopologyInfo `json:"topology,omitempty"`
}

type TopologyInfo struct {
	Nodes []*NUMANode `json:"nodes,omitempty"`
}

// NUMANode contains NUMA nodes information
type NUMANode struct {
	ID *int64 `json:"ID,omitempty"`
}
