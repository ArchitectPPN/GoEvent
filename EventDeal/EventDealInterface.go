package EventDeal

import "EventTools/EventParams"

type DealInterface interface {
	TriggerEvent(eventKey string, eventType string, listenEventParam EventParams.ListenEventParams) bool
	RestoreEvent(eventKey string, eventType string) bool
}
