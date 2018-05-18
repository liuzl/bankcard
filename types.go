package bankcard

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type CardType struct {
	Type     string   `json:"type"`
	Prefixes []string `json:"prefixes"`
}

type Bank struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	LocalTitle string `json:"localTitle"`
	EngTitle   string `json:"engTitle"`
	Url        string `json:"url"`
	Color      string `json:"color"`
	Prefixes   []int  `json:"prefixes"`
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
