package analyse

import (
	"testing"
)

func TestNew(t *testing.T) {
	n := NewColumn()

	t.Log(n.String())
	for i := 0; i < 150; i++ {

		t.Log(i+1, n.Next().String())
	}
}
