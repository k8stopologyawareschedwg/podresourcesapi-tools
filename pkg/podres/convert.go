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

import (
	kubeletpodresourcesv1 "k8s.io/kubelet/pkg/apis/podresources/v1"
)

func ConvertListPodResourcesResponseFromK(resp *kubeletpodresourcesv1.ListPodResourcesResponse) *ListPodResourcesResponse {
	var podResources []*PodResources
	for _, podRes := range resp.PodResources {
		var podResContainers []*ContainerResources
		for _, c := range podRes.Containers {
			podResContainers = append(podResContainers, &ContainerResources{
				Name:    c.Name,
				CpuIds:  c.CpuIds,
				Devices: ConvertDevicesFromK(c.Devices),
				Memory:  ConvertMemoryFromK(c.Memory),
			})
		}

		podResources = append(podResources, &PodResources{
			Name:       podRes.Name,
			Namespace:  podRes.Namespace,
			Containers: podResContainers,
		})
	}

	return &ListPodResourcesResponse{
		PodResources: podResources,
	}
}

func ConvertAllocatableResourcesResponseFromK(resp *kubeletpodresourcesv1.AllocatableResourcesResponse) *AllocatableResourcesResponse {
	return &AllocatableResourcesResponse{
		CpuIds:  resp.CpuIds,
		Devices: ConvertDevicesFromK(resp.Devices),
		Memory:  ConvertMemoryFromK(resp.Memory),
	}
}

func ConvertDevicesFromK(containerDevices []*kubeletpodresourcesv1.ContainerDevices) []*ContainerDevices {
	var cDevices []*ContainerDevices
	for _, d := range containerDevices {
		deviceTopologyInfo := ConvertTopologyInfoFromK(d.Topology)
		cDevices = append(cDevices, &ContainerDevices{
			ResourceName: d.ResourceName,
			DeviceIds:    d.DeviceIds,
			Topology:     deviceTopologyInfo,
		})
	}

	return cDevices
}

func ConvertMemoryFromK(containerMemory []*kubeletpodresourcesv1.ContainerMemory) []*ContainerMemory {
	var cMemory []*ContainerMemory
	for _, m := range containerMemory {
		memoryTopologyInfo := ConvertTopologyInfoFromK(m.Topology)
		cMemory = append(cMemory, &ContainerMemory{
			MemoryType: m.MemoryType,
			Size_:      m.Size_,
			Topology:   memoryTopologyInfo,
		})
	}

	return cMemory
}

func ConvertTopologyInfoFromK(topologyInfo *kubeletpodresourcesv1.TopologyInfo) *TopologyInfo {
	if topologyInfo == nil {
		return nil
	}

	var numaNodes []*NUMANode
	for _, numaNode := range topologyInfo.Nodes {
		numaNodeID := numaNode.ID
		numaNodes = append(numaNodes, &NUMANode{
			ID: &numaNodeID,
		})
	}

	return &TopologyInfo{
		Nodes: numaNodes,
	}
}
