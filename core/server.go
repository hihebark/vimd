package core

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

//ServeMux
type ServeMux struct {
	mutex sync.RWMutex
	filew filesWrap
}

//FileWrap
type fileWrap struct {
	Name string
	Data template.HTML
}
type filesWrap struct {
	filewrap []fileWrap
}

func newFilesWrap(list []string) *filesWrap {
	var fileswrap []fileWrap
	for _, v := range list {
		fileswrap = append(fileswrap, fileWrap{Name: v})
	}
	return &filesWrap{fileswrap}
}

//StartServer open the port 7069.
func StartServer(list []string) {
	fmt.Println("+ Stating server on: localhost:7069 | [::1]:7069")
	fmt.Println("+ To exit hit Ctrl+c ...")
	x := &ServeMux{filew: *newFilesWrap(list)}
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
		indexpage(w, r, x.filew.filewrap[0])
		return
	case r.URL.Path == "/assets/css/github.css":
		x.mutex.RLock()
		defer x.mutex.RUnlock()
		r.Header.Set("Content-Type", "text/css")
		http.ServeFile(w, r, "template/assets/css/github.css")
	case r.URL.Path == "/assets/css/syntax.css":
		x.mutex.RLock()
		defer x.mutex.RUnlock()
		r.Header.Set("Content-Type", "text/css")
		http.ServeFile(w, r, "template/assets/css/syntax.css")
	case r.URL.Query().Get("p") != "":
		return
	default:
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

//indexpage
func indexpage(w http.ResponseWriter, r *http.Request, filew fileWrap) {
	htmlTemplate := template.New("index.html")
	htmlTemplate, err := htmlTemplate.ParseFiles("template/index.html")
	if err != nil {
		fmt.Printf("! Error html parser\n\t\t%v\n", err)
	}
	htmlTemplate.Execute(w, filew)
}
