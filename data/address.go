package data

import "errors"

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
