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
	MailerSettings *MailerSettings
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
	triggeredAlerts := make(chan triggeredAlert)

	mailer := MailGun{}

	go func() {
		for {
			tables, fetcherErr := w.getNBPRates()
			if fetcherErr != nil {
				continue
			}

			tables.ToMap(currencies, recheck)

			time.Sleep(time.Duration(w.settings.ScrapeInterval) * time.Second)
		}
	}()

	for {
		select {
		case <-recheck:
			go func() {
				checkAlerts(w, currencies, triggeredAlerts)
			}()
			log.Info("Running alerts recheck after fetching data from nbp api")
		case t := <-triggeredAlerts:
			go func(trigAlert *triggeredAlert) {
				log.Debug("Spawned go routine for notyfing")
				mailer.NotifyViaMail(w.settings.MailerSettings, trigAlert)
			}(&t)
		case <-time.After(time.Duration(w.settings.ScrapeInterval/3) * time.Second):
			go func() {
				checkAlerts(w, currencies, triggeredAlerts)
			}()
			log.Infof("Running alerts recheck after %dsec", int(w.settings.ScrapeInterval/3))
		}
	}

}
