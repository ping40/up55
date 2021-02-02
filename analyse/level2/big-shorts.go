package level2

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"github.com/ping40/up55/util"
)

const ()

func GenBigShorts(f *excelize.File, row int) {

	f.SetRowHeight("Sheet1", row, 70)
	f.MergeCell("Sheet1", fmt.Sprintf("A%d", row), fmt.Sprintf("G%d", row))
	myC := util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), `财务健康状况：
空头情形：公司的客户的财务健康状况，如果客户崩溃了，就完了
`)

}
