package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type URIParser struct {
	URI   *url.URL
	Proxy *httputil.ReverseProxy
}

func NewURIParser(uri string) (*URIParser, error) {
	parsedUri, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	return &URIParser{
		URI:   parsedUri,
		Proxy: httputil.NewSingleHostReverseProxy(parsedUri),
	}, nil
}

func (u *URIParser) Handle(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	r.Host = u.URI.Host
	u.Proxy.ServeHTTP(w, r)
}
