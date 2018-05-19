package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/liuzl/bankcard"
	"github.com/liuzl/dl"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	fileUrl = "https://raw.githubusercontent.com/MardaWang0518/BankCardDemo/master/app/src/main/assets/bankinfo.txt"
	numReg  = regexp.MustCompile(`(\d+)`)

	bankFile  = `src/github.com/liuzl/bankcard/cn_banks.go`
	indexFile = `src/github.com/liuzl/bankcard/cn_index.go`
)

type Bank struct {
	BankName string     `json:"bankName"`
	BankCode string     `json:"bankCode"`
	Patterns []*Pattern `json:"patterns"`
}

type Pattern struct {
	Reg      string `json:"reg"`
	CardType string `json:"cardType"`
}

func main() {
	gopath, found := os.LookupEnv("GOPATH")
	if !found {
		log.Fatal("Missing $GOPATH environment variable")
	}

	resp := dl.DownloadUrl(fileUrl)
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
	text := strings.Trim(strings.TrimSpace(strings.Replace(resp.Text, "@", "", -1)), ";")
	var items []*Bank
	err := json.Unmarshal([]byte(text), &items)
	if err != nil {
		log.Fatal(err)
	}

	out1 := bytes.Buffer{}
	out2 := bytes.Buffer{}
	out1.WriteString("package bankcard\n\n")
	out2.WriteString("package bankcard\n\n")
	out1.WriteString("func init() {\n\tbanks := make(map[string]*BankPro)\n")
	out2.WriteString("func init() {\n\tindex := make(map[int]string)\n")

	maxLen := 0
	for _, item := range items {
		for _, p := range item.Patterns {
			b := &bankcard.BankPro{
				Bank: bankcard.Bank{
					Name:       strings.ToLower(item.BankCode),
					Country:    "cn",
					LocalTitle: item.BankName,
				},
				CardType: strings.ToLower(p.CardType),
			}
			matches := numReg.FindAllString(p.Reg, -1)
			n := len(matches)
			if n < 2 {
				log.Fatal(p.Reg)
			}
			if b.Length, err = strconv.Atoi(matches[n-1]); err != nil {
				log.Fatal(err)
			}
			prefixLen := len(matches[0])
			b.Length += prefixLen
			if prefixLen > maxLen {
				maxLen = prefixLen
			}
			key := strconv.Quote(b.Key())
			for i := 0; i < n-1; i++ {
				prefix, err := strconv.Atoi(matches[i])
				if err != nil {
					log.Fatal(err)
				}
				b.Prefixes = append(b.Prefixes, prefix)
				out2.WriteString(fmt.Sprintf("\tindex[%d] = %s\n", prefix, key))
			}
			out1.WriteString(fmt.Sprintf("\tbanks[%s] = %s\n", key, b.Code(1)))
		}
	}
	out1.WriteString("\tBanksProMap[\"cn\"] = banks\n")
	out1.WriteString(fmt.Sprintf("\tMaxLenProMap[\"cn\"] = %d\n", maxLen))
	out1.WriteString("}")
	out2.WriteString("\tIndexProMap[\"cn\"] = index\n")
	out2.WriteString("}")
	bankFile = filepath.Join(gopath, bankFile)
	indexFile = filepath.Join(gopath, indexFile)
	if err = ioutil.WriteFile(bankFile, out1.Bytes(), os.FileMode(0664)); err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile(indexFile, out2.Bytes(), os.FileMode(0664)); err != nil {
		log.Fatal(err)
	}
}
