package fetch

import (
	"fmt"
	"time"

	"github.com/ping40/up55/util"
)

//http://quotes.money.163.com/service/zcfzb_600519.html?type=year

func Fetch(code string) {
	if util.CodeFileExists(code) {
		fmt.Println("文件存在，", code)
		return
	}

	fmt.Println("文件不存在，", code)
	dataDir := util.MakeDataDirectory()

	url := util.MakeZCFZBURL(code)
	fileName := util.MakeZCFZBFileName(code)
	if err := util.DownloadAndSave(url, dataDir, fileName); err != nil {
		fmt.Println("fetch ", fileName, "failed ")
		panic(err)
	}
	time.Sleep(time.Second * 2)

	url = util.MakeLRBURL(code)
	fileName = util.MakeLRBFileName(code)
	if err := util.DownloadAndSave(url, dataDir, fileName); err != nil {
		fmt.Println("fetch ", fileName, "failed ")
		panic(err)
	}
	time.Sleep(time.Second * 2)

	url = util.MakeXJLLBURL(code)
	fileName = util.MakeXJLLBFileName(code)
	if err := util.DownloadAndSave(url, dataDir, fileName); err != nil {
		fmt.Println("fetch ", fileName, "failed ")
		panic(err)
	}
}
