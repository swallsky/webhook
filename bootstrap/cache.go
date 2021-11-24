package bootstrap

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var runtimePath string

// 初始化
func init() {
	runtimePath = "runtime"
	if dir, err := os.Getwd(); err == nil {
		runtimePath = dir + "/runtime/"
	}
	// 如果目录不存在，则创建
	if err := os.MkdirAll(runtimePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
}

// 创建文件夹
func CreateDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}

// 写文件
func WriteFile(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(content) //写入文件

	//Flush将缓存的文件真正写入到文件中
	err = write.Flush()
	return err
}

// 读文件
func ReadFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
