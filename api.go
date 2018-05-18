package bankcard

import (
	"fmt"
	"github.com/liuzl/goutil"
	"strconv"
	"unicode"
)

var Banks = make(map[string]*Bank)
var Index = make(map[int]string)

func FindBank(card string) (*Bank, error) {
	if len(card) < MaxLen || !goutil.StringIs(card, unicode.IsDigit) {
		return nil, fmt.Errorf("format error")
	}

	var bank string
	for i := MaxLen; i >= 1; i-- {
		i, err := strconv.Atoi(card[:i])
		if err != nil {
			return nil, err
		}
		if bank = Index[i]; bank != "" {
			break
		}
	}
	if bank == "" {
		return nil, fmt.Errorf("not found")
	}
	return Banks[bank], nil
}
