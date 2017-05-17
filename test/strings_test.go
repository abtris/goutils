package test

import (
	"testing"

	"github.com/printfhome/goutils"
)

func TestURLEncode(t *testing.T) {
	testStr := "999ss999asf2323323asd"

	out := goutils.StringsReplaceAllNonNumeric(testStr)
	if out != 9999992323323 {
		t.Error(out)
	}

	t.Log(out)
}
