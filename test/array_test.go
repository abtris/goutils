package test

import (
	"testing"

	"github.com/printfhome/goutils"
)

func TestArrayRemoveItems(t *testing.T) {

	in := []string{"1", "2", "3", "4", "4", "3", "78", "u"}
	items1 := []string{"1", "2", "3", "4", "4", "3", "78", "u"}
	items2 := []string{"3", "4", "4"}

	out := goutils.ArrayStringMoveItems(in, items1...)
	if len(out) != 0 {
		t.Error(out)
	}
	out = goutils.ArrayStringMoveItems(in, items2...)
	if out[0] != "1" || out[1] != "2" || out[2] != "78" || out[3] != "u" {
		t.Error(out)
	}

	t.Log(out)
}
