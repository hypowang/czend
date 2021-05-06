package util

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// check file or directory exists
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//check file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	if info == nil || info.IsDir(){
		return false
	}
	return true
}

// CreatNestedFile 给定path创建文件，如果目录不存在就递归创建
func CreatNestedFile(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if !Exists(basePath) {
		err := os.MkdirAll(basePath, 0700)
		if err != nil {
			Log().Warning("无法创建目录，%s", err)
			return nil, err
		}
	}

	return os.Create(path)
}

// 移除路径最后的`/`
func RemoveSlash(path string) string {
	if len(path) > 1 {
		return strings.TrimSuffix(path, "/")
	}
	return path
}

// RelativePath 获取相对可执行文件的路径
func RelativePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	e, _ := os.Executable()
	return filepath.Join(filepath.Dir(e), path)
}

// FixPath 将path中的反斜杠'\'替换为'/'
func FixPath(old string) string {
	return path.Clean(strings.ReplaceAll(old, "\\", "/"))
}
