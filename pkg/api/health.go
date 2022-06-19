package api

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func RunHealthCheck(port uint) error {
	resp, reqErr := http.Get(fmt.Sprintf("http://localhost:%d/health", port))
	if reqErr != nil {
		return reqErr
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("Invalid health status: %s", resp.Status)
	}
	return nil
}

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	if connectorErr := s.con.HealthCheck(); connectorErr != nil {
		log.Debugf("Health check connector error: %v", connectorErr)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
