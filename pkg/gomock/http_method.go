package gomock

// HTTPMethod is custom type for HTTP method
type HTTPMethod int

// HTTPMethodCount is the number of available HTTP method
const HTTPMethodCount = 7

// common HTTP method
const (
	Get HTTPMethod = iota
	Head
	Post
	Put
	Delete
	Options
	Patch
)

// String convert from HTTPMethod to string type.
func (h HTTPMethod) String() string {
	switch h {
	case Get:
		return "GET"
	case Head:
		return "HEAD"
	case Post:
		return "POST"
	case Put:
		return "PUT"
	case Delete:
		return "DELETE"
	case Options:
		return "OPTIONS"
	case Patch:
		return "PATCH"
	default:
		return ""
	}
}
