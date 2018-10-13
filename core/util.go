package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//GITHUBAPIURL link to api to access github/markdown
const GITHUBAPIURL string = "https://api.github.com/markdown/raw"

//MarkdowntoHTML convert given markdown data to html.
func MarkdowntoHTML(data, token string) string {
	req, err := http.NewRequest("POST", GITHUBAPIURL, bytes.NewBufferString(data))
	if err != nil {
		fmt.Printf("! Error on request\n\t\t%v\n", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	if len(token) != 0 {
		req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("! Error on response\n\t\t%v\n", err)
	}
	if resp.StatusCode != 200 {
		return "Error with rate limit. report to <a href='https://github.com/hihebark/pickle/issues'>issue #3</a> or just use -token argument!"
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
