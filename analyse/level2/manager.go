package level2

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"github.com/ping40/up55/util"
)

const ()

func GenManagement(f *excelize.File, row int) {

	f.SetRowHeight("Sheet1", row, 70)
	f.MergeCell("Sheet1", fmt.Sprintf("A%d", row), fmt.Sprintf("G%d", row))
	myC := util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), `管理：
1：薪酬：高额薪酬不如高额红利，慷慨期权计划不如严格限制的股票期权。 红利意味着风险
2：性格：让亲朋好友富裕起来（关联交易），要求诚实
3：

`)

}
