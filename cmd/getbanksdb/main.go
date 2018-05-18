package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	svnCmd  = `svn export https://github.com/ramoona/banks-db/trunk/banks --force`
	srcGlob = "./banks/*/*.json"
	dstFile = `src/github.com/liuzl/bankcard/banks.go`
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

func svnExport() {
	cmd := exec.Command("/bin/bash", "-c", svnCmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(stderr)
	if err != nil {
		log.Fatal(err, string(data))
	}
	outputBuf := bufio.NewReader(stdout)

	for {
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		log.Println(string(output))
	}

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Fetching " + svnCmd + " from Github")
	//svnExport()
	files, err := filepath.Glob(srcGlob)
	if err != nil {
		log.Fatal(err)
	}
	out1 := bytes.Buffer{}
	out2 := bytes.Buffer{}
	out1.WriteString("package bankcard\n\n")
	out2.WriteString("package bankcard\n\n")
	out1.WriteString("func init() {\n")
	out2.WriteString("func init() {\n")
	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		var bank Bank
		if err = json.Unmarshal(b, &bank); err != nil {
			log.Fatal(err)
		}
		key := strconv.Quote(bank.Key())
		out1.WriteString(fmt.Sprintf("\tBanks[%s] = %s\n",
			key, bank.Code(1)))
		for _, prefix := range bank.Prefixes {
			out2.WriteString(fmt.Sprintf("\tIndex[%d] = %s\n", prefix, key))
		}
	}
	out1.WriteString("}")
	out2.WriteString("}")
	fmt.Println(out1.String())
	fmt.Println(out2.String())
}
