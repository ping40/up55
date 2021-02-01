package level2

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"github.com/ping40/up55/util"
)

const (
	GROWTH_TOTAL_INCOME = "主营业务收入增长率(%)"
	GROWTH_NET_PROFIT   = "净利润增长率(%)"
	GROWTH_NET_ASSET    = "净资产增长率(%)"
	GROWTH_ASSET        = "总资产增长率(%)"

	GROWTH_NET_MONEY        = " 期末现金及现金等价物余额(万元)"
	GROWTH_INCOME_PER_SHARE = "基本每股收益"
)

func GenGrowth(f *excelize.File, columns int, xjllbResult, growthResult, lrbResult [][]string, row int) {
	// GROWTH_TOTAL_INCOME ="主营业务收入增长率(%)"
	myC := util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), GROWTH_TOTAL_INCOME)
	rIndex := util.GetRowIndex(GROWTH_TOTAL_INCOME, growthResult)
	growthTotalIncomeList := growthResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), growthTotalIncomeList[i+1])
	}

	//GROWTH_NET_PROFIT ="净利润增长率(%)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), GROWTH_NET_PROFIT)
	rIndex = util.GetRowIndex(GROWTH_NET_PROFIT, growthResult)
	growthNetProfitList := growthResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), growthNetProfitList[i+1])
	}

	//GROWTH_NET_ASSET    = "净资产增长率(%)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), GROWTH_NET_ASSET)
	rIndex = util.GetRowIndex(GROWTH_NET_ASSET, growthResult)
	growthNetAssetList := growthResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), growthNetAssetList[i+1])
	}

	//GROWTH_ASSET        = "资产增长率(%)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), GROWTH_ASSET)
	rIndex = util.GetRowIndex(GROWTH_ASSET, growthResult)
	growthAssetList := growthResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), growthAssetList[i+1])
	}

	//GROWTH_INCOME_PER_SHARE = "基本每股收益"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), GROWTH_INCOME_PER_SHARE)
	rIndex = util.GetRowIndex(GROWTH_INCOME_PER_SHARE, lrbResult)
	growthIncomePerShareList := lrbResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), growthIncomePerShareList[i+1])
	}

	//GROWTH_NET_MONEY        = "期末现金及现金等价物余额(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), GROWTH_NET_MONEY)
	rIndex = util.GetRowIndex(GROWTH_NET_MONEY, xjllbResult)
	growthNetMoenyList := xjllbResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), growthNetMoenyList[i+1])
	}

	row++
	f.SetRowHeight("Sheet1", row, 200)
	f.MergeCell("Sheet1", fmt.Sprintf("A%d", row), fmt.Sprintf("M%d", row))
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), `公司的成长性观察：销售增长来源：
1：销售更多的产品或服务（如啤酒整体市场增加或者本公司份额增加）
2：提高价格
3：销售新的产品或服务
4：购买其他公司 （危险）

同一个产品或服务增长：要么价格提升，要么行业规模增加，要么本公司份额增加

销售增长很难伪造的，盈利增长的欺骗方法很多，如改变税率，股份数，一次性所得，削减成本
一般情况，公司赢利增长超过销售增长一段时间，也是值得怀疑的
`)

}
