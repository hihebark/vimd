package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hihebark/pickle/core"
)

var (
	f *string
)

func init() {
	f = flag.String("f", "", "Markdown file")
}

func main() {
	fmt.Printf("I'am pickle!\n")
	flag.Parse()
	if *f == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if _, err := os.Stat(*f); os.IsExist(err) {
		fmt.Printf("! File does not exist. Path: %s\n", *f)
		os.Exit(2)
	}
	file, err := os.OpenFile(*f, os.O_RDONLY, 0555)
	defer file.Close()
	if err != nil {
		fmt.Printf("! Error on reading file check permission.\n%v\n", err)
		os.Exit(1)
	}
	//curl https://api.github.com/markdown/raw -X "POST" -H "Content-Type: text/plain" -d "Hello world github/linguist#1 **cool**, and #1!"
	core.DoRequest("https://api.github.com/markdown/raw", "## Hello World")
}
