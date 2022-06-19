package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	log "github.com/sirupsen/logrus"
)

type Mailer interface {
	NotifyViaMail(m *MailerSettings, ta *triggeredAlert) (string, error)
}

type MailGun struct {
}

type MailerSettings struct {
	Domain string
	ApiKey string
}

func (mgun *MailGun) NotifyViaMail(ms *MailerSettings, ta *triggeredAlert) (string, error) {

	mg := mailgun.NewMailgun(ms.Domain, ms.ApiKey)
	m := mg.NewMessage(
		fmt.Sprintf("Currency Alerter <postmaster@%s>", ms.Domain),
		fmt.Sprintf("Currency Alert %s", ta.alert.Currency),
		ta.alert.buildAlertMailContent(ta.currentValue),
		ta.alert.Email,
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, m)

	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("ID: %s Resp: %s\n", id, resp)
	return id, err
}
