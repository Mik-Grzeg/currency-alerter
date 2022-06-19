package worker

import (
	"currency-alerter/pkg/connector"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type WorkerSettings struct {
	ClientTimeout  uint
	ScrapeInterval uint
}

type Worker struct {
	con      *connector.Connector
	client   *http.Client
	settings *WorkerSettings
}

func NewWorker(settings *WorkerSettings, user, passwd, addr, dbname string) *Worker {

	return &Worker{
		settings: settings,
		con:      connector.NewConnector(user, passwd, "tcp", addr, dbname),
		client:   &http.Client{Timeout: time.Duration(settings.ClientTimeout) * time.Second},
	}
}

func (w *Worker) Run() error {
	currencies := make(map[string]float32)
	recheck := make(chan bool)

	go func() {
		for {
			tables := w.getNBPRates()
			tables.ToMap(currencies, recheck)

			time.Sleep(time.Duration(w.settings.ScrapeInterval) * time.Second)
		}
	}()

	for {
		select {
		case <-recheck:
			checkAlerts(w, currencies)
			log.Info("Running alerts recheck after fetching data from nbp api")
		case <-time.After(time.Duration(w.settings.ScrapeInterval/3) * time.Second):
			checkAlerts(w, currencies)
			log.Infof("Running alerts recheck after %dsec", int(w.settings.ScrapeInterval/3))
		}
	}

	return nil
}
