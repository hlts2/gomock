package gomock

import (
	"bytes"
	"regexp"
	"strings"
)

// newRegepxRoute parse a route and returns a routeRegexp
func newRegepxRoute(route string) (*regexp.Regexp, error) {
	cnt := strings.Count(route, "?")
	var routetpl = make([]byte, 0, len(route)+cnt)

	for i, uPoint := range routetpl {
		switch ch := string(uPoint); {
		case ch == "?":
			routetpl = append(routetpl, '\\', route[i])
		case ch == "}":
			cnt := 0
			for _, v := range routetpl {
				if string(v) == "{" {
					break
				}
				cnt++
			}

			routetpl = routetpl[:cnt]
			routetpl = append(routetpl, '.', '*', '?')
		default:
			for _, v := range []byte(string(uPoint)) {
				routetpl = append(routetpl, v)
			}
		}
	}

	return regexp.Compile(bytes.NewBuffer(routetpl).String())
}
