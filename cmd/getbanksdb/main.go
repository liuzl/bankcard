package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
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
	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		var bank Bank
		if err = json.Unmarshal(b, &bank); err != nil {
			log.Fatal(err)
		}
		fmt.Println(bank.Key())
	}
}
