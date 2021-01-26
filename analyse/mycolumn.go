package analyse

var chars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type myColumn struct {
	val int
}

func (this *myColumn) Next() *myColumn {
	this.val += 1
	return this
}

func NewColumn() *myColumn {
	return &myColumn{}
}

func (this *myColumn) String() string {
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

	return str

}
