package utils

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// If 模拟简单的三元操作
func If(condition bool, trueVal, falseVal any) any {
	if condition {
		return trueVal
	}
	return falseVal
}

// GetRunPath 获取执行目录作为默认目录
func GetRunPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		return ""
	}
	return currentPath
}

// GetFileDirectoryToCaller 根据运行堆栈信息获取文件目录，skip 默认1
func GetFileDirectoryToCaller(opts ...int) (directory string, ok bool) {
	var filename string
	directory = ""
	skip := 1
	if opts != nil {
		skip = opts[0]
	}
	if _, filename, _, ok = runtime.Caller(skip); ok {
		fmt.Println("filename", filename)
		directory = path.Dir(filename)
	}
	return
}

// GetCurrentAbPathByExecutable 获取当前执行文件绝对路径
func GetCurrentAbPathByExecutable() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	res, _ := filepath.EvalSymlinks(exePath)
	return filepath.Dir(res), nil
}

// GetCurrentPath 获取当前执行文件路径
func GetCurrentPath() (dir string, err error) {
	dir, err = GetCurrentAbPathByExecutable()
	fmt.Println("GetCurrentAbPathByExecutable", dir)
	if err != nil {
		return "", err
	}

	tmpDir, err := filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		return "", err
	}
	fmt.Println("tmpDir", tmpDir)
	v0, _ := GetFileDirectoryToCaller(0)
	fmt.Println("v0", v0)
	v1, _ := GetFileDirectoryToCaller(1)
	fmt.Println("v1", v1)
	v2, _ := GetFileDirectoryToCaller(2)
	fmt.Println("v2", v2)
	//go run 模式下的情况
	if strings.Contains(dir, tmpDir) {
		var ok bool
		if dir, ok = GetFileDirectoryToCaller(2); !ok {
			return "", errors.New("failed to get path")
		}
	}
	return dir, nil
}
