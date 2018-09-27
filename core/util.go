package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const GITHUBAPI string = "https://api.github.com/markdown/raw"

//MarkdowntoHTML convert given markdown data to html.
func MarkdowntoHTML(data string) string {
	req, err := http.NewRequest("POST", GITHUBAPI, bytes.NewBufferString(data))
	if err != nil {
		fmt.Printf("! Error on request\n\t\t%v\n", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	if os.Getenv("PICKLETOKEN") != "" {
		req.Header.Set("Authorizations", fmt.Sprintf("token %s", os.Getenv("PICKLETOKEN")))
	}
	client := &http.Client{}
	fmt.Printf("- Getting html from github ...\n")
	resp, err := client.Do(req)
	fmt.Printf("%v\n", resp.StatusCode)
	if err != nil {
		fmt.Printf("! Error on response\n\t\t%v\n", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
