package util

import (
	"os"
)

func Write(fileName, content string) error {

	dstFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	dstFile.WriteString(content)

	return nil
}

func FileExists(fullFileName string) bool {
	_, err := os.Lstat(fullFileName)
	return !os.IsNotExist(err)
}
