package core

import (
	"fmt"
	"net/http"
	"path"
	"sync"
)

type metadata struct {
	commit string
	date   string
}
type file struct {
	path     string
	name     string
	metadata metadata
}

type Server struct {
	mutex sync.RWMutex
	token string
	files []file
}

func NewServ(dirpath, token, assets string, reload bool) *Server {
	fmt.Printf("Initialising server\n")
	fileList := getFileList(dirpath)

	server := &Server{
		token: token,
		files: mdFetcher(fileList),
	}
	return server
}

func (s *Server) Start() error {
	fmt.Printf("Starting Server....\n")
	err := http.ListenAndServe(":7069", s)
	if err != nil {
		return err
	}
	return nil
}

func mdFetcher(paths []string) []file {
	var files []file
	for _, v := range paths {
		files = append(files, file{
			path: v,
			name: path.Base(v),
			metadata: metadata{
				commit: "",
				date:   "",
			},
		})
	}
	return files
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

/*
import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/hihebark/pickle/log"
)

// ServeMux just a mutex
type ServeMux struct {
	mutex  sync.RWMutex
	wraps  Wraps
	token  string
	static string
}

// Wrap .
type Wrap struct {
	Commit string
	Date   string
	Name   string
}

// Wraps .
type Wraps struct {
	Wraps   []Wrap
	Content template.HTML
	Name    string
}

// StartServer open the port 7069.
func StartServer(list []string, token, static string) {
	log.Inf("Stating server on: localhost:7069 | [::1]:7069")
	log.Inf("To exit hit Ctrl+c ...")
	var ws []Wrap
	for _, name := range list {
		commit := strings.Replace(execute("git", []string{"log", "--format='%s'", "-n 1", name}), "'", "", -1)
		date := strings.Replace(execute("git", []string{"log", "--format='%cr'", "-n 1", name}), "'", "", -1)
		ws = append(ws, Wrap{commit, date, name})
	}
	x := &ServeMux{
		wraps: Wraps{
			Wraps: ws,
		},
		token:  token,
		static: static,
	}
	err := http.ListenAndServe(":7069", x)
	if err != nil {
		log.Err("Error on starting server %v", err)
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
			log.Inf("Processing with unknown page %d - redirecting to home", k)
			http.Redirect(w, r, "/index", http.StatusFound)
		}
		log.Inf("Processing with %s file", x.wraps.Wraps[k].Name)
		indexpage(w, r, x.wraps, k, x.token)
		return
	default:
		path := r.URL.Path[1:]
		data, _ := ioutil.ReadFile(string(path))
		switch {
		case strings.HasSuffix(path, "jpg") || strings.HasSuffix(path, "jpeg") || strings.HasSuffix(path, "png"):
			r.Header.Add("Content-type", "image/*")
			w.Write(data)
		default:
			http.Redirect(w, r, "/index", http.StatusFound)
		}
		return
	}
}
func indexpage(w http.ResponseWriter, r *http.Request, wraps Wraps, key int, token string) {
	htmlTemplate, err := template.New("index.html").Parse(TEMPLATE)
	if err != nil {
		log.Err("Error html parser %v", err)
	}
	wraps.Name = wraps.Wraps[key].Name
	wraps.Content = template.HTML(MarkdowntoHTML(contentFile(wraps.Name), token))
	htmlTemplate.Execute(w, wraps)
}
*/
