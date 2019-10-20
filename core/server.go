package core

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"sync"

	socketio "github.com/googollee/go-socket.io"
)

type Server struct {
	mutex    sync.RWMutex
	port     string
	socket   *socketio.Server
	rendring rendring
	token    string
	watcher  bool
}

type rendring struct {
	Current string
	Content template.HTML
	Files   []file
}

type file struct {
	Path     string
	Name     string
	Metadata metadata
}

type metadata struct {
	Commit string
	Date   string
}

func NewServ(port, dirpath, token string, watch bool) *Server {
	fmt.Printf("[*] Initialising server\n")
	fileList := getFileList(dirpath)
	socket, err := socketio.NewServer(nil)
	if err != nil {
		fmt.Printf("[ERR] socketio NewServer %v\n", err)
	}
	dir := "./.git"
	if path.Base(dirpath) != dirpath {
		dir = filepath.Join(dirpath, "/.git")
	}
	server := &Server{
		port:    port,
		watcher: watch,
		token:   token,
		socket:  socket,
		rendring: rendring{
			Current: "Error",
			Content: "",
			Files:   mdFetcher(fileList, dir),
		},
	}
	return server
}

func (s *Server) initSocket() {
	s.socket.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		fmt.Println("connected:", so.ID())
		return nil
	})
	s.socket.OnEvent("/", "hello", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	s.socket.OnDisconnect("/", func(so socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})
	go s.socket.Serve()
	defer s.socket.Close()
}

func (s *Server) Start() error {
	if s.watcher {
		fmt.Printf("[+] Starting socket....\n")
		s.initSocket()
		http.Handle("/socket.io/", s.socket)
	}
	fmt.Printf("[+] Starting server at http://localhost:%s/\n", s.port)
	err := http.ListenAndServe(":"+s.port, s)
	if err != nil {
		return err
	}
	return nil
}

func mdFetcher(paths []string, dir string) []file {
	var files []file
	for _, v := range paths {
		files = append(files, file{
			Path: v,
			Name: path.Base(v),
			Metadata: metadata{
				Commit: getGitCommit(dir, path.Base(v)),
				Date:   getGitDate(dir, path.Base(v)),
			},
		})
	}
	return files
}

func (s *Server) contain(str string) (int, bool) {
	for k, v := range s.rendring.Files {
		if v.Name == str {
			return k, true
		}
	}
	return 0, false
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/":
		s.mutex.RLock()
		defer s.mutex.RUnlock()
		s.render(w, r, 0)
		return
	case r.URL.Path != "/":
		s.mutex.RLock()
		defer s.mutex.RUnlock()
		key, isIn := s.contain(path.Base(r.URL.Path))
		if isIn {
			s.render(w, r, key)
		} else {
			path := r.URL.Path[1:]
			data, _ := ioutil.ReadFile(string(path))
			if strings.HasSuffix(path, "jpg") || strings.HasSuffix(path, "jpeg") || strings.HasSuffix(path, "png") {
				r.Header.Add("Content-type", "image/*")
				w.Write(data)
			} else {
				s.notFoundPage(w, r)
			}
		}
		return
	default:
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

func (s *Server) notFoundPage(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, err := template.New("404.html").Parse(NOTFOUNDPAGE)
	if err != nil {
		fmt.Printf("[ERR] Error html parser \n%v\n", err)
	}
	htmlTemplate.Execute(w, s.rendring)
}

func (s *Server) render(w http.ResponseWriter, r *http.Request, key int) {
	htmlTemplate, err := template.New("index.html").Parse(TEMPLATE)
	if err != nil {
		fmt.Printf("[ERR] Error html parser \n%v\n", err)
	}
	if len(s.rendring.Files) != 0 && key <= len(s.rendring.Files) {
		s.rendring.Current = s.rendring.Files[key].Name
		s.rendring.Content = template.HTML(MarkdowntoHTML(contentFile(s.rendring.Files[key].Path), s.token))
	} else {
		s.rendring.Content = "No <b>markdown</b> file found"
	}
	htmlTemplate.Execute(w, s.rendring)
}
