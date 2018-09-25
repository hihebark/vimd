package core

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

type ServeMux struct {
	mutex sync.RWMutex
	md    Markdown
}

//Markdown
type Markdown struct {
	Name string
	Data string
}

//DoRequest
func DoRequest(url, data string) string {
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
	return string(body)
	//	fmt.Printf("%v\n", string(body))
}

//StartListning start listning to the given port
func StartServer(md Markdown) {
	fmt.Printf("Stating server on http://localhost:7069/ | http://[::1]:7069/\n")
	mux := &ServeMux{md: md}
	err := http.ListenAndServe(":7069", mux)
	if err != nil {
		fmt.Printf("! Error on starting server %v\n", err)
	}
}

//ServeHTTP hundle results route
func (mutex *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch {
	case r.URL.Path == "/":
		mutex.mutex.RLock()
		defer mutex.mutex.RUnlock()
		ShowResultsFile(w, r, mutex.md)
		return
	/*case r.URL.Path == "css/":
	mutex.mutex.RLock()
	defer mutex.mutex.RUnlock()
	http.ServeFile(w, r, "template/css/")
	return*/
	default:
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

}

//ShowResultsFile see if the directory is in the data/results and show all json files in it
func ShowResultsFile(w http.ResponseWriter, r *http.Request, data Markdown) {
	htmlTemplate := template.New("index.html")
	htmlTemplate, err := htmlTemplate.ParseFiles("template/index.html")
	if err != nil {
		fmt.Printf("! Error html parser %v\n", err)
	}
	htmlTemplate.Execute(w, data)
}
