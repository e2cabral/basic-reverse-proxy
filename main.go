package main

import (
	"net/http"
	"reverse-proxy-test/data"
	"slices"
)

func main() {
	var proxyServices data.ProxyServices
	proxyServices.Load()
	services, err := proxyServices.GetServices()
	if err != nil {
		panic(err)
	}

	proxies := proxyServices.ParseToProxy(services)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index := slices.IndexFunc(services, func(s data.URLService) bool {
			return r.Header.Get("X-Route-To") == s.Key
		})
		if index == -1 {
			http.Error(w, "Invalid service key", http.StatusBadRequest)
			return
		}
		proxies[index].Handle(w, r)
	})
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
