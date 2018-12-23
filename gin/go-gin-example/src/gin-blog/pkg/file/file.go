package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// 获取后缀
func GetExt(filename string) string {
	return path.Ext(filename)
}

// 检查是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// 检查权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// 如果不存在创建路径
func IsNotExistMkDir(src string) error {
	if exist := CheckExist(src); exist == true {
		if err := Mkdir(src); err != nil {
			return err
		}
	}
	return nil
}

// 创建路径
func Mkdir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
