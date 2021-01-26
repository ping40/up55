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
	columns := len(zcfzbResult)

	fullFileName := util.MakeHumanReadFileName(code)

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", columns)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(fullFileName); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
