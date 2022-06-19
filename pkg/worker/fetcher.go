package worker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Tables []struct {
	Table         string       `json:"table"`
	No            string       `json:"no"`
	EffectiveDate string       `json:"effectiveDate"`
	Rates         []Currencies `json:"rates"`
}

func (t *Tables) ToMap(c map[string]float32, reprocess chan bool) {
	for _, table := range *t {
		for _, rate := range table.Rates {
			c[rate.Code] = rate.Mid
		}
	}
	reprocess <- true
}

type Currencies struct {
	Currency string  `json:"currency"`
	Code     string  `json:"code"`
	Mid      float32 `json:"mid"`
}

func (w *Worker) getNBPRates() (*Tables, error) {
	endpoint := "http://api.nbp.pl/api/exchangerates/tables/A"

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Errorf("Error occurred while creating new request. %+v", err)
		return nil, err
	}

	response, err := w.client.Do(req)
	if err != nil {
		log.Errorf("Error sending request to API endpoint. %+v", err)
		return nil, err
	}
	// Close the connection to reuse it
	defer response.Body.Close()
	log.Infof("NBP api responded with: %d status code", response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//Failed to read response.
		log.Errorf("There was an error reading response body: %+v", err)
		return nil, err
	}

	var tables Tables
	if errBodyDecode := json.Unmarshal(body, &tables); errBodyDecode != nil {
		log.Errorf("Error parsing body %+v", errBodyDecode)
		return nil, errBodyDecode
	}
	return &tables, nil
}
