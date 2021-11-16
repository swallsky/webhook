package app

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/gin-gonic/gin"
)

// shell文件测试
func ShellTest(c *gin.Context) {
	command := `./bin/test.sh .`
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute shell:%s finished with output:\n%s", command, string(output))
}

// 执行各分支脚本
func ShellBin(token string, ref string) string {
	fileName := path.Join("./bin", token+".sh")
	//判断文件是否存在
	if _, err := os.Stat(fileName); err == nil {
		command := fileName + ` ` + ref
		cmd := exec.Command("/bin/bash", "-c", command)
		output, err := cmd.Output()
		if err == nil {
			return fmt.Sprintf("Execute shell:%s finished with output:\n%s", command, string(output))
		} else {
			return fmt.Sprintf("Execute shell:%s failed with error:%s", command, err.Error())
		}
	} else {
		return fmt.Sprintf("Shell file:%s is not found.\n", fileName)
	}
}
