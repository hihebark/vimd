package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hihebark/pickle/core"
)

var (
	file   *string
	token  *string
	static *string
)

func init() {
	file = flag.String("file", "", "Markdown file")
	token = flag.String("token", "", "Github api token.")
	static = flag.String("static", ".", "Static file image, video, ...")
}

func main() {
	fmt.Printf("\n  [I'am pickle]\n\n")
	flag.Parse()
	var list []string
	if *file == "" {
		list = core.Mdfileslist()
		if len(list) == 0 {
			fmt.Printf("! Error no markdown file in this directory.")
			os.Exit(0)
		}
	} else {
		list = append(list, *file)
	}
	core.StartServer(list, *token, *static)
}
