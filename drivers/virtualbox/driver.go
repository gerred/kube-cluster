// Copyright 2015 The kube-cluster Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package virtualbox

import (
	"fmt"
	"os/exec"
	"runtime"
)

const (
	VBoxManageBin   = "VBoxManage"
	VBoxHeadlessBin = "VBoxHeadless"
)

type Virtualbox struct {
	mgmtbin     string
	headlessbin string
	envName     string
}

// todo(carlos): test for minimum VBox version

func New(envName string) (*Virtualbox, error) {
	dotExe := ""
	if runtime.GOOS == "windows" {
		dotExe = ".exe"
	}
	mgmtbin, err := exec.LookPath(VBoxManageBin+dotExe)
	if err != nil {
		return nil, fmt.Errorf("could not find %s: %v", VBoxManageBin, err)
	}
	headlessbin, err := exec.LookPath(VBoxHeadlessBin+dotExe)
	if err != nil {
		return nil, fmt.Errorf("could not find %s: %v", VBoxHeadlessBin, err)
	}

	v := &Virtualbox{
		mgmtbin:     mgmtbin,
		headlessbin: headlessbin,
		envName:     envName,
	}

	return v, nil
}

func (v *Virtualbox) GenerateCerts() {
}

func (v *Virtualbox) GetTokens() {
}

func (v *Virtualbox) ProvisionMaster() {
}

func (v *Virtualbox) ConfigureMaster() {
}

func (v *Virtualbox) ProvisionNode() {
}

func (v *Virtualbox) ConfigureNode() {
}
