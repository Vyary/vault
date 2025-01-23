package proxy

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct{}

func New(port string) (*http.Server, error) {
  target, err := url.Parse("localhost:" + port)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	server := &http.Server{
		Addr:    ":443",
		Handler: proxy,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}

	return server, nil
}
