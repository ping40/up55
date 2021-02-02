package level2

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"github.com/ping40/up55/util"
)

const ()

func GenRepayAbility(f *excelize.File, columns int, repayingAbilityResult [][]string, row int) {

	for _, v := range repayingAbilityResult {
		h := strings.TrimSpace(v[0])
		if h == "" {
			break
		}

		myC := util.NewColumn()
		f.SetCellValue("Sheet1", myC.String(row), v[0])
		for i := 0; i < len(v)-1; i++ {
			f.SetCellValue("Sheet1", myC.String(row), v[i+1])
		}

		row++
	}

	f.SetRowHeight("Sheet1", row, 70)
	f.MergeCell("Sheet1", fmt.Sprintf("A%d", row), fmt.Sprintf("G%d", row))
	myC := util.NewColumn()
	f.SetCellValue("Sheet1", myC.String(row), `财务健康状况：
流动速率：  流动资产/流动负债， > 1.5
速动比率：  (流动资产-库存)/流动负债, >1.0
`)

}
