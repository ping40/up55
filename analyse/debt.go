package analyse

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/ping40/up55/util"
)

const (
	DEBT_TOTAL            = "负债合计(万元)"
	DEBT_SHORT            = "短期借款(万元)"
	DEBT_LONG             = "长期借款(万元)"
	DEBT_WILL_PAY_INVOICE = "应付票据(万元)"
	DEBT_WILL_PAY_MONEY   = "应付账款(万元)"
)

func genDebt(f *excelize.File, columns int, zcfzbResult [][]string, lrbResult [][]string, valMap map[string][]string, row int) {
	//DEBT_TOTAL = "负债合计(万元)"
	myC := util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), DEBT_TOTAL)
	rIndex := util.GetRowIndex(DEBT_TOTAL, zcfzbResult)
	debtTotalList := zcfzbResult[rIndex]
	valMap[DEBT_TOTAL] = debtTotalList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), debtTotalList[i+1])
	}

	//DEBT_SHORT = "短期借款(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), DEBT_SHORT)
	rIndex = util.GetRowIndex(DEBT_SHORT, zcfzbResult)
	debtShortList := zcfzbResult[rIndex]
	valMap[DEBT_SHORT] = debtShortList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), debtShortList[i+1])
	}

	//DEBT_LONG = "长期借款(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), DEBT_LONG)
	rIndex = util.GetRowIndex(DEBT_LONG, zcfzbResult)
	debtLongList := zcfzbResult[rIndex]
	valMap[DEBT_LONG] = debtLongList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), debtLongList[i+1])
	}

	//	DEBT_WILL_PAY = "应付票据及应付账款"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), DEBT_WILL_PAY_INVOICE)
	rIndex = util.GetRowIndex(DEBT_WILL_PAY_INVOICE, zcfzbResult)
	debtWillPayList := zcfzbResult[rIndex]
	valMap[DEBT_WILL_PAY_INVOICE] = debtWillPayList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), debtWillPayList[i+1])
	}

	//	DEBT_WILL_PAY_MONEY   = "应付账款(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), DEBT_WILL_PAY_MONEY)
	rIndex = util.GetRowIndex(DEBT_WILL_PAY_MONEY, zcfzbResult)
	debtWillPayMoneyList := zcfzbResult[rIndex]
	valMap[DEBT_WILL_PAY_MONEY] = debtWillPayMoneyList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), debtWillPayMoneyList[i+1])
	}

}
