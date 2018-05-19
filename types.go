package bankcard

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Bank struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	LocalTitle string `json:"localTitle"`
	EngTitle   string `json:"engTitle"`
	Url        string `json:"url"`
	Color      string `json:"color"`
	Prefixes   []int  `json:"prefixes"`
}

type BankPro struct {
	Bank
	CardType string `json:"card_type"`
	Length   int    `json:"length"`
}

func (b *BankPro) Key() string {
	return fmt.Sprintf("%s_%s_%s_%d", b.Country, b.Name, b.CardType, b.Length)
}

func (b *Bank) Key() string {
	return b.Country + "_" + b.Name
}

func (b *Bank) Code(n int) string {
	ind := ""
	for i := 0; i < n; i++ {
		ind += "\t"
	}
	output := bytes.Buffer{}
	output.WriteString("&Bank{\n")
	output.WriteString(fmt.Sprintf("%s\t%s,%s,%s,%s,%s,%s,\n", ind,
		strconv.Quote(b.Name),
		strconv.Quote(b.Country),
		strconv.Quote(b.LocalTitle),
		strconv.Quote(b.EngTitle),
		strconv.Quote(b.Url),
		strconv.Quote(b.Color)))
	var strs []string
	for _, i := range b.Prefixes {
		strs = append(strs, strconv.Itoa(i))
	}
	output.WriteString(fmt.Sprintf("%s\t[]int{%s},\n", ind, strings.Join(strs, ",")))
	output.WriteString(ind + "}")
	return output.String()
}

func (b *BankPro) Code(n int) string {
	ind := ""
	for i := 0; i < n; i++ {
		ind += "\t"
	}
	output := bytes.Buffer{}
	output.WriteString(fmt.Sprintf("&BankPro{\n%s\tBank{\n", ind))
	output.WriteString(fmt.Sprintf("\t%s\t%s,%s,%s,%s,%s,%s,\n", ind,
		strconv.Quote(b.Name),
		strconv.Quote(b.Country),
		strconv.Quote(b.LocalTitle),
		strconv.Quote(b.EngTitle),
		strconv.Quote(b.Url),
		strconv.Quote(b.Color)))
	var strs []string
	for _, i := range b.Prefixes {
		strs = append(strs, strconv.Itoa(i))
	}
	output.WriteString(fmt.Sprintf("\t%s\t[]int{%s},\n", ind, strings.Join(strs, ",")))
	output.WriteString(ind + "\t},\n")
	output.WriteString(fmt.Sprintf("%s%s, %d,\n%s}",
		ind, strconv.Quote(b.CardType), b.Length, ind))
	return output.String()
}
