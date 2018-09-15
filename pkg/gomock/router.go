package gomock

import (
	"io/ioutil"
	"net/http"
)

// Router represetns HTTP router interface
type Router interface {
	Get(path string, response Response)
	Head(path string, response Response)
	Post(path string, response Response)
	Put(path string, response Response)
	Delete(path string, response Response)
	Options(path string, response Response)
	Patch(path string, response Response)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type router struct {
	trees     []Trie
	endpoints []Endpoint
}

func (r *router) Get(path string, response Response) {
	r.trees[Get].Insert(path, response)
}

func (r *router) Head(path string, response Response) {
	r.trees[Head].Insert(path, response)
}

func (r *router) Post(path string, response Response) {
	r.trees[Post].Insert(path, response)
}

func (r *router) Put(path string, response Response) {
	r.trees[Put].Insert(path, response)
}

func (r *router) Delete(path string, response Response) {
	r.trees[Delete].Insert(path, response)
}

func (r *router) Options(path string, response Response) {
	r.trees[Options].Insert(path, response)
}

func (r *router) Patch(path string, response Response) {
	r.trees[Patch].Insert(path, response)
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	resp, ok := r.getRoute(req.URL.String(), req.Method)
	if !ok {
		return
	}

	if resp.Code < 100 || resp.Code > 500 {
		return
	}

	for key, value := range resp.Headers {
		w.Header().Add(key, value)
	}

	d, err := ioutil.ReadFile(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(resp.Code)
	w.Write(d)
}

// NewRouter returns Router(*router) object
func NewRouter(endPoints []Endpoint) Router {
	r := &router{
		trees:     make([]Trie, 0, HTTPMethodCount),
		endpoints: endPoints,
	}

	for _, endPoint := range endPoints {
		r.addRoute(endPoint)
	}

	return r
}

func (r *router) addRoute(endpoint Endpoint) {
	switch endpoint.Request.Method {
	case Get.String():
		r.Get(endpoint.Request.Path, endpoint.Response)
	case Head.String():
		r.Head(endpoint.Request.Path, endpoint.Response)
	case Post.String():
		r.Post(endpoint.Request.Path, endpoint.Response)
	case Put.String():
		r.Put(endpoint.Request.Path, endpoint.Response)
	case Delete.String():
		r.Delete(endpoint.Request.Path, endpoint.Response)
	case Patch.String():
		r.Patch(endpoint.Request.Path, endpoint.Response)
	}
}

func (r *router) getRoute(path, method string) (Response, bool) {
	switch method {
	case Get.String():
		return r.trees[Get].Search(path)
	case Head.String():
		return r.trees[Head].Search(path)
	case Post.String():
		return r.trees[Post].Search(path)
	case Put.String():
		return r.trees[Put].Search(path)
	case Delete.String():
		return r.trees[Delete].Search(path)
	case Patch.String():
		return r.trees[Patch].Search(path)
	default:
		return Response{}, false
	}
}
