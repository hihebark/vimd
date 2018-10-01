package core

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// ServeMux just a mutex
type ServeMux struct {
	mutex sync.RWMutex
	wraps Wraps
	token string
}

type Wrap struct {
	Commit string
	Date   string
	Name   string
}
type Wraps struct {
	Wraps   []Wrap
	Content template.HTML
	Name    string
}

//StartServer open the port 7069.
func StartServer(list []string, token string) {
	fmt.Println("+ Stating server on: localhost:7069 | [::1]:7069")
	fmt.Println("+ To exit hit Ctrl+c ...")
	var ws []Wrap
	for _, v := range list {
		commit := strings.Replace(execute("git", []string{"log", "--format='%s'", "-n 1", v}), "'", "", -1)
		date := strings.Replace(execute("git", []string{"log", "--format='%cr'", "-n 1", v}), "'", "", -1)
		ws = append(ws, Wrap{commit, date, v})
	}
	x := &ServeMux{
		wraps: Wraps{
			Wraps: ws,
		},
		token: token,
	}
	err := http.ListenAndServe(":7069", x)
	if err != nil {
		fmt.Printf("! Error on starting server\n\t\t%v\n", err)
	}
}

//ServeHTTP hundle results route
func (x *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/index":
		x.mutex.RLock()
		defer x.mutex.RUnlock()
		indexpage(w, r, x.wraps, 0, x.token)
		return
	case r.URL.Query().Get("f") != "":
		x.mutex.RLock()
		defer x.mutex.RUnlock()
		k, _ := strconv.Atoi(r.URL.Query().Get("f"))
		if k > len(x.wraps.Wraps) {
			fmt.Printf("! Processing with unknown page %d - redirecting to home\n", k)
			http.Redirect(w, r, "/index", http.StatusFound)
		}
		fmt.Printf("* Processing with %s file\n", x.wraps.Wraps[k])
		indexpage(w, r, x.wraps, k, x.token)
		return
	default:
		http.Redirect(w, r, "/index", http.StatusFound)
		return
	}
}
func indexpage(w http.ResponseWriter, r *http.Request, wraps Wraps, key int, token string) {
	htmlTemplate, err := template.New("index.html").Parse(TEMPLATE)
	if err != nil {
		fmt.Printf("! Error html parser\n\t\t%v\n", err)
	}
	wraps.Name = wraps.Wraps[key].Name
	wraps.Content = template.HTML(MarkdowntoHTML(contentFile(wraps.Name), token))
	htmlTemplate.Execute(w, wraps)
}
