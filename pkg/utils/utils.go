package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// UtilsDir 此文件相对于根路径,所在的文件夹
const UtilsDir = "pkg/utils"

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

	if err != nil {
		return "", err
	}

	tmpDir, err := filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		return "", err
	}

	//go run 模式下的情况
	if strings.Contains(dir, tmpDir) {
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			dir = path.Dir(filename)
		}

		dir = strings.Replace(dir, UtilsDir, "", 1)
		fmt.Println("dir", dir)
	}
	return dir, nil
}
