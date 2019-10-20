package main

import (
	"flag"
	"fmt"

	"github.com/hihebark/vimd/core"
)

const VERSION = "0.2.0"

var (
	path   *string
	token  *string
	save   *bool
	output *string
	watch  *bool
	port   *string
)

func init() {
	path = flag.String("p", ".", "Path that contain the markdown file(s)")
	port = flag.String("port", "7069", "The serve port")
	token = flag.String("token", "", "Github personal accesss token")
	save = flag.Bool("s", false, "Save the output into an html file(s)")
	output = flag.String("o", "", "The name and path of the output rendred HTML")
	watch = flag.Bool("watch", false, "Reload when change is detected on the path (soon)")
}

func main() {
	fmt.Printf("\n[ Vimd ] - %s\n\n", VERSION)
	flag.Parse()
	if *save && *path != "" && *output != "" {
		isFile, err := core.IsFile(*path)
		if err != nil {
			fmt.Printf("[Err] While checking the state of the file%v\n", err)
		}
		if isFile {
			fmt.Printf("[+] Saving file into: %s\n", *output)
			core.SaveFileHTML(*path, *output, *token)
		} else {
			fmt.Printf("[ERR] Can't save a directory into an html file\n")
		}
	} else {
		server := core.NewServ(*port, *path, *token, *watch)
		err := server.Start()
		if err != nil {
			fmt.Printf("[ERR] While executing server.Start:\n%v\n", err)
		}
	}
}
