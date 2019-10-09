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
	output *string
	watch  *bool
)

func init() {
	path = flag.String("path", "", "Path that contain the markdown file(s)")
	token = flag.String("token", "", "Token ...")
	save = flag.Bool("save", false, "Save the output into an html file(s)")
	output = flag.String("output", "", "")
	watch = flag.Bool("watch", false, "Reload when change is detected on the path")
}

func main() {
	fmt.Printf("-=======-\n")
	fmt.Printf("| Vismd | - 0.2.0\n")
	fmt.Printf("-=======-\n")
	flag.Parse()
	if *save && *path != "" && *output != "" {
		isFile, err := core.IsFile(*path)
		if err != nil {
			fmt.Printf("Err %v\n", err)
		}
		if isFile {
			fmt.Printf("Saving file into current directory\n")
			core.SaveFileHTML(*path, *output, *token)
		} else {
			fmt.Printf("[ERR] Error Cant save a directory into an html file\n")
		}
	} else {
		server := core.NewServ(*path, *token, *watch)
		err := server.Start()
		if err != nil {
			fmt.Printf("[ERR] Error while executing server.Start:\n%v\n", err)
		}
	}
}
