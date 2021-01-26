package analyse

import (
	"fmt"
	"github.com/ping40/up55/util"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	zzc = "资产总计(万元)"
)

func genHumanRead(code string, zcfzbResult, lrbResult, xjllbResult [][]string) {
	columns := len(zcfzbResult[0]) - 1

	fullFileName := util.MakeHumanReadFileName(code)

	f := excelize.NewFile()
	f.SetColWidth("Sheet1", "A", "Z", 15)

	//第一行
	myC := NewColumn()
	f.SetCellValue("Sheet1", myC.String(1), "时间数据")
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(1), zcfzbResult[0][i+1])
	}

	//第二行 总资产
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(2), zzc)
	rIndex := getRowIndex(zzc, zcfzbResult)

	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(2), zcfzbResult[rIndex][i+1])
	}

	// Set active sheet of the workbook.

	// Save spreadsheet by the given path.
	if err := f.SaveAs(fullFileName); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func getRowIndex(columnKey string, result [][]string) int {
	for i := 0; i < len(result); i++ {
		fmt.Println("all key:", result[i][0])
		if result[i][0] == columnKey {
			return i
		}
	}
	fmt.Println("找不到 关键词： ", columnKey)
	panic("找不到 关键词： ")
}
