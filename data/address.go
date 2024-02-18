package data

import (
	"errors"
	"log"
	"reverse-proxy-test/proxy"
)

type URLService struct {
	Key string
	URL string
}

type ProxyServices struct {
	Services []URLService
}

func (ps *ProxyServices) Load() {
	posts := URLService{Key: "posts", URL: "https://jsonplaceholder.typicode.com/posts"}
	users := URLService{Key: "users", URL: "https://jsonplaceholder.typicode.com/users"}
	albums := URLService{Key: "albums", URL: "https://jsonplaceholder.typicode.com/albums"}

	ps.Services = append(ps.Services, posts, users, albums)
}

func (ps *ProxyServices) GetServices() ([]URLService, error) {
	if len(ps.Services) > 0 {
		return ps.Services, nil
	}

	return nil, errors.New("no services available. try load them before")
}

func (ps *ProxyServices) ParseToProxy(services []URLService) []*proxy.URIParser {
	var proxies []*proxy.URIParser

	for _, service := range services {
		log.Printf("SERVICE %s | URL %s", service.Key, service.URL)
		if proxyParser, err := proxy.NewURIParser(service.URL); err == nil {
			proxies = append(proxies, proxyParser)
		} else {
			panic(errors.New("no address to proxy"))
		}
	}

	return proxies
}
