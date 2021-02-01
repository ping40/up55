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
