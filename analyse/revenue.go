package analyse

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/ping40/up55/util"
)

const (
	REVENUE_TOTAL   = "营业总收入(万元)"
	COST_SALE       = "销售费用(万元)"
	COST_MANAGE     = "管理费用(万元)"
	COST_FIN        = "财务费用(万元)"
	COST_RD         = "研发费用(万元)"
	COST_ALL_PROFIT = "营业利润(万元)"
	COST_NET_PROFIT = "净利润(万元)"
	REVENUE_COST    = "营业总成本(万元)"
)

func genRevenue(f *excelize.File, columns int, zcfzbResult [][]string, lrbResult [][]string, valMap map[string][]string, row int) {
	//DEBT_TOTAL = "负债合计(万元)"
	myC := util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), REVENUE_TOTAL)
	rIndex := util.GetRowIndex(REVENUE_TOTAL, lrbResult)
	revenueTotalList := lrbResult[rIndex]
	valMap[REVENUE_TOTAL] = revenueTotalList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), revenueTotalList[i+1])
	}

	//	REVENUE_COST  = "营业总成本(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), REVENUE_COST)
	rIndex = util.GetRowIndex(REVENUE_COST, lrbResult)
	revenueCostList := lrbResult[rIndex]
	valMap[COST_SALE] = revenueCostList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), revenueCostList[i+1])
	}

	//	COST_SALE     = "销售费用(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), COST_SALE)
	rIndex = util.GetRowIndex(COST_SALE, lrbResult)
	costSaleList := lrbResult[rIndex]
	valMap[COST_SALE] = costSaleList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), costSaleList[i+1])
	}

	//	COST_MANAGE   = "管理费用(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), COST_MANAGE)
	rIndex = util.GetRowIndex(COST_MANAGE, lrbResult)
	costManageList := lrbResult[rIndex]
	valMap[COST_MANAGE] = costManageList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), costManageList[i+1])
	}

	//	COST_FIN      = "财务费用(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), COST_FIN)
	rIndex = util.GetRowIndex(COST_FIN, lrbResult)
	costFinList := lrbResult[rIndex]
	valMap[COST_FIN] = costFinList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), costFinList[i+1])
	}

	//	COST_RD       = "研发费用(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), COST_RD)
	rIndex = util.GetRowIndex(COST_RD, lrbResult)
	costRDList := lrbResult[rIndex]
	valMap[COST_RD] = costRDList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), costRDList[i+1])
	}

	//	COST_ALL_PROFIT = "营业利润(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), COST_ALL_PROFIT)
	rIndex = util.GetRowIndex(COST_ALL_PROFIT, lrbResult)
	costAllProfitList := lrbResult[rIndex]
	valMap[COST_ALL_PROFIT] = costRDList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), costAllProfitList[i+1])
	}

	//	COST_NET_PROFIT = "净利润(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), COST_NET_PROFIT)
	rIndex = util.GetRowIndex(COST_NET_PROFIT, lrbResult)
	costNetProfitList := lrbResult[rIndex]
	valMap[COST_NET_PROFIT] = costRDList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), costNetProfitList[i+1])
	}

	//	营业成本%"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "营业成本%")
	for i := 0; i < columns; i++ {
		roe := util.GenPercent(revenueTotalList[i+1], revenueCostList[i+1])
		f.SetCellValue("Sheet1", myC.String(row), roe)
	}

	//	销售成本%"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "销售成本%")
	for i := 0; i < columns; i++ {
		roe := util.GenPercent(revenueTotalList[i+1], costSaleList[i+1])
		f.SetCellValue("Sheet1", myC.String(row), roe)
	}

	//	管理费用%"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "管理费用%")
	for i := 0; i < columns; i++ {
		roe := util.GenPercent(revenueTotalList[i+1], costManageList[i+1])
		f.SetCellValue("Sheet1", myC.String(row), roe)
	}

	//	研发费用%"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "研发费用%")
	for i := 0; i < columns; i++ {
		roe := util.GenPercent(revenueTotalList[i+1], costRDList[i+1])
		f.SetCellValue("Sheet1", myC.String(row), roe)
	}

	//	财务费用%"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "财务费用%")
	for i := 0; i < columns; i++ {
		roe := util.GenPercent(revenueTotalList[i+1], costFinList[i+1])
		f.SetCellValue("Sheet1", myC.String(row), roe)
	}

	//	营业利润%"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "营业利润%")
	for i := 0; i < columns; i++ {
		roe := util.GenPercent(revenueTotalList[i+1], costAllProfitList[i+1])
		f.SetCellValue("Sheet1", myC.String(row), roe)
	}

	//	净利润%"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "净利润%")
	for i := 0; i < columns; i++ {
		roe := util.GenPercent(revenueTotalList[i+1], costNetProfitList[i+1])
		f.SetCellValue("Sheet1", myC.String(row), roe)
	}

	//	营业收入年增加%"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "营业收入年增加%")
	for i := 0; i < columns-1; i++ {
		roe := util.GenGowth(revenueTotalList[i+1], revenueTotalList[i+2])
		f.SetCellValue("Sheet1", myC.String(row), roe)
	}
}
