package connector

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type Connector struct {
	Db  *sql.DB
	cfg mysql.Config
}

func (c *Connector) HealthCheck() error {
	var pingErr error
	for i := 1; i < 4; i++ {
		time.Sleep(time.Duration(i*i) * time.Second)
		pingErr = c.Db.Ping()

		if pingErr == nil {
			return nil
		}
		log.Infof("Pinging database - retry: %d", i)
	}

	if pingErr != nil {
		log.Errorf("Database does not response to pings: %v", pingErr)
		return pingErr
	}
	return nil
}

func NewConnector(user, passwd, net, addr, dbname string) *Connector {
	cfg := mysql.Config{
		User:   user,
		Passwd: passwd,
		Net:    net,
		Addr:   addr,
		DBName: dbname,
	}
	log.Debug(cfg.FormatDSN())

	db, dbErr := sql.Open("mysql", cfg.FormatDSN())
	if dbErr != nil {
		log.Fatalf("Connection with database cannot be established: %v", dbErr)
	}

	log.Info("Database connector created")

	connection := Connector{
		Db:  db,
		cfg: cfg,
	}

	if errMysqlHealthCheck := connection.HealthCheck(); errMysqlHealthCheck != nil {
		return nil
	}
	connection.migrate()

	return &connection
}

const migrateAlertTableQuery = `
CREATE TABLE IF NOT EXISTS alerts(
	id	INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	money DECIMAL(5,2) NOT NULL,
	currency VARCHAR(128) NOT NULL,
	operator CHAR(1) NOT NULL,
	email VARCHAR(255) NOT NULL,
	triggered BOOLEAN DEFAULT false NOT NULL,
	CONSTRAINT UC_Alert UNIQUE(currency, email, operator)
);
`

func (c *Connector) migrate() {
	_, insertErr := c.Db.Exec(migrateAlertTableQuery)
	if insertErr != nil {
		log.Fatalf("Failed to create table: %v", insertErr)
	}

	log.Debugf("Table schema migrated")
}
