package util

import (
	"os"
	"path/filepath"
)

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

func CodeFileExists(code string) bool {
	fileName := filepath.Join(currentDir, DataDirectory, MakeZCFZBFileName(code))
	return fileExists(fileName)
}

func fileExists(fullFileName string) bool {
	_, err := os.Lstat(fullFileName)
	return !os.IsNotExist(err)
}

func MakeDataDirectory() string {
	return filepath.Join(currentDir, DataDirectory)
}

func MakeZCFZBFileName(code string) string {
	return ZCFZB + code + ".csv"
}
func MakeZCFZBURL(code string) string {
	return "http://quotes.money.163.com/service/" + ZCFZB + "_" + code + ".html?type=year"
}

func MakeLRBFileName(code string) string {
	return LRB + code + ".csv"
}
func MakeLRBURL(code string) string {
	return "http://quotes.money.163.com/service/" + LRB + "_" + code + ".html?type=year"
}

func MakeXJLLBFileName(code string) string {
	return XJLLB + code + ".csv"
}
func MakeXJLLBURL(code string) string {
	return "http://quotes.money.163.com/service/" + XJLLB + "_" + code + ".html?type=year"
}
