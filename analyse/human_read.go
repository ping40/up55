package analyse

import (
	"fmt"
	"time"

	"github.com/ping40/up55/util"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	ZZC     = "资产总计(万元)"
	ZJLR    = "净利润(万元)"
	XSGDQY  = "少数股东权益(万元)"
	XSGDSY  = "少数股东损益(万元)"
	MYTOTAL = "归属于母公司股东权益合计(万元)"
	MYROE   = "归属于母公司所有者的净利润(万元)"
	MYMPS   = "基本每股收益"
)

func genHumanRead(code string, zcfzbResult, lrbResult, xjllbResult [][]string) {
	columns := len(zcfzbResult[0]) - 1

	fullFileName := util.MakeHumanReadFileName(fmt.Sprintf("%s-%d", code, time.Now().UnixNano()))

	f := excelize.NewFile()
	f.SetColWidth("Sheet1", "A", "Z", 14)
	valMap := map[string][]string{}

	genROE(f, columns, zcfzbResult, lrbResult, valMap)
	genShare(f, columns, zcfzbResult, lrbResult, 15)
	// Set active sheet of the workbook.

	// Save spreadsheet by the given path.
	if err := f.SaveAs(fullFileName); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func genShare(f *excelize.File, columns int, zcfzbResult [][]string, lrbResult [][]string, row int) {

	//第row行 基本每股收益
	myC := NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), MYMPS)
	rIndex := getRowIndex(MYMPS, lrbResult)
	mympsList := lrbResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), mympsList[i+1])
	}

	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row+1), "分红")

}

func getRowIndex(columnKey string, result [][]string) int {
	//fmt.Println("getRowIndex \n ")
	for i := 0; i < len(result); i++ {
		//	fmt.Println("all key:", result[i][0])
		if result[i][0] == columnKey {
			return i
		}
	}
	fmt.Println("找不到 关键词： ", columnKey)
	panic("找不到 关键词： ")
}
