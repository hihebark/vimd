package core

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hihebark/pickle/log"
)

func IsFile(path string) (bool, error) {
	file, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return !file.Mode().IsDir(), nil
}

func SaveFileHTML() {
	return
}

//Mdfileslist return a list of markdown file .md or .markdown
func Mdfileslist() []string {
	var listfiles []string
	files, err := filepath.Glob("*.*")
	if err != nil {
		log.Err("Error on listing files on this directory %v", err)
	}
	for _, v := range files {
		if filepath.Ext(v) == ".md" || filepath.Ext(v) == ".markdown" {
			listfiles = append(listfiles, v)
		}
	}
	return listfiles
}

func getFileList(dirpath string) []string {
	var fileList []string
	files, err := filepath.Glob("*.*")
	if err != nil {
		log.Err("Error on listing files on this directory %v", err)
	}
	for _, file := range files {
		if filepath.Ext(file) == ".md" || filepath.Ext(file) == ".markdown" {
			fileList = append(fileList, file)
		}
	}
	return fileList
}

func contentFile(f string) string {
	if _, err := os.Stat(f); os.IsExist(err) {
		log.Err("File does not exist. Path: %s", f)
		os.Exit(2)
	}
	file, err := os.OpenFile(f, os.O_RDONLY, 0555)
	defer file.Close()
	if err != nil {
		log.Err("Error on reading file check permission. %v", err)
		os.Exit(1)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Err("Error reading file. %v", err)
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
		log.Err("Error while executing %v", err)
		return ""
	}
	return strings.Replace(string(cmd), "\n", "", -1)
}
