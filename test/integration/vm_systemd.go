// +build integration

/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package integration

import (
	"strings"
	"testing"

	"k8s.io/minikube/test/integration/util"
)

func testVMSystemd(t *testing.T) {
	t.Parallel()
	minikubeRunner := util.MinikubeRunner{
		Args:       *args,
		BinaryPath: *binaryPath,
		T:          t}

	expectedStr := "0 loaded units listed"
	sshCmdOutput := minikubeRunner.RunCommand("ssh -- systemctl --state=failed --all --no-pager", true)
	if !strings.Contains(sshCmdOutput, expectedStr) {
		t.Logf("Expected no systemd units to be failed got:\n %s", sshCmdOutput)
	}
}
