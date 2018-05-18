package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/liuzl/bankcard"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var (
	svnCmd    = `svn export https://github.com/ramoona/banks-db/trunk/banks --force`
	srcGlob   = "./banks/*/*.json"
	bankFile  = `src/github.com/liuzl/bankcard/banks.go`
	indexFile = `src/github.com/liuzl/bankcard/index.go`
)

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
	gopath, found := os.LookupEnv("GOPATH")
	if !found {
		log.Fatal("Missing $GOPATH environment variable")
	}
	log.Println("Fetching " + svnCmd + " from Github")
	svnExport()
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
	maxLen := 0
	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		var bank bankcard.Bank
		if err = json.Unmarshal(b, &bank); err != nil {
			log.Fatal(err)
		}
		key := strconv.Quote(bank.Key())
		out1.WriteString(fmt.Sprintf("\tBanks[%s] = %s\n",
			key, bank.Code(1)))
		for _, prefix := range bank.Prefixes {
			out2.WriteString(fmt.Sprintf("\tIndex[%d] = %s\n", prefix, key))
			p := strconv.Itoa(prefix)
			if len(p) > maxLen {
				maxLen = len(p)
			}
		}
	}
	out1.WriteString("}")
	out1.WriteString(fmt.Sprintf("\n\nvar MaxLen = %d", maxLen))
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
