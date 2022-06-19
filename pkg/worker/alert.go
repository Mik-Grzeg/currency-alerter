package worker

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
