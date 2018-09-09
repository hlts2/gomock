package gomock

import (
	"io/ioutil"
	"net/http"

	"github.com/kpango/glg"
)

// Server is core API mock server interface
type Server interface {
	Serve() error
}

type server struct {
	Logger *glg.Glg
	Config Config
}

// NewServer returns Server(*server) object
func NewServer(config Config) Server {
	return &server{
		Config: config,
		Logger: glg.New(),
	}
}

func (s *server) Serve() error {
	port := s.Config.Port

	s.Logger.Info("Starting app on " + port)

	return http.ListenAndServe(":"+port, s)
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.Logger.Info(req.Method + " " + req.URL.String())

	machedEndpointIdx := s.Config.Endpoints.GetMachingEndpointIndex(req.Method, req.URL.String())
	if machedEndpointIdx < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response := s.Config.Endpoints[machedEndpointIdx].Response

	d, err := ioutil.ReadFile(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for key, value := range response.Headers {
		if _, ok := w.Header()[key]; ok {
			w.Header().Add(key, value)
		} else {
			w.Header().Set(key, value)
		}
	}

	w.WriteHeader(response.Code)
	w.Write(d)
}
