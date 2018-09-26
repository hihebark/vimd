package core

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	//"strings"
	"sync"
)

type ServeMux struct {
	mutex sync.RWMutex
	md    Markdown
}

//Markdown
type Markdown struct {
	Name string
	Data template.HTML
}

const GITHUBAPI string = "https://api.github.com/markdown/raw"

//MarkdowntoHTML convert given markdown data to html.
func MarkdowntoHTML(data string) string {
	req, err := http.NewRequest("POST", GITHUBAPI, bytes.NewBufferString(data))
	if err != nil {
		fmt.Printf("! Error on request\n\t\t%v\n", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("! Error on response\n\t\t%v\n", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

//StartServer open the port 7069.
func StartServer(md Markdown) {
	fmt.Println("+ Stating server on: localhost:7069 | [::1]:7069")
	fmt.Println("- To exit hit Ctrl+c ...")
	mux := &ServeMux{md: md}
	err := http.ListenAndServe(":7069", mux)
	if err != nil {
		fmt.Printf("! Error on starting server\n\t\t%v\n", err)
	}
}

//ServeHTTP hundle results route
func (mutex *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/":
		mutex.mutex.RLock()
		defer mutex.mutex.RUnlock()
		showResults(w, r, mutex.md)
		return
	case r.URL.Path == "/assets/css/github.css":
		mutex.mutex.RLock()
		defer mutex.mutex.RUnlock()
		r.Header.Set("Content-Type", "text/css")
		http.ServeFile(w, r, "template/assets/css/github.css")
	case r.URL.Path == "/assets/css/syntax.css":
		mutex.mutex.RLock()
		defer mutex.mutex.RUnlock()
		r.Header.Set("Content-Type", "text/css")
		http.ServeFile(w, r, "template/assets/css/syntax.css")
	/*
		case strings.HasSuffix(r.URL.Path, ".css"):
			css := []string{
				"pickle.css",
				"syntax.css",
				"github.css",
			}
			for _, v := range css {
				mutex.mutex.RLock()
				defer mutex.mutex.RUnlock()
				r.Header.Set("Content-Type", "text/css")
				http.ServeFile(w, r, "template/assets/css/"+v)
			}
			return
	*/
	default:
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

//ShowResultsFile see if the directory is in the data/results and show all json files in it
func showResults(w http.ResponseWriter, r *http.Request, data Markdown) {
	htmlTemplate := template.New("index.html")
	htmlTemplate, err := htmlTemplate.ParseFiles("template/index.html")
	if err != nil {
		fmt.Printf("! Error html parser\n\t\t%v\n", err)
	}
	htmlTemplate.Execute(w, data)
}
