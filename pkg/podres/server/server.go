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

package server

import (
	"google.golang.org/grpc"

	"k8s.io/klog/v2"
	podresourcesapi "k8s.io/kubelet/pkg/apis/podresources/v1"
	podresourcesapiv1alpha1 "k8s.io/kubelet/pkg/apis/podresources/v1alpha1"
	"k8s.io/kubernetes/pkg/kubelet/apis/podresources"
	"k8s.io/kubernetes/pkg/kubelet/util"

	fakeprovider "github.com/k8stopologyawareschedwg/podresourcesapi-tools/pkg/podres/manager/fake"
)

type Config struct {
	PodresourcesDirectory string `json:"podresourcesDirectory"`
}

func Run(conf Config) error {
	klog.InfoS("Running podresources server", "configuration", conf)

	socket, err := util.LocalEndpoint(conf.PodresourcesDirectory, podresources.Socket)
	if err != nil {
		klog.V(2).InfoS("Failed to get local endpoint for PodResources endpoint", "err", err)
		return err
	}

	klog.InfoS("Local endpoint created", "path", socket)

	podsProvider := fakeprovider.Pods{}
	devicesProvider := fakeprovider.Devices{}
	cpusProvider := fakeprovider.CPUs{}
	memoryProvider := fakeprovider.Memory{}

	server := grpc.NewServer()
	podresourcesapiv1alpha1.RegisterPodResourcesListerServer(server, podresources.NewV1alpha1PodResourcesServer(podsProvider, devicesProvider))
	podresourcesapi.RegisterPodResourcesListerServer(server, podresources.NewV1PodResourcesServer(podsProvider, devicesProvider, cpusProvider, memoryProvider))

	lst, err := util.CreateListener(socket)
	if err != nil {
		klog.ErrorS(err, "Failed to create listener for podResources endpoint")
		return err
	}

	klog.InfoS("Serving...")

	return server.Serve(lst)
}
