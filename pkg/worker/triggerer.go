package worker

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

const getAllAlertsQuery = `SELECT id, money, currency, operator, email FROM alerts WHERE triggered = false;`

func (w *Worker) getAlerts() []*Alert {
	rows, queryErr := w.con.Db.Query(getAllAlertsQuery)
	if queryErr != nil {
		log.Errorf("There has been error while processing query: %v", queryErr)
	}

	defer rows.Close()

	alerts := []*Alert{}
	for rows.Next() {
		var (
			id       uint
			money    float32
			currency string
			operator string
			email    string
		)

		rows.Scan(&id, &money, &currency, &operator, &email)
		alerts = append(alerts, &Alert{
			Id:        id,
			Money:     money,
			Currency:  currency,
			Operator:  operator,
			Email:     email,
			Triggered: false,
		})
	}

	return alerts
}

func buildTriggeredUpdateQuery(ids []uint) ([]interface{}, string) {
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	query := `UPDATE alerts SET triggered = true WHERE id IN(?` + strings.Repeat(",?", len(args)-1) + `)`
	return args, query
}

func (w *Worker) triggerAlerts(alerts []*Alert, c map[string]float32) error {
	ids := make([]uint, 0, len(alerts))
	for _, alert := range alerts {
		if !alert.Triggered && alert.isAlertTriggered(c[alert.Currency]) {
			ids = append(ids, alert.Id)
			log.Debugf("Triggering : for %s\n with state %t with condition x %s %f Current rate of %s %f", alert.Email, alert.Triggered, alert.Operator, alert.Money, alert.Currency, c[alert.Currency])
		}
	}

	if len(ids) == 0 {
		return nil
	}

	idsToPass, triggeringQuery := buildTriggeredUpdateQuery(ids)
	results, updateErr := w.con.Db.Exec(triggeringQuery, idsToPass...)
	if updateErr != nil {
		log.Errorf("Failed executing updating query: %v", updateErr)
		return updateErr
	}

	rowsAffected, rowsAffectedErr := results.RowsAffected()
	if rowsAffectedErr != nil {
		log.Errorf("Could not retrieve how many rows were affected", updateErr)
		return updateErr
	}

	if rowsAffected != int64(len(idsToPass)) {
		log.Warningf("Number (%d) of updated triggered alerts does not match passed amount (%d)", rowsAffected, len(idsToPass))
	}

	return nil
}

func checkAlerts(w *Worker, c map[string]float32) {
	alerts := w.getAlerts()
	log.Info("Checking alerts ...")

	w.triggerAlerts(alerts, c)
}
