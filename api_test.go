package bankcard

import (
	"testing"
)

func TestFindBank(t *testing.T) {
	//b, err := FindBank("527594112231")
	b, err := FindBank("6222000200124846494")
	if err != nil {
		t.Error(err)
	}
	t.Log(b)
}
