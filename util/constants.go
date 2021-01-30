package util

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	DataDirectory = "data"
	ZCFZB         = "zcfzb"
	LRB           = "lrb"
	XJLLB         = "xjllb"
	Growth        = "growth"
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

func MakeGrowthFileName(code string) string {
	return Growth + code + ".csv"
}

func MakeHumanReadFileName(code string) string {
	return filepath.Join(currentDir, DataDirectory, "humanread_"+code+".xlsx")
}

func MakeZCFZBURL(code string) string {
	return "http://quotes.money.163.com/service/" + ZCFZB + "_" + code + ".html?type=year"
}

// 成长性
func MakeGrowthURL(code string) string {
	return "http://quotes.money.163.com/service/zycwzb_" + code + ".html?type=year&part=cznl"
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

func GetRowIndex(columnKey string, result [][]string) int {
	//fmt.Println("GetRowIndex \n ")
	for i := 0; i < len(result); i++ {
		//	fmt.Println("all key:", result[i][0])
		if result[i][0] == columnKey {
			return i
		}
	}
	fmt.Println("找不到 关键词： ", columnKey)
	panic("找不到 关键词： ")
}
