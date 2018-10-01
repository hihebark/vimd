package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//Mdfileslist return a list of markdown file .md or .markdown
func Mdfileslist() []string {
	var listfiles []string
	files, err := filepath.Glob("*.*")
	if err != nil {
		fmt.Printf("! Error on listing files on this directory %v\n", err)
	}
	for _, v := range files {
		if filepath.Ext(v) == ".md" || filepath.Ext(v) == ".markdown" {
			listfiles = append(listfiles, v)
		}
	}
	return listfiles
}

func contentFile(f string) string {
	if _, err := os.Stat(f); os.IsExist(err) {
		fmt.Printf("! File does not exist. Path: %s\n", f)
		os.Exit(2)
	}
	file, err := os.OpenFile(f, os.O_RDONLY, 0555)
	defer file.Close()
	if err != nil {
		fmt.Printf("! Error on reading file check permission.\n\t\t%v\n", err)
		os.Exit(1)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("! Error reading file.\n\t\t%v\n", err)
	}
	return string(content)
}

func execute(pathExec string, args []string) string {
	path, err := exec.LookPath(pathExec)
	if err != nil {
		return ""
	}
	cmd, err := exec.Command(path, args...).CombinedOutput()
	if err != nil {
		fmt.Printf("! Error while executing %v\n", err)
		return ""
	}
	return strings.Replace(string(cmd), "\n", "", -1)
}
