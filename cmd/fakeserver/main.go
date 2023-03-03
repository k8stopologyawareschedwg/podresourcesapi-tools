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

package main

import (
	"flag"
	"os"

	"k8s.io/klog/v2"

	"github.com/k8stopologyawareschedwg/podresourcesapi-tools/pkg/podres/server"
)

func main() {
	klog.InitFlags(flag.CommandLine)

	var podresDir string

	flag.StringVar(&podresDir, "podresources-directory", "/var/run", "base directory on which create the podresources socket")
	flag.Parse()

	if err := server.Run(server.Config{
		PodresourcesDirectory: podresDir,
	}); err != nil {
		klog.ErrorS(err, "server run failed", "directory", podresDir)
		os.Exit(1)
	}
}
