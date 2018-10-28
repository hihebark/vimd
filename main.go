package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hihebark/pickle/core"
	"github.com/hihebark/pickle/log"
)

var (
	file   *string
	token  *string
	static *string
	save   *bool
)

func init() {
	file = flag.String("file", "", "Markdown file")
	token = flag.String("token", "", "Github api token.")
	static = flag.String("static", ".", "Static file image, video, ...")
	save = flag.Bool("save", false, "Save as HTML file.")
}

func main() {
	fmt.Printf("\n[I'am pickle!]\n\n")
	flag.Parse()
	var list []string
	if *file == "" {
		list = core.Mdfileslist()
		if len(list) == 0 {
			log.Err("Error no markdown file in this directory.")
			//fmt.Errorf("! Error no markdown file in this directory.")
			os.Exit(0)
		}
	} else {
		list = append(list, *file)
	}
	if *save {
		//
	}
	core.StartServer(list, *token, *static)
}
