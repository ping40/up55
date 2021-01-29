package analyse

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	DEBT_TOTAL = "负债合计(万元)"
)

func genDebt(f *excelize.File, columns int, zcfzbResult [][]string, lrbResult [][]string, valMap map[string][]string, row int) {
	//DEBT_TOTAL = "负债合计(万元)"
	myC := NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), DEBT_TOTAL)
	rIndex := getRowIndex(DEBT_TOTAL, zcfzbResult)
	debtTotalList := zcfzbResult[rIndex]
	valMap[DEBT_TOTAL] = debtTotalList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), debtTotalList[i+1])
	}

}
