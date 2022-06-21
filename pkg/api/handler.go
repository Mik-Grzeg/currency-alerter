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
		w.Header().Set("Content-Type", "application/json")

		var newAlert Alert
		if unmarshallErr := json.Unmarshal(reqBody, &newAlert); unmarshallErr != nil {
			log.Errorf("Serializing received data has failed: %+v", unmarshallErr)
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode("err: Serializing received data has failed")
		}

		createdAlertErr := s.AddAlert(&newAlert)
		if createdAlertErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode("err: Creating an alert failed. Something went wrong")
		} else {
			log.Debugf("Configured an alert: %v ", newAlert)

			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(newAlert)
		}
	})
}
