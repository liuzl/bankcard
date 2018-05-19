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

func BinCheck(bin string) bool {
	var t = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	odd := len(bin) & 1
	var sum int
	for i, c := range bin {
		if c < '0' || c > '9' {
			return false
		}
		if i&1 == odd {
			sum += t[c-'0']
		} else {
			sum += int(c - '0')
		}
	}
	return sum%10 == 0
}

func FindBank(card string) (*BankPro, error) {
	if len(card) < MaxLen || !goutil.StringIs(card, unicode.IsDigit) {
		return nil, fmt.Errorf("format error")
	}
	if !BinCheck(card) {
		return nil, fmt.Errorf("invalid card number")
	}

	for k, m := range BanksProMap {
		for i := MaxLenProMap[k]; i >= 1; i-- {
			n, err := strconv.Atoi(card[:i])
			if err != nil {
				return nil, err
			}
			if bank := IndexProMap[k][n]; bank != "" {
				if len(card) == m[bank].Length {
					return m[bank], nil
				}
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
