package bankcard

import (
	"fmt"
	"github.com/liuzl/goutil"
	"strconv"
	"unicode"
)

var (
	Banks = make(map[string]*Bank)
	Index = make(map[int]string)

	BanksProMap  = make(map[string]map[string]*BankPro)
	IndexProMap  = make(map[string]map[int]string)
	MaxLenProMap = make(map[string]int)
)

func FindBank(card string) (*BankPro, error) {
	if len(card) < MaxLen || !goutil.StringIs(card, unicode.IsDigit) {
		return nil, fmt.Errorf("format error")
	}

	for k, m := range BanksProMap {
		for i := MaxLenProMap[k]; i >= 1; i-- {
			n, err := strconv.Atoi(card[:i])
			if err != nil {
				return nil, err
			}
			if bank := IndexProMap[k][n]; bank != "" {
				return m[bank], nil
			}
		}
	}

	for i := MaxLen; i >= 1; i-- {
		n, err := strconv.Atoi(card[:i])
		if err != nil {
			return nil, err
		}
		if bank := Index[n]; bank != "" {
			return &BankPro{Bank: *Banks[bank]}, nil
		}
	}
	return nil, fmt.Errorf("not found")
}
