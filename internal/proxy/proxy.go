package proxy

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewSecure(targetURL string) (*http.Server, error) {
	target, err := url.Parse(targetURL)
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

func New() *http.Server {
	return &http.Server{
		Addr:    ":80",
		Handler: redirectToHTTPS(),
	}
}

func redirectToHTTPS() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpsURL := "https://" + r.Host + r.URL.RequestURI()
		http.Redirect(w, r, httpsURL, http.StatusMovedPermanently)
	})
}
