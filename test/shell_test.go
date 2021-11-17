package test

import (
	"os/exec"
	"testing"
)

//shell 测试
func TestShell(t *testing.T) {
	command := `./bin/test.sh .`
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		t.Logf("Execute shell:%s failed with error:%s", command, err.Error())
		return
	}
	t.Logf("Execute shell:%s finished with output:\n%s", command, string(output))
}
