package gomock

import (
	"io/ioutil"
	"net/http"

	"github.com/hlts2/lilty"
	"github.com/kpango/glg"
	"github.com/pkg/errors"
)

// ErrUnsupportedHTTPMethod is error of unsupported HTTP method
var ErrUnsupportedHTTPMethod = errors.New("unsupported HTTP method")

// Server is core API mock server interface
type Server interface {
	lilty.Lilty
}

type server struct {
	lilty.Lilty
	endpoints Endpoints
}

// NewServer returns Server(*server) object
func NewServer(endpoints Endpoints) (Server, error) {
	s := &server{
		Lilty:     lilty.Default(),
		endpoints: endpoints,
	}

	err := s.loadRoute()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *server) loadRoute() error {
	for _, e := range s.endpoints {
		switch e.Request.Method {
		case string(lilty.GET):
			s.Lilty.Get(e.Request.Path, s.ServeHTTP)
		case string(lilty.HEAD):
			s.Lilty.Head(e.Request.Path, s.ServeHTTP)
		case string(lilty.POST):
			s.Lilty.Post(e.Request.Path, s.ServeHTTP)
		case string(lilty.PUT):
			s.Lilty.Put(e.Request.Path, s.ServeHTTP)
		case string(lilty.DELETE):
			s.Lilty.Delete(e.Request.Path, s.ServeHTTP)
		case string(lilty.OPTIONS):
			s.Lilty.Options(e.Request.Path, s.ServeHTTP)
		case string(lilty.PATCH):
			s.Lilty.Patch(e.Request.Path, s.ServeHTTP)
		default:
			return errors.WithMessage(ErrUnsupportedHTTPMethod, e.Request.Method)
		}
	}

	return nil
}

func (s *server) ServeHTTP(ctxt *lilty.Context) {
	glg.Info(ctxt.Request.Method + " " + ctxt.Request.URL.String())
	for _, e := range s.endpoints {
		if e.Request.Method == ctxt.Request.Method && e.Request.Path == ctxt.Route() {
			for key, value := range e.Response.Headers {
				ctxt.SetResponseHeader(key, value)
			}

			d, err := ioutil.ReadFile(e.Response.Body)
			if err != nil {
				ctxt.SetStatusCode(http.StatusNoContent)
				return
			}

			ctxt.SetStatusCode(e.Response.Code)
			ctxt.Writer.Write(d)
		}
	}
}
