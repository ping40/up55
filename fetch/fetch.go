package fetch

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ping40/up55/util"
)

//http://quotes.money.163.com/service/zcfzb_600519.html?type=year

const (
	DataDirectory = "data"
	ZCFZB         = "zcfzb"
	ZCFZB1        = "zcfzb"
)

var currentDir string

func init() {
	currentDir, _ = os.Getwd()
}

func Fetch(code string) {
	if fileExists(code) {
		fmt.Println("文件存在，", code)
		return
	}

	fmt.Println("文件不存在，", code)

}

func fileExists(code string) bool {
	fileName := filepath.Join(currentDir, DataDirectory, makeZCFZBFileName(code))
	return util.FileExists(fileName)
}

func makeZCFZBFileName(code string) string {
	return ZCFZB + code + ".csv"
}
