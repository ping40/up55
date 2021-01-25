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
	LRB           = "lrb"
	XJLLB         = "xjllb"
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
	dataDir := makeDataDirectory()

	url := makeZCFZBURL(code)
	fileName := makeZCFZBFileName(code)
	if err := util.DownloadAndSave(url, dataDir, fileName); err != nil {
		fmt.Println("fetch ", fileName, "failed ")
		panic(err)
	}

	url = makeLRBURL(code)
	fileName = makeLRBFileName(code)
	if err := util.DownloadAndSave(url, dataDir, fileName); err != nil {
		fmt.Println("fetch ", fileName, "failed ")
		panic(err)
	}

	url = makeXJLLBURL(code)
	fileName = makeXJLLBFileName(code)
	if err := util.DownloadAndSave(url, dataDir, fileName); err != nil {
		fmt.Println("fetch ", fileName, "failed ")
		panic(err)
	}
}

func fileExists(code string) bool {
	fileName := filepath.Join(currentDir, DataDirectory, makeZCFZBFileName(code))
	return util.FileExists(fileName)
}

func makeDataDirectory() string {
	return filepath.Join(currentDir, DataDirectory)
}

func makeZCFZBFileName(code string) string {
	return ZCFZB + code + ".csv"
}
func makeZCFZBURL(code string) string {
	return "http://quotes.money.163.com/service/" + ZCFZB + "_" + code + ".html?type=year"
}

func makeLRBFileName(code string) string {
	return LRB + code + ".csv"
}
func makeLRBURL(code string) string {
	return "http://quotes.money.163.com/service/" + LRB + "_" + code + ".html?type=year"
}

func makeXJLLBFileName(code string) string {
	return XJLLB + code + ".csv"
}
func makeXJLLBURL(code string) string {
	return "http://quotes.money.163.com/service/" + XJLLB + "_" + code + ".html?type=year"
}
