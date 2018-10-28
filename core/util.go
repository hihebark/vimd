package core

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hihebark/pickle/log"
)

//GITHUBAPIURL link to api to access github/markdown
const GITHUBAPIURL string = "https://api.github.com/markdown/raw"

//MarkdowntoHTML convert given markdown data to html.
func MarkdowntoHTML(data, token string) string {
	req, err := http.NewRequest("POST", GITHUBAPIURL, bytes.NewBufferString(data))
	if err != nil {
		log.Err("Error on request %v", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	if len(token) != 0 {
		req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Err("Error on response %v", err)
	}
	if resp.StatusCode != 200 {
		log.Dbg("Rate limit\nThe Response: %v\n%v", errors.New("Rate limit"), resp.Header)
		return "Error with rate limit. report to <a href='https://github.com/hihebark/pickle/issues'>issue #3</a> or just use -token argument!"
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
