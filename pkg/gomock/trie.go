package gomock

import (
	"unsafe"
)

// Trie is base trie tree interface{}
type Trie interface {
	Insert(path string, response Response)
	Search(path string) (Response, bool)
}

type trie struct {
	root *node
}

type node struct {
	child    *node
	bros     *node
	path     string
	partPath string
	queries  Queries
	response Response
}

// NewTrie returns Trie(*trie) object
func NewTrie() Trie {
	return &trie{
		root: new(node),
	}
}

func (t *trie) Insert(path string, response Response) {
	if len(path) == 0 || path[0] != '/' {
		return
	}

	if len(path) == 1 {
		t.root.update("/", path, Queries{}, response)
		return
	}

	t.insert(path, response)
}

func (t *trie) insert(path string, response Response) {
	bytesPath := stringToBytes(path)

	n := t.root.update("/", path, Queries{}, Response{})

	for i := 1; i < len(bytesPath); i++ {
		startPos, endPos := blockPos(bytesPath, i)

		var queries Queries
		if len(bytesPath) > endPos && bytesPath[endPos] == '?' {
			queries = parseQuery(bytesToString(bytesPath, endPos+1, len(bytesPath)))
			if startPos == endPos {
				n.queries = queries
				n.response = response
			} else {
				n.update(bytesToString(bytesPath, startPos, endPos), path, queries, response)
			}
			break
		}

		n = n.update(bytesToString(bytesPath, startPos, endPos), path, queries, response)

		i = endPos
	}
}

// nextBlockPos returns start pos of next block
func blockPos(path []byte, startPos int) (int, int) {
	endPos := startPos
	for i := startPos; i < len(path) && path[i] != '/' && path[i] != '?'; i++ {
		endPos++
	}

	return startPos, endPos
}

func (t *trie) Search(path string) (Response, bool) {
	if len(path) == 0 || path[0] != '/' {
		return Response{}, false
	}

	n := t.root

	if n = n.get("/"); n == nil {
		return Response{}, false
	}

	if len(path) == 1 {
		return n.response, true
	}

	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	return t.search(path, n)
}

func (t *trie) search(path string, n *node) (Response, bool) {
	bytesPath := stringToBytes(path)

	for i := 1; i < len(bytesPath); i++ {
		startPos, endPos := blockPos(bytesPath, i)

		var queries Queries
		if len(bytesPath) > endPos && bytesPath[endPos] == '?' {
			queries = parseQuery(bytesToString(bytesPath, endPos+1, len(bytesPath)))
			if startPos == endPos {
				if n.queries.Match(queries) {
					return n.response, true
				}
			}
		}

		if n = n.get(bytesToString(bytesPath, startPos, endPos)); n == nil {
			return Response{}, false
		}

		if len(bytesPath) > endPos && bytesPath[endPos] == '?' && n.queries.Match(queries) {
			return n.response, true
		}

		if endPos >= len(bytesPath) {
			return n.response, true
		}

		i = endPos
	}

	return Response{}, false
}

func (n *node) get(partPath string) *node {
	child := n.child
	for child != nil {
		if child.partPath == partPath || child.partPath == "*" {
			return child
		}
		child = child.bros
	}

	return child
}

func (n *node) insert(partPath, path string, queries []Query, response Response) *node {
	newNode := &node{
		bros:     n.child,
		path:     path,
		partPath: partPath,
		response: response,
	}

	if len(queries) != 0 {
		newNode.queries = queries
	}

	n.child = newNode
	return newNode
}

func (n *node) update(partPath, path string, queries []Query, response Response) *node {
	child := n.get(partPath)
	if child == nil {
		child = n.insert(partPath, path, queries, response)
	}

	return child
}

func stringToBytes(path string) []byte {
	return *(*[]byte)(unsafe.Pointer(&path))
}

func bytesToString(b []byte, start, end int) string {
	return (*(*string)(unsafe.Pointer(&b)))[start:end]
}
