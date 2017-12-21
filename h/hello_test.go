package h

import (
	"strings"
	"testing"
)

func TestPrintHello(t *testing.T) {
	arg := "Anika"
	ret := PrintHello(arg)
	expected_ret := "Hello " + arg
	if strings.Compare(ret, expected_ret) != 0 {
		t.Error("Expected \"%s\", instead got \"%s\"", expected_ret, ret)
	}
}