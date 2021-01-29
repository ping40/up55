package analyse

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func genROE(f *excelize.File, columns int, zcfzbResult [][]string, lrbResult [][]string, valMap map[string][]string) {
	//第一行
	myC := NewColumn()
	f.SetCellValue("Sheet1", myC.String(1), "时间数据")
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(1), zcfzbResult[0][i+1])
	}

	//第二行 总资产
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(2), ZZC)
	rIndex := getRowIndex(ZZC, zcfzbResult)
	zzcList := zcfzbResult[rIndex]
	valMap[ZZC] = zzcList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(2), zzcList[i+1])
	}

	//第三行 总净利润
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(3), ZJLR)
	rIndex = getRowIndex(ZJLR, lrbResult)
	zjlrList := lrbResult[rIndex]
	valMap[ZJLR] = zjlrList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(3), zjlrList[i+1])
	}

	//第四行 总roe
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(4), "总roe")
	for i := 0; i < columns; i++ {
		roe := getAROE(zzcList[i+1], zjlrList[i+1])
		fmt.Println("roe:", roe)
		f.SetCellValue("Sheet1", myC.String(4), roe)
	}
	// 第5行 空行

	// 第6行 少数股东权益(万元)
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(6), XSGDQY)
	rIndex = getRowIndex(XSGDQY, zcfzbResult)
	xsgdqyList := zcfzbResult[rIndex]
	valMap[XSGDQY] = xsgdqyList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(6), xsgdqyList[i+1])
	}

	//第7行 少数股东损益(万元)
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(7), XSGDSY)
	rIndex = getRowIndex(XSGDSY, lrbResult)
	xsgdsyList := lrbResult[rIndex]
	valMap[XSGDSY] = xsgdsyList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(7), xsgdsyList[i+1])
	}

	//第8行 少数股东roe
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(8), "少数股东roe")
	for i := 0; i < columns; i++ {
		roe := getAROE(xsgdqyList[i+1], xsgdsyList[i+1])
		fmt.Println("roe:", roe)
		f.SetCellValue("Sheet1", myC.String(8), roe)
	}

	//MYTOTAL = "归属于母公司股东权益合计(万元)"

	// 第9行 空行

	// 第10行 归属于母公司股东权益合计(万元)
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(10), MYTOTAL)
	rIndex = getRowIndex(MYTOTAL, zcfzbResult)
	myTotalList := zcfzbResult[rIndex]
	valMap[MYTOTAL] = myTotalList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(10), myTotalList[i+1])
	}

	// 第11行 归属于母公司所有者的净利润(万元)
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(11), MYROE)
	rIndex = getRowIndex(MYROE, lrbResult)
	myRoeList := lrbResult[rIndex]
	valMap[MYROE] = myRoeList
	for i := 0; i < columns; i++ {
		f.SetCellValue("Sheet1", myC.String(11), myRoeList[i+1])
	}

	//第12行 股东roe
	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(12), "股东roe")
	for i := 0; i < columns; i++ {
		roe := getAROE(myTotalList[i+1], myRoeList[i+1])
		fmt.Println("roe:", roe)
		f.SetCellValue("Sheet1", myC.String(12), roe)
	}

	myC = NewColumn()
	f.SetCellValue("Sheet1", myC.String(13), "观察点：股东roe，和少数股东roe 大小的差别")

}

func getAROE(s1 string, s2 string) string {
	s1f, e1 := strconv.ParseFloat(s1, 10)
	s2f, e2 := strconv.ParseFloat(s2, 10)
	if e1 == nil && e2 == nil && s1f > 0 {
		return fmt.Sprintf("%.2f%%", 100*s2f/s1f)
	}
	return "--"
}
