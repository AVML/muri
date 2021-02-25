// muri URIs can be used to consume or reference hypermedia resources
// bundled inside another resource.
// Example(the encoding is ignored for brevity) :
// muri://http://example.com/a.zip,a.mp4
// The muri above points to a resource ("example.com/a.zip") and
// and a subresource within it. A client should download the parent
// resource (a.zip)
// and render the subresources (a.mp4) to resolve the muri.
package muri

import (
	"errors"
	"net/url"
	"strings"
)

const Scheme = "muri://"

// Encode receives a succesion of URIs (parent -> children)
// and returns a muri
func Encode(sa ...string) string {
	switch len(sa) {
	case 1:
		return sa[0]
	case 0:
		return ""
	}
	for k := range sa {
		sa[k] = url.QueryEscape(sa[k])
	}
	s := strings.Join(sa, ",")
	return Scheme + url.QueryEscape(s)
}

func Decode(s string) ([]string, error) {
	if !strings.HasPrefix(s, Scheme) {
		// it's actually a URI?
		return []string{s}, nil
	}
	s = strings.TrimPrefix(s, Scheme)
	ps, err := url.QueryUnescape(s)
	if err != nil {
		return nil, err
	}
	if ps == "" {
		return nil, errors.New("Invalid muri")
	}
	sa := strings.Split(ps, ",")
	if len(sa) == 0 {
		return nil, errors.New("Invalid muri")
	}
	for k := range sa {
		sa[k], err = url.QueryUnescape(sa[k])
		if err != nil {
			return nil, err
		}
	}
	return sa, nil
}

func AddParent(parent, child string) (string, error) {
	sa, err := Decode(parent)
	if err != nil {
		return "", err
	}
	sa = append(sa, child)
	return Encode(sa...), nil

}
