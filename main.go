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
	fmt.Printf("\n  [I'am pickle]\n\n")
	flag.Parse()
	var list []string
	if *f == "" {
		list = core.Mdfileslist()
		if len(list) == 0 {
			fmt.Printf("! Error no markdown file in this directory.")
			os.Exit(0)
		}
		fmt.Printf("Files %v\n", list)
	} else {
		list = append(list, *f)
	}
	core.StartServer(list)
	//htmldata := core.MarkdowntoHTML(string(data))
	//core.StartServer(core.Markdown{file.Name(), template.HTML(htmldata)})
}
