package analyse

import "strconv"

type myRow struct {
	val int
}

func (this *myRow) Next() *myRow {
	this.val += 1
	return this
}

func NewRow() *myRow {
	return &myRow{val: 1}
}

func (this *myRow) String() string {
	return strconv.Itoa(this.val)
}
