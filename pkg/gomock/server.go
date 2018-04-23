package gomock

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type Server interface {
	Launch() error
}

type server struct {
	Config Config
}

func NewServer(config Config) Server {
	return &server{
		Config: config,
	}
}

func (s *server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	for _, val := range s.Config.Endpoints {
		if val.Path == req.URL.Path {
			if val.Method == req.Method {

				dir, err := os.Getwd()
				if err != nil {
					resp.WriteHeader(http.StatusInternalServerError)
					log.Println(err)
					return
				}

				path := path.Join(dir, strings.Replace(val.ResponseFile, "..", "", -1))

				d, err := ioutil.ReadFile(path)
				if err != nil {
					resp.WriteHeader(http.StatusNoContent)
					log.Println(err)
					return
				}

				resp.Header().Set("Content-Type", "application/json")
				resp.Write(d)
			}
		}
	}
}

func (s *server) Launch() error {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "8080"
	}

	http.Handle("/", s)

	fmt.Println("Starting app on " + port)
	return http.ListenAndServe(":"+port, nil)
}
