package main

//https://blog.csdn.net/zhangcucmb/article/details/82683450
import (
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
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
		// handle error
	}
	//打印返回值
	fmt.Println(string(body))

}
