package main

import (
	"flag"
	"fmt"

	"github.com/hihebark/pickle/core"
)

var (
	path   *string
	token  *string
	save   *bool
	reload *bool
	assets *string
)

func init() {
	path = flag.String("path", "", "Path that contain the markdown file(s)")
	token = flag.String("token", "", "Token ...")
	save = flag.Bool("save", false, "Save the output into an html file(s)")
	reload = flag.Bool("reload", false, "Reload when change is detected on the path")
	assets = flag.String("assets", "", "")
}

func main() {
	fmt.Printf("[ gomd ] - 0.2.0\n")
	flag.Parse()
	if *save && *path != "" {
		isFile, err := core.IsFile(*path)
		if err != nil {
			fmt.Printf("Err %v\n", err)
		}
		if isFile {
			fmt.Printf("Saving file into current directory\n")
			core.SaveFileHTML()
		} else {
			fmt.Printf("[ERR] Error Cant save a directory into an html file\n")
		}
	} else {
		server := core.NewServ(*path, *token, *assets, *reload)
		err := server.Start()
		if err != nil {
			fmt.Printf("[ERR] Error while executing server.Start:\n%v\n", err)
		}
	}
}

/*
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
*/
