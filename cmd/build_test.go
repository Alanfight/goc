/*
 Copyright 2020 Qiniu Cloud (qiniu.com)

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

package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var baseDir string

func init() {
	baseDir, _ = os.Getwd()
}

func TestGeneratedBinary(t *testing.T) {
	startTime := time.Now()

	workingDir := filepath.Join(baseDir, "../tests/samples/simple_project")
	gopath := ""

	os.Setenv("GOPATH", gopath)
	os.Setenv("GO111MODULE", "on")

	buildFlags, buildOutput = "", ""
	args := []string{"."}
	runBuild(args, workingDir)

	obj := filepath.Join(workingDir, "simple-project")
	fInfo, err := os.Lstat(obj)
	assert.Equal(t, err, nil, "the binary should be generated.")
	assert.Equal(t, startTime.Before(fInfo.ModTime()), true, obj+"new binary should be generated, not the old one")

	cmd := exec.Command("go", "tool", "objdump", "simple-project")
	cmd.Dir = workingDir
	out, _ := cmd.CombinedOutput()
	cnt := strings.Count(string(out), "main.registerSelf")
	assert.Equal(t, cnt > 0, true, "main.registerSelf function should be in the binary")

	cnt = strings.Count(string(out), "GoCover")
	assert.Equal(t, cnt > 0, true, "GoCover varibale should be in the binary")
}

func TestBuildBinaryName(t *testing.T) {
	startTime := time.Now()

	workingDir := filepath.Join(baseDir, "../tests/samples/simple_project2")
	gopath := ""

	os.Setenv("GOPATH", gopath)
	os.Setenv("GO111MODULE", "on")

	buildFlags, buildOutput = "", ""
	args := []string{"."}
	runBuild(args, workingDir)

	obj := filepath.Join(workingDir, "simple-project")
	fInfo, err := os.Lstat(obj)
	assert.Equal(t, err, nil, "the binary should be generated.")
	assert.Equal(t, startTime.Before(fInfo.ModTime()), true, obj+"new binary should be generated, not the old one")

	cmd := exec.Command("go", "tool", "objdump", "simple-project")
	cmd.Dir = workingDir
	out, _ := cmd.CombinedOutput()
	cnt := strings.Count(string(out), "main.registerSelf")
	assert.Equal(t, cnt > 0, true, "main.registerSelf function should be in the binary")

	cnt = strings.Count(string(out), "GoCover")
	assert.Equal(t, cnt > 0, true, "GoCover varibale should be in the binary")
}

// test if goc can get variables in internal package
func TestBuildBinaryForInternalPackage(t *testing.T) {
	startTime := time.Now()

	workingDir := filepath.Join(baseDir, "../tests/samples/simple_project_with_internal")
	gopath := ""

	os.Setenv("GOPATH", gopath)
	os.Setenv("GO111MODULE", "on")

	buildFlags, buildOutput = "", ""
	args := []string{"."}
	runBuild(args, workingDir)

	obj := filepath.Join(workingDir, "simple-project")
	fInfo, err := os.Lstat(obj)
	assert.Equal(t, err, nil, "the binary should be generated.")
	assert.Equal(t, startTime.Before(fInfo.ModTime()), true, obj+"new binary should be generated, not the old one")

	cmd := exec.Command("go", "tool", "objdump", "simple-project")
	cmd.Dir = workingDir
	out, _ := cmd.CombinedOutput()
	cnt := strings.Count(string(out), "GoCacheCover")
	assert.Equal(t, cnt > 0, true, "GoCacheCover variable for internal package should be in the binary")

	cnt = strings.Count(string(out), "internal.GoCover")
	assert.Equal(t, cnt > 0, true, "internal.GoCover varibale should be in the binary")
}
