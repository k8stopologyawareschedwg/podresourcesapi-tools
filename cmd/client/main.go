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
 * Copyright 2022 The Kubernetes Authors.
 */

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"k8s.io/klog/v2"
	kubeletpodresourcesv1 "k8s.io/kubelet/pkg/apis/podresources/v1"
	"k8s.io/kubernetes/pkg/kubelet/apis/podresources"

	"github.com/k8stopologyawareschedwg/podresourcesapi-tools/pkg/podres"
	podresdefault "github.com/k8stopologyawareschedwg/podresourcesapi-tools/pkg/podres/defaults"
)

const (
	apiCallList           = "list"
	apiCallGetAllocatable = "get-allocatable"
)

// we fill our own structs to avoid the problem when default int value(0) removed from the json
func selectAction(apiName string) (func(cli kubeletpodresourcesv1.PodResourcesListerClient) error, error) {
	if apiName == apiCallList {
		return func(cli kubeletpodresourcesv1.PodResourcesListerClient) error {
			resp, err := cli.List(context.TODO(), &kubeletpodresourcesv1.ListPodResourcesRequest{})
			if err != nil {
				return err
			}

			listPodResourcesResp := podres.ConvertListPodResourcesResponseFromK(resp)
			if err := json.NewEncoder(os.Stdout).Encode(listPodResourcesResp); err != nil {
				return err
			}

			return nil
		}, nil
	}
	if apiName == apiCallGetAllocatable {
		return func(cli kubeletpodresourcesv1.PodResourcesListerClient) error {
			resp, err := cli.GetAllocatableResources(context.TODO(), &kubeletpodresourcesv1.AllocatableResourcesRequest{})
			if err != nil {
				return err
			}

			allocatableResourcesResponse := podres.ConvertAllocatableResourcesResponseFromK(resp)
			if err := json.NewEncoder(os.Stdout).Encode(allocatableResourcesResponse); err != nil {
				return err
			}

			return nil
		}, nil
	}
	return func(cli kubeletpodresourcesv1.PodResourcesListerClient) error {
		return nil
	}, fmt.Errorf("unknown API %q", apiName)
}

func main() {
	klog.InitFlags(flag.CommandLine)

	var podresSock string

	flag.StringVar(&podresSock, "podresources-socket", podresdefault.SocketPath, "podresources server socket to connect to")
	flag.Parse()

	args := flag.Args()

	apiName := "list"
	if len(args) == 1 {
		apiName = args[0]
	}

	action, err := selectAction(apiName)
	if err != nil {
		klog.ErrorS(err, "unknown API", "api", apiName)
		os.Exit(1)
	}

	cli, conn, err := podresources.GetV1Client(podresSock, podresdefault.Timeout, podresdefault.MaxSize)
	if err != nil {
		klog.ErrorS(err, "failed to connect", "api", apiName, "socketPath", podresSock)
		os.Exit(2)
	}
	defer conn.Close()

	if err := action(cli); err != nil {
		klog.ErrorS(err, "call failed", "api", apiName)
		os.Exit(4)
	}
}
