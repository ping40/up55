package main

//https://blog.csdn.net/zhangcucmb/article/details/82683450
import (
	"fmt"
	"github.com/ping40/up55/util"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var (
	currentDir string
)

func main() {
	currentDir, _ = os.Getwd()
	print("currentDir:", currentDir)

	urlTest := "http://quotes.money.163.com/service/lrb_600519.html?type=year"
	//传递的参数:
	//dataReader := strings.NewReader("")

	newReq, err := http.NewRequest("GET", urlTest, nil)

	if err != nil {
		return
	}
	//	newReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//	newReq.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := http.DefaultClient.Do(newReq)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//打印返回值
	fmt.Println(string(body))

	dirName := filepath.Join(currentDir, "data")
	fileName := filepath.Join(dirName, "a.csv")
	util.CreateDirIfNotExists(dirName)
	util.Write(fileName, string(body))

}
