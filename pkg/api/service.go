package api

import (
	log "github.com/sirupsen/logrus"
)

const upsertAlertQuery = `
INSERT INTO alerts (money, currency, operator, email) 
VALUES(?, ?, ?, ?)
ON DUPLICATE KEY UPDATE money=?, triggered=?;
`

func (s *Server) AddAlert(alert *Alert) error {

	result, insertErr := s.con.Db.Exec(upsertAlertQuery, alert.Money, alert.Currency, alert.Operator, alert.Email, alert.Money, false)
	if insertErr != nil {
		log.Errorf("Failed to insert alert: %v with error: %v", alert, insertErr)
		return insertErr
	}

	if _, lastInsertIdErr := result.LastInsertId(); lastInsertIdErr != nil {
		return lastInsertIdErr
	}

	return nil
}
