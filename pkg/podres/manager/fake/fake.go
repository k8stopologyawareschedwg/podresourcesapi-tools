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
 * Copyright 2023 The Kubernetes Authors.
 */

package fake

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	podresourcesapi "k8s.io/kubelet/pkg/apis/podresources/v1"
)

type Devices struct{}

func (dd Devices) UpdateAllocatedDevices() {
	klog.V(3).InfoS("Invoked fake UpdateAllocatedDevices()")
}

func (dd Devices) GetDevices(podUID, containerName string) []*podresourcesapi.ContainerDevices {
	klog.V(3).InfoS("Invoked fake GetDevices", "podUID", podUID, "containerName", containerName)
	return nil
}

func (dd Devices) GetAllocatableDevices() []*podresourcesapi.ContainerDevices {
	klog.V(3).InfoS("Invoked fake GetAllocatableDevices()")
	return nil
}

type Pods struct{}

func (pp Pods) GetPods() []*corev1.Pod {
	klog.V(3).InfoS("Invoked fake GetPods()")
	return nil
}

type CPUs struct{}

func (cc CPUs) GetCPUs(podUID, containerName string) []int64 {
	klog.V(3).InfoS("Invoked fake GetCPUs()", "podUID", podUID, "containerName", containerName)
	return nil
}

func (cc CPUs) GetAllocatableCPUs() []int64 {
	klog.V(3).InfoS("Invoked fake GetAllocatableCPUs()")
	return nil
}

type Memory struct{}

func (mm Memory) GetMemory(podUID, containerName string) []*podresourcesapi.ContainerMemory {
	klog.V(3).InfoS("Invoked fake GetMemory()", "podUID", podUID, "containerName", containerName)
	return nil
}

func (mm Memory) GetAllocatableMemory() []*podresourcesapi.ContainerMemory {
	klog.V(3).InfoS("Invoked fake GetAllocatableMemory()")
	return nil
}
