package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func IsFile(path string) (bool, error) {
	file, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return !file.Mode().IsDir(), nil
}

func SaveFileHTML(path string, output string, token string) {
	html := MarkdowntoHTML(contentFile(path), token)
	err := ioutil.WriteFile(output, []byte(html), 0755)
	if err != nil {
		fmt.Printf("[ERR] Unable to write file: %v\n", err)
	}
	return
}

//Mdfileslist return a list of markdown file .md or .markdown
func getFileList(dirpath string) []string {
	var fileList []string
	isfile, err := IsFile(dirpath)
	if err != nil {
		fmt.Printf("[ERR] On detecting if it's a file or a directory \n%v\n", err)
		return fileList
	}
	if isfile {
		return []string{dirpath}
	}
	files, err := filepath.Glob(filepath.Join(dirpath, "*.*"))
	if err != nil {
		fmt.Printf("[Err] On listing files on this directory \n%v\n", err)
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
		fmt.Printf("[Err]File does not exist. Path: %s\n", f)
		os.Exit(2)
	}
	file, err := os.OpenFile(f, os.O_RDONLY, 0555)
	defer file.Close()
	if err != nil {
		fmt.Printf("[Err] on reading file check permission. %v\n", err)
		os.Exit(1)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("[Err] While reading the file. %v\n", err)
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
		fmt.Printf("[Err] While executing %s %v\n%v\n", pathExec, args, err)
		return ""
	}
	return strings.Replace(string(cmd), "\n", "", -1)
}

func getGitCommit(dir, fileName string) string {
	return strings.Replace(execute("git", []string{"--git-dir", dir, "log", "--format='%s'", "-n 1", "--", fileName}), "'", "", -1)
}

func getGitDate(dir, fileName string) string {
	return strings.Replace(execute("git", []string{"--git-dir", dir, "log", "--format='%cr'", "-n 1", "--", fileName}), "'", "", -1)
}
