package api

import (
	"currency-alerter/pkg/api/middleware"
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

func (s *Server) Run() error {
	r := mux.NewRouter()

	r.Handle("/alert", middleware.MiddlewareCorsHeader(middleware.MiddlewareLogging(s.createNewAlertHandler()))).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/health", middleware.MiddlewareLogging(s.HealthCheckHanlder())).Methods(http.MethodGet)

	log.Infof("Starting server on port :%d", s.settings.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.settings.Port), r))

	return nil
}
