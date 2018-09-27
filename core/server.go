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
	filew fileWrap
}

type fileWrap struct {
	List []string
	Name string
	Data template.HTML
}

//StartServer open the port 7069.
func StartServer(list []string) {
	fmt.Println("+ Stating server on: localhost:7069 | [::1]:7069")
	fmt.Println("+ To exit hit Ctrl+c ...")
	filew := fileWrap{List: list}
	x := &ServeMux{filew: filew}
	err := http.ListenAndServe(":7069", x)
	if err != nil {
		fmt.Printf("! Error on starting server\n\t\t%v\n", err)
	}
}

//ServeHTTP hundle results route
func (x *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/":
		x.mutex.RLock()
		defer x.mutex.RUnlock()
		indexpage(w, r, x.filew, x.filew.List[0])
		return
	case strings.Contains(r.URL.Path, ".css"):
		x.mutex.RLock()
		defer x.mutex.RUnlock()
		r.Header.Set("Content-Type", "text/css")
		http.ServeFile(w, r, "template/"+r.URL.Path)
	case r.URL.Query().Get("f") != "":
		x.mutex.RLock()
		defer x.mutex.RUnlock()
		k, _ := strconv.Atoi(r.URL.Query().Get("f"))
		fmt.Printf("* Processing with %s file\n", x.filew.List[k])
		indexpage(w, r, x.filew, x.filew.List[k])
		return
	default:
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}
func indexpage(w http.ResponseWriter, r *http.Request, filew fileWrap, name string) {
	htmlTemplate := template.New("index.html")
	htmlTemplate, err := htmlTemplate.ParseFiles("template/index.html")
	if err != nil {
		fmt.Printf("! Error html parser\n\t\t%v\n", err)
	}
	filew.Name = name
	filew.Data = template.HTML(MarkdowntoHTML(contentFile(filew.Name)))
	htmlTemplate.Execute(w, filew)
}
