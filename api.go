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
	if len(card) < 5 || !goutil.StringIs(card, unicode.IsDigit) {
		return nil, fmt.Errorf("format error")
	}
	i, err := strconv.Atoi(card[:5])
	if err != nil {
		return nil, err
	}
	bank := Index[i]
	if bank == "" {
		if len(card) >= 6 {
			if i, err = strconv.Atoi(card[:6]); err != nil {
				return nil, err
			}
			bank = Index[i]
		}
	}
	if bank == "" {
		return nil, fmt.Errorf("not found")
	}
	return Banks[bank], nil
}
