package gomock

import (
	"fmt"
	"net/http"

	"github.com/kpango/glg"
)

// Server is core API mock server interface
type Server interface {

	// Server starts server
	Serve() error

	// ServeTLS starts server with TLS mode
	ServeTLS(crtPath, keyPath string) error
}

type server struct {
	port   string
	router Router
}

// NewServer returns Server(*server) object
func NewServer(config *Config) Server {
	fmt.Println("hoge")
	return &server{
		port:   config.Port,
		router: NewRouter(config.Endpoints),
	}
}

func (s *server) Serve() error {
	glg.Info("Starting app on " + s.port)
	return http.ListenAndServe(":"+s.port, s.router)
}

func (s *server) ServeTLS(crtPath, keyPath string) error {
	glg.Info("Starting app on " + s.port)
	return http.ListenAndServeTLS(":"+s.port, crtPath, keyPath, s.router)
}
