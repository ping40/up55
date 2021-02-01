package level2

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	_ "image/png"
	"io/ioutil"

	"github.com/ping40/up55/util"
)

const (
	PROFIT_ROE             = "净资产收益率(%)"
	PROFIT_ALL_ASSET       = "总资产净利润率(%)"
	PROFIT_ALL_ASSET_ROUND = "总资产周转率(次)"
	PROFIT_SALES           = "销售净利率(%)" //报表给的数据有错误的
	PROFIT_SALES_1         = "归属于母公司所有者的净利润(万元)"
	PROFIT_SALES_2         = "营业总收入(万元)"
	PROFIT_ADDED           = "现金及现金等价物净增加额(万元)"

	PROFIT_MAIN_TOTAL       = "经营活动现金流入小计(万元)"
	PROFIT_MAIN_TOTAL_ADDED = "经营活动产生的现金流量净额(万元)"
)

func GenProfit(f *excelize.File, columns int, profitAbilityResult, operationAbilityResult, lrbResult, xjllbResult [][]string, row int) {
	//PROFIT_ROE = "净资产收益率(%)"
	myC := util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), PROFIT_ROE)
	rIndex := util.GetRowIndex(PROFIT_ROE, profitAbilityResult)
	profitROEList := profitAbilityResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), profitROEList[i+1])
	}

	//PROFIT_ALL_ASSET = "总资产净利润率(%)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), PROFIT_ALL_ASSET)
	rIndex = util.GetRowIndex(PROFIT_ALL_ASSET, profitAbilityResult)
	profitAllAssetList := profitAbilityResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), profitAllAssetList[i+1])
	}

	//权益乘数
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "权益乘数")
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), util.GenPercent(profitAllAssetList[i+1], profitROEList[i+1]))
	}

	// PROFIT_ALL_ASSET_ROUND = "总资产周转率(次)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), PROFIT_ALL_ASSET_ROUND)
	rIndex = util.GetRowIndex(PROFIT_ALL_ASSET_ROUND, operationAbilityResult)
	profitAllAssetRoundList := operationAbilityResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), profitAllAssetRoundList[i+1])
	}

	//PROFIT_SALES           = "销售净利率(%)" //报表给的数据有错误的
	//PROFIT_SALES_1         = "归属于母公司所有者的净利润(万元)"
	//PROFIT_SALES_2         = "营业总收入(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), PROFIT_SALES)
	rIndex = util.GetRowIndex(PROFIT_SALES_1, lrbResult)
	PROFIT_SALES_1_List := lrbResult[rIndex]
	rIndex = util.GetRowIndex(PROFIT_SALES_2, lrbResult)
	PROFIT_SALES_2_List := lrbResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), util.GenPercent(PROFIT_SALES_2_List[i+1], PROFIT_SALES_1_List[i+1]))
	}

	//PROFIT_ADDED = "现金及现金等价物净增加额(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "自由现金流比率")
	rIndex = util.GetRowIndex(PROFIT_ADDED, xjllbResult)
	profitAddedList := xjllbResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), util.GenPercent(PROFIT_SALES_2_List[i+1], profitAddedList[i+1]))
	}

	//PROFIT_MAIN_TOTAL       = "销售商品、提供劳务收到的现金(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), PROFIT_MAIN_TOTAL)
	rIndex = util.GetRowIndex(PROFIT_MAIN_TOTAL, xjllbResult)
	profitMainTotal := xjllbResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), profitMainTotal[i+1])
	}

	//	PROFIT_MAIN_TOTAL_ADDED = "经营活动产生的现金流量净额(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), PROFIT_MAIN_TOTAL_ADDED)
	rIndex = util.GetRowIndex(PROFIT_MAIN_TOTAL_ADDED, xjllbResult)
	profitMainTotalAdded := xjllbResult[rIndex]
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), profitMainTotalAdded[i+1])
	}

	//PROFIT_ADDED = "现金及现金等价物净增加额(万元)"
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "经营活动产生的自由现金流比率")
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), util.GenPercent(profitMainTotal[i+1], profitMainTotalAdded[i+1]))
	}

	row++
	f.SetRowHeight("Sheet1", row, 300)
	f.MergeCell("Sheet1", fmt.Sprintf("A%d", row), fmt.Sprintf("G%d", row))
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), `公司的收益性观察：
自由现金流比率=现金及现金等价物净增加额(万元)/营业收入
自由现金流比率 >= 5%, roe> 15%: 优秀企业
2014~2019， roe>15%, 520家， >10%, 1300家
2014~2019， 自由现金流比率>10%, 440家

低风险公司，
`)

	f.MergeCell("Sheet1", fmt.Sprintf("H%d", row), fmt.Sprintf("M%d", row))
	file, err := ioutil.ReadFile("./img/shouyilv.png")
	if err != nil {
		panic(err)
	}

	if err = f.AddPictureFromBytes("Sheet1", fmt.Sprintf("H%d", row), `{
        "x_scale": 0.7,
        "y_scale": 0.7      
    }`, "excel logo", ".png", file); err != nil {
		panic(err)
	}
}
