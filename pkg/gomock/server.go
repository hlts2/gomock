package gomock

import (
	"net/http"

	"github.com/kpango/glg"
	"github.com/pkg/errors"
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
	return &server{
		port:   config.Port,
		router: NewRouter(config.Endpoints),
	}
}

func (s *server) Serve() error {
	glg.Info("Starting app on " + s.port)
	err := http.ListenAndServe(":"+s.port, s.router)
	if err != nil {
		return errors.Wrap(err, "faild to listen")
	}
	return nil
}

func (s *server) ServeTLS(crtPath, keyPath string) error {
	glg.Info("Starting app on " + s.port)
	err := http.ListenAndServeTLS(":"+s.port, crtPath, keyPath, s.router)
	if err != nil {
		return errors.Wrap(err, "faild to listen")
	}
	return nil
}
