package util

import (
	"fmt"
	"strconv"
)

func GenPercent(base string, upNumber string) string {
	s1f, e1 := strconv.ParseFloat(base, 10)
	s2f, e2 := strconv.ParseFloat(upNumber, 10)
	if e1 == nil && e2 == nil && s1f > 0 {
		return fmt.Sprintf("%.2f%%", 100*s2f/s1f)
	}
	return "--"
}

func GenGowth(thisYear string, oldYear string) string {
	ty, e1 := strconv.ParseFloat(thisYear, 10)
	oy, e2 := strconv.ParseFloat(oldYear, 10)
	if e1 == nil && e2 == nil && oy > 0 {
		return fmt.Sprintf("%.2f%%", 100*(ty-oy)/oy)
	}
	return "--"
}
