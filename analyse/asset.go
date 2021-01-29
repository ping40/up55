package analyse

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	ASSET_FLOW       = "流动资产合计(万元)"
	ASSET_MONEY      = "货币资金(万元)"
	ASSET_INVISIBLE  = "无形资产(万元)"
	ASSSET_WILL_GET  = "应收" // 名称中有 应收，但是不包含 长期应收的情况
	ASSSET_COMMODITY = "存货(万元)"
	ASSSET_HAVE_GOT  = "预付款项(万元)"

	ASSET_NO_FLOW       = "非流动资产合计(万元)"
	ASSET_FIXED         = "固定资产(万元)"
	ASSET_REPUTATION    = "商誉(万元)"
	ASSET_LONG_WILL_GET = "长期应收款(万元)"
)

func genAsset(f *excelize.File, columns int, zcfzbResult [][]string, lrbResult [][]string, valMap map[string][]string, row int) {
	myC := NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ASSET_FLOW)
	rIndex := getRowIndex(ASSET_FLOW, zcfzbResult)
	assetFlowList := zcfzbResult[rIndex]
	valMap[ASSET_FLOW] = assetFlowList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), assetFlowList[i+1])
	}

	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ASSET_MONEY)
	rIndex = getRowIndex(ASSET_MONEY, zcfzbResult)
	assetMoneyList := zcfzbResult[rIndex]
	valMap[ASSET_MONEY] = assetMoneyList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), assetMoneyList[i+1])
	}

	// 应收
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ASSSET_WILL_GET)
	assetWillGetList := getWillGetList(zcfzbResult)
	//valMap[ASSSET_WILL_GET] = assetInvisibleList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), strconv.FormatInt(assetWillGetList[i], 10))
	}

	//"存货(万元)"
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ASSSET_COMMODITY)
	rIndex = getRowIndex(ASSSET_COMMODITY, zcfzbResult)
	assetCommodityList := zcfzbResult[rIndex]
	valMap[ASSSET_COMMODITY] = assetMoneyList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), assetCommodityList[i+1])
	}

	//	"预付款项(万元)"
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ASSSET_HAVE_GOT)
	rIndex = getRowIndex(ASSSET_HAVE_GOT, zcfzbResult)
	assetHaveGotList := zcfzbResult[rIndex]
	valMap[ASSSET_HAVE_GOT] = assetHaveGotList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), assetHaveGotList[i+1])
	}

	// 非流动资产合计(万元)
	row++
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ASSET_NO_FLOW)
	rIndex = getRowIndex(ASSET_NO_FLOW, zcfzbResult)
	assetNoFlowList := zcfzbResult[rIndex]
	valMap[ASSET_NO_FLOW] = assetNoFlowList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), assetNoFlowList[i+1])
	}

	//"固定资产(万元)"
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ASSET_FIXED)
	rIndex = getRowIndex(ASSET_FIXED, zcfzbResult)
	assetfixedList := zcfzbResult[rIndex]
	valMap[ASSET_FIXED] = assetfixedList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), assetfixedList[i+1])
	}

	//无形资产
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ASSET_INVISIBLE)
	rIndex = getRowIndex(ASSET_INVISIBLE, zcfzbResult)
	assetInvisibleList := zcfzbResult[rIndex]
	valMap[XSGDQY] = assetInvisibleList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), assetInvisibleList[i+1])
	}

	//ASSET_REPUTATION = "商誉(万元)"
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ASSET_REPUTATION)
	rIndex = getRowIndex(ASSET_REPUTATION, zcfzbResult)
	assetReputationList := zcfzbResult[rIndex]
	valMap[ASSET_REPUTATION] = assetReputationList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), assetReputationList[i+1])
	}

	//ASSET_LONG_WILL_GET = "长期应收款(万元)"
	row++
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), ASSET_LONG_WILL_GET)
	rIndex = getRowIndex(ASSET_LONG_WILL_GET, zcfzbResult)
	assetLongWillGetList := zcfzbResult[rIndex]
	valMap[ASSET_LONG_WILL_GET] = assetLongWillGetList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(row), assetLongWillGetList[i+1])
	}

}

//应收  集合
func getWillGetList(result [][]string) []int64 {
	//fmt.Println("getRowIndex \n ")
	var indices []int
	for i := 0; i < len(result); i++ {
		//	fmt.Println("all key:", result[i][0])
		if strings.Contains(result[i][0], "应收") && !strings.Contains(result[i][0], "长期应收") {
			indices = append(indices, i)
		}
	}
	if len(indices) == 0 {
		fmt.Println("找不到 应收记录")
		panic("找不到 应收记录 ")
	}

	var list []int64
	for i := 1; i < len(result[0]); i++ {
		var r int64 = 0
		for _, k := range indices {
			r += getInt64(result[k][i])
		}
		list = append(list, r)

	}

	return list

}
