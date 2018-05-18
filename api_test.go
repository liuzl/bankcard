package bankcard

import (
	"testing"
)

func TestFindBank(t *testing.T) {
	b, err := FindBank("527594112231")
	if err != nil {
		t.Error(err)
	}
	t.Log(b)
}
