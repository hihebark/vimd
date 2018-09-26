package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
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
	fmt.Printf("  [I'am pickle!]\n")
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
		fmt.Printf("! Error on reading file check permission.\n\t\t%v\n", err)
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("! Error reading file.\n\t\t%v\n", err)
	}
	htmldata := core.MarkdowntoHTML(string(data))
	core.StartServer(core.Markdown{file.Name(), template.HTML(htmldata)})
}
