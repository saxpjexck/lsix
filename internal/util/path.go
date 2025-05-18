package util

import (
	"os"
	"path/filepath"
)

// 获取执行文件的目录
func GetBinDir() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exePath)
}
