package analyse

import (
	"fmt"
	"time"

	"github.com/ping40/up55/analyse/level2"

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

func genHumanRead(code string, zcfzbResult, lrbResult, xjllbResult, growthResult, profitAbilityResult, repayingAbilityResult, operationAbilityResult [][]string) {
	columns := len(zcfzbResult[0]) - 1

	fullFileName := util.MakeHumanReadFileName(fmt.Sprintf("%s-%d", code, time.Now().UnixNano()))

	f := excelize.NewFile()
	f.SetColWidth("Sheet1", "A", "Z", 14)
	valMap := map[string][]string{}

	row := 1
	genROE(f, columns, zcfzbResult, lrbResult, profitAbilityResult, valMap, row)

	row = 15
	genShare(f, columns, zcfzbResult, lrbResult, row)
	// Set active sheet of the workbook.

	//资产
	row = 20
	genAsset(f, columns, zcfzbResult, lrbResult, valMap, row)

	//资产
	row = 33
	genDebt(f, columns, zcfzbResult, lrbResult, valMap, row)

	//营业收入
	row = 40
	genRevenue(f, columns, zcfzbResult, lrbResult, valMap, row)

	//成长性
	row = 50
	level2.GenGrowth(f, columns, xjllbResult, growthResult, lrbResult, row)

	//收益性
	row = 60
	level2.GenProfit(f, columns, profitAbilityResult, operationAbilityResult, lrbResult, xjllbResult, row)

	//财务健康风险
	row = 80
	level2.GenRepayAbility(f, columns, repayingAbilityResult, row)

	//财务健康风险
	row = 100
	level2.GenBigShorts(f, row)

	// Save spreadsheet by the given path.
	if err := f.SaveAs(fullFileName); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func genShare(f *excelize.File, columns int, zcfzbResult [][]string, lrbResult [][]string, row int) {

	//第row行 基本每股收益
	myC := util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), MYMPS)
	rIndex := util.GetRowIndex(MYMPS, lrbResult)
	mympsList := lrbResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), mympsList[i+1])
	}

	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row+1), "分红")

}
