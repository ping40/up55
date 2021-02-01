package analyse

import (
	"fmt"
	"strconv"

	"github.com/ping40/up55/util"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	ALL_ASSET_PROFIT = "总资产净利润率(%)"
	NET_ASSET_PROFIT = "净资产收益率(%)"
)

func genROE(f *excelize.File, columns int, zcfzbResult [][]string, lrbResult, profitAbilityResult [][]string, valMap map[string][]string, row int) {
	//第一行
	myC := util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "时间数据")
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), zcfzbResult[0][i+1])
	}

	//第二行 总资产
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ZZC)
	rIndex := util.GetRowIndex(ZZC, zcfzbResult)
	zzcList := zcfzbResult[rIndex]
	valMap[ZZC] = zzcList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), zzcList[i+1])
	}

	//第三行 总净利润
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ZJLR)
	rIndex = util.GetRowIndex(ZJLR, lrbResult)
	zjlrList := lrbResult[rIndex]
	valMap[ZJLR] = zjlrList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), zjlrList[i+1])
	}

	//第四行 总roe
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ALL_ASSET_PROFIT)
	rIndex = util.GetRowIndex(ALL_ASSET_PROFIT, profitAbilityResult)
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), profitAbilityResult[rIndex][i+1])
	}
	// 第5行 空行

	// 第6行 少数股东权益(万元)
	row++
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), XSGDQY)
	rIndex = util.GetRowIndex(XSGDQY, zcfzbResult)
	xsgdqyList := zcfzbResult[rIndex]
	valMap[XSGDQY] = xsgdqyList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), xsgdqyList[i+1])
	}

	//第7行 少数股东损益(万元)
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), XSGDSY)
	rIndex = util.GetRowIndex(XSGDSY, lrbResult)
	xsgdsyList := lrbResult[rIndex]
	valMap[XSGDSY] = xsgdsyList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), xsgdsyList[i+1])
	}

	//第8行 少数股东roe
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "少数股东roe")
	for i := 0; i < columns; i++ {
		roe := util.GenPercent(xsgdqyList[i+1], xsgdsyList[i+1])
		fmt.Println("roe:", roe)
		f.SetCellValue("Sheet1", myC.String(row), roe)
	}

	//MYTOTAL = "归属于母公司股东权益合计(万元)"

	// 第9行 空行

	// 第10行 归属于母公司股东权益合计(万元)
	row++
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), MYTOTAL)
	rIndex = util.GetRowIndex(MYTOTAL, zcfzbResult)
	myTotalList := zcfzbResult[rIndex]
	valMap[MYTOTAL] = myTotalList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), myTotalList[i+1])
	}

	// 第11行 归属于母公司所有者的净利润(万元)
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), MYROE)
	rIndex = util.GetRowIndex(MYROE, lrbResult)
	myRoeList := lrbResult[rIndex]
	valMap[MYROE] = myRoeList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), myRoeList[i+1])
	}

	//第12行 股东roe
	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), NET_ASSET_PROFIT)
	rIndex = util.GetRowIndex(NET_ASSET_PROFIT, profitAbilityResult)
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), profitAbilityResult[rIndex][i+1])
	}

	row++
	myC = util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), "观察点：股东roe，和少数股东roe 大小的差别")

}

func getInt64(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		return 0
	} else {
		return i
	}
}
