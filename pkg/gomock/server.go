package gomock

import (
	"io/ioutil"
	"net/http"

	"github.com/kpango/glg"
)

type Server interface {
	Serve() error
}

type server struct {
	Logger *glg.Glg
	Config Config
}

func NewServer(config Config) Server {
	return &server{
		Config: config,
		Logger: glg.New(),
	}
}

func (s *server) Serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Info(req.Method + " " + req.URL.String())
		s.ServeHTTP(w, req)
	})

	port := s.Config.Port

	s.Logger.Info("Starting app on " + port)

	return http.ListenAndServe(":"+port, nil)
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	machedEndpointIdx := s.Config.Endpoints.GetMachingEndpointIndex(req.Method, req.URL.String())
	if machedEndpointIdx < 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response := s.Config.Endpoints[machedEndpointIdx].Response

	d, err := ioutil.ReadFile(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	for key, value := range response.Headers {
		w.Header().Set(key, value)
	}

	w.WriteHeader(response.Code)
	w.Write(d)
}
