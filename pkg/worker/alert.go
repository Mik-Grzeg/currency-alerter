package worker

import "fmt"

type Alert struct {
	Id        uint
	Money     float32
	Currency  string
	Operator  string
	Email     string
	Triggered bool
}

func (a *Alert) isAlertTriggered(currentRate float32) bool {
	switch {
	case a.Operator == ">":
		return currentRate > a.Money
	case a.Operator == "<":
		return currentRate < a.Money
	default:
		return false
	}
}

func (a *Alert) buildAlertMailContent(currentValue float32) string {
	var operatorToString string
	if operatorToString = "lower"; a.Operator == ">" {
		operatorToString = "greater"
	}

	return fmt.Sprintf(
		"Hi, alert that you have set up has been triggered.\n%s exchange rate is %.2f. You expected it to be %s than %.2f\nBye",
		a.Currency, currentValue, operatorToString, a.Money)
}
