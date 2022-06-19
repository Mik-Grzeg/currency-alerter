package api

import (
	log "github.com/sirupsen/logrus"
)

const upsertAlertQuery = `
INSERT INTO alerts (money, currency, operator, email) 
VALUES(?, ?, ?, ?)
ON DUPLICATE KEY UPDATE money=?, triggered=?;
`

func (s *Server) AddAlert(alert *Alert) (*int64, error) {

	result, insertErr := s.con.Db.Exec(upsertAlertQuery, alert.Money, alert.Currency, alert.Operator, alert.Email, alert.Money, false)
	if insertErr != nil {
		log.Errorf("Failed to insert alert: %v with error: %v", alert, insertErr)
		return nil, insertErr
	}

	lastInsertId, lastInsertIdErr := result.LastInsertId()
	if lastInsertIdErr != nil {
		return nil, lastInsertIdErr
	}
	log.Debugf("New data inserted to alert with id: %d", lastInsertId)

	return &lastInsertId, nil
}
