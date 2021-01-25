package util

import "os"

func CreateDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(path, 0777) //0777也可以os.ModePerm
		os.Chmod(path, 0777)
	}else if err != nil{
		return err
	}
	return nil
}