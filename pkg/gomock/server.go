package gomock

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type server struct {
	Config Config
}

func NewServer(config Config) error {
	s := &server{
		Config: config,
	}

	http.Handle("/", s)

	port := config.Port

	fmt.Println("Starting app on " + port)

	return http.ListenAndServe(":"+port, nil)
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	machedEndpointIdx := s.Config.Endpoints.GetMachingEndpointIndex(req.Method, req.URL.Path)
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
