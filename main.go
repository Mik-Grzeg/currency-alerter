package main

import (
	"currency-alerter/pkg/api"
	"currency-alerter/pkg/worker"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var mysqlFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "mysql-user",
		Usage:   "mysql user login",
		Value:   "default",
		EnvVars: []string{"MYSQL_USER"},
	},
	&cli.StringFlag{
		Name:    "db-name",
		Usage:   "database name",
		Value:   "default",
		EnvVars: []string{"MYSQL_DATABASE"},
	},
	&cli.StringFlag{
		Name:    "db-addr",
		Usage:   "database address",
		Value:   "mysql:3306",
		EnvVars: []string{"DB_ADDRESS"},
	},
	&cli.StringFlag{
		Name:    "mysql-passwd",
		Usage:   "mysql user password",
		Value:   "password",
		EnvVars: []string{"MYSQL_PASSWORD"},
	},
}

var serverPort = &cli.UintFlag{
	Name:  "port",
	Usage: "Port that the server will be server on",
	Value: 8000,
}

func main() {
	app := &cli.App{
		Name:                 "currency-alerter",
		Usage:                "",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "api",
				Usage: "Start web server",
				Flags: append(
					mysqlFlags,
					serverPort,
					&cli.BoolFlag{
						Name:  "health",
						Usage: "Running health check",
						Value: false,
					},
				),
				Action: func(c *cli.Context) error {
					if c.Bool("health") {
						return api.RunHealthCheck(c.Uint("port"))
					}

					return api.NewServer(
						&api.ServerSettings{
							Port: c.Uint("port"),
						},
						c.String("mysql-user"),
						c.String("mysql-passwd"),
						c.String("db-addr"),
						c.String("db-name"),
					).Run()
				},
			},
			{
				Name:  "worker",
				Usage: "Asynchronous worker",
				Flags: append(
					mysqlFlags,
					&cli.UintFlag{
						Name:  "http-client-timeout",
						Usage: "Time until the connection will be terminated",
						Value: 10,
					},
					&cli.UintFlag{
						Name:  "nbp-fetch-interval",
						Usage: "Interval between api calls in order to retrieve nbp rates",
						Value: 60 * 30,
					},
					&cli.StringFlag{
						Name:    "mailgun-domain",
						Usage:   "MailGun domain",
						EnvVars: []string{"MAILGUN_DOMAIN"},
					},
					&cli.StringFlag{
						Name:    "mailgun-apikey",
						Usage:   "MailGun api key",
						EnvVars: []string{"MAILGUN_APIKEY"},
					},
				),
				Action: func(c *cli.Context) error {
					return worker.NewWorker(
						&worker.WorkerSettings{
							ClientTimeout:  c.Uint("http-client-timeout"),
							ScrapeInterval: c.Uint("nbp-fetch-interval"),
							MailerSettings: &worker.MailerSettings{
								Domain: c.String("mailgun-domain"),
								ApiKey: c.String("mailgun-apikey"),
							},
						},
						c.String("mysql-user"),
						c.String("mysql-passwd"),
						c.String("db-addr"),
						c.String("db-name"),
					).Run()
				},
			},
		},
		Before: setLogLevel,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Be more verbose when logging stuff",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func setLogLevel(c *cli.Context) error {
	if c.IsSet("debug") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	return nil
}
