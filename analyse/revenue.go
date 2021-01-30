package analyse

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	REVENUE_TOTAL = "营业总收入(万元)"
	COST_SALE     = "销售费用(万元)"
	COST_MANAGE   = "管理费用(万元)"
	COST_FIN      = "财务费用(万元)"
	COST_RD       = "研发费用(万元)"
)

func genRevenue(f *excelize.File, columns int, zcfzbResult [][]string, lrbResult [][]string, valMap map[string][]string, row int) {
	//DEBT_TOTAL = "负债合计(万元)"
	myC := NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), REVENUE_TOTAL)
	rIndex := getRowIndex(REVENUE_TOTAL, lrbResult)
	revenueTotalList := lrbResult[rIndex]
	valMap[REVENUE_TOTAL] = revenueTotalList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), revenueTotalList[i+1])
	}

	//	COST_SALE     = "销售费用(万元)"
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), COST_SALE)
	rIndex = getRowIndex(COST_SALE, lrbResult)
	costSaleList := lrbResult[rIndex]
	valMap[COST_SALE] = costSaleList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), costSaleList[i+1])
	}

	//	COST_MANAGE   = "管理费用(万元)"
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), COST_MANAGE)
	rIndex = getRowIndex(COST_MANAGE, lrbResult)
	costManageList := lrbResult[rIndex]
	valMap[COST_MANAGE] = costManageList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), costManageList[i+1])
	}

	//	COST_FIN      = "财务费用(万元)"
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), COST_FIN)
	rIndex = getRowIndex(COST_FIN, lrbResult)
	costFinList := lrbResult[rIndex]
	valMap[COST_FIN] = costFinList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), costFinList[i+1])
	}

	//	COST_RD       = "研发费用(万元)"
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), COST_RD)
	rIndex = getRowIndex(COST_RD, lrbResult)
	costRDList := lrbResult[rIndex]
	valMap[COST_RD] = costRDList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), costRDList[i+1])
	}
}
