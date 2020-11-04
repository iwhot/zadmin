package common

import "os"

//判断文件或者文件夹是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//清空文件内容
func ClearFile(name string) error {
	return os.Truncate(name, 0)
}

//追加写入文件
func AppendToFile(fname, content string) error {
	// 以追加模式打开文件
	txt, err := os.OpenFile(fname, os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer txt.Close()
	str := []byte(content)
	// 写入文件
	n, err := txt.Write(str)
	// 当 n != len(b) 时，返回非零错误
	if err == nil && n != len(str) {
		return err
	}
	return nil
}

//调用os.MkdirAll递归创建文件夹
func CreateFile(filePath string) error {
	if !IsExists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}
