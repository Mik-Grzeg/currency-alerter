package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (s *Server) createNewAlertHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)

		var newAlert Alert
		json.Unmarshal(reqBody, &newAlert)
		w.Header().Set("Content-Type", "application/json")

		createdAlertErr := s.AddAlert(&newAlert)
		if createdAlertErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("err: Creating an alert failed. Something went wrong")
		} else {
			log.Debugf("Configured an alert: %v ", newAlert)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newAlert)
		}
	})
}
