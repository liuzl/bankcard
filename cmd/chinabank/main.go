package main

import (
	"encoding/json"
	"fmt"
	"github.com/liuzl/dl"
	"log"
	"strings"
)

var (
	fileUrl = "https://raw.githubusercontent.com/MardaWang0518/BankCardDemo/master/app/src/main/assets/bankinfo.txt"
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
	resp := dl.DownloadUrl(fileUrl)
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
	text := strings.Trim(strings.TrimSpace(strings.Replace(resp.Text, "@", "", -1)), ";")
	var items []Bank
	err := json.Unmarshal([]byte(text), &items)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(items)
}
