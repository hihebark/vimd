package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//DoRequest
func DoRequest(url, data string) {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
	if err != nil {
		fmt.Printf("! Error on request %v\n", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("! Error on response %v\n")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%v\n", string(body))
}
