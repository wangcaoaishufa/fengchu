package util

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// RootPath 获取此项目的绝对路径
// 如果是以 go build 生成的二进制文件运行，则返回 bin 目录的上级目录的绝对路径
// 如果是以 go run 运行，则返回在此项目的绝对路径
func RootPath() string {
	var binDir string

	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	binDir = filepath.Dir(filepath.Dir(exePath))

	tmpDir := os.TempDir()
	if strings.Contains(exePath, tmpDir) {
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			binDir = filepath.Dir(filepath.Dir(filepath.Dir(filename)))
		}
	}

	return binDir
}

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			log.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				log.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
