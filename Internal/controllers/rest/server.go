package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HTTPServer struct {
	server       *http.Server
	proxyHandler *ProxyHandler
}

func NewHTTPServer(proxyHandler *ProxyHandler) *HTTPServer {
	return &HTTPServer{proxyHandler: proxyHandler}
}

func (s HTTPServer) Start(port string) error {
	s.server = &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      s.router(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.server.ListenAndServe()
}

func (s HTTPServer) router() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]bool{"Are your server still alive?": true})

		w.WriteHeader(http.StatusOK)
	})

	//router.HandleFunc("/api/proxy", s.proxyHandler.DO).Methods(http.MethodGet)
	router.Path("/api/proxy").HandlerFunc(s.proxyHandler.Do).Methods(http.MethodGet)

	router.Path("/api/proxy/list").HandlerFunc(s.proxyHandler.List).Methods(http.MethodGet)
	// TODO: Get by key

	return router
}
