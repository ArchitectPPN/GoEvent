package EventDealFactory

import (
	"EventTools/EventDeal"
)

type BaseEventDealFactoryInterface interface {
	CreateEventDeal(eventType string) (EventDeal.DealInterface, error)
}
