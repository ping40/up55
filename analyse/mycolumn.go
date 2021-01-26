package analyse

import "strconv"

var chars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type myColumn struct {
	val int
}

func NewColumn() *myColumn {
	return &myColumn{}
}

func (this *myColumn) String(rows int) string {
	yu := this.val % 26
	k := this.val / 26
	arrs := make([]string, 0)
	arrs = append(arrs, string(chars[yu]))

	for k > 0 {
		yu = k % 26
		k = k / 26
		arrs = append(arrs, string(chars[yu-1]))
	}
	str := ""
	for end := 0; end < len(arrs); end++ {
		str = arrs[end] + str
	}
	str += strconv.Itoa(rows)

	this.val += 1

	return str

}
