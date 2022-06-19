package api

import (
	"currency-alerter/pkg/connector"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type ServerSettings struct {
	Port uint
}

type Server struct {
	settings *ServerSettings
	con      *connector.Connector
}

func NewServer(settings *ServerSettings, user, passwd, addr, dbname string) *Server {
	return &Server{
		settings: settings,
		con:      connector.NewConnector(user, passwd, "tcp", addr, dbname),
	}
}

func middlewareCorsHeader(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%s [URL: %s] %s %s %s", r.Method, r.URL, r.Host, r.RemoteAddr, r.RequestURI)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-type, Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if (*r).Method == "OPTIONS" {
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (s *Server) Run() error {
	r := mux.NewRouter()

	alertHandler := http.HandlerFunc(s.createNewAlert)
	r.Handle("/alert", middlewareCorsHeader(alertHandler)).Methods(http.MethodPost, http.MethodOptions)

	// r.HandleFunc("/alert", s.createNewAlert).Methods(http.MethodPost)
	r.HandleFunc("/health", s.healthCheck).Methods(http.MethodGet)

	log.Infof("Starting alert api server on port %d", s.settings.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.settings.Port), r))

	return nil
}
