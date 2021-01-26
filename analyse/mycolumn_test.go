package analyse

import (
	"testing"
)

func TestNew(t *testing.T) {
	n := NewColumn()

	for i := 0; i < 150; i++ {
		t.Log(i+1, n.String(1))
	}
}
