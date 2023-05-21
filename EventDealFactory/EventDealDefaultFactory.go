package EventDealFactory

import (
	"EventTools/EventDeal"
	"EventTools/EventDealTypeConst"
	"errors"
)

type DefaultEventDealFactory struct {
	eventMap map[string]EventDeal.DealInterface
}

// CreateEventDeal 创建事件处理实例
func (defaultEventDealFactory *DefaultEventDealFactory) CreateEventDeal(eventType string) (EventDeal.DealInterface, error) {
	defaultEventDealFactory.bindEventDeal(EventDealTypeConst.WaitCdsUrlConst, &EventDeal.WaitCdsUrl{})
	defaultEventDealFactory.bindEventDeal(EventDealTypeConst.WaitStageModelForProductionConst, &EventDeal.WaitModelForProduction{})

	// 检查是否存在对应的处理方式
	event, exist := defaultEventDealFactory.eventMap[eventType]
	if !exist {
		return nil, errors.New(eventType + " 处理的类型不存在")
	}

	return event, nil
}

// BindEventDeal 绑定事件处理方式
func (defaultEventDealFactory *DefaultEventDealFactory) bindEventDeal(eventType string, eventDeal EventDeal.DealInterface) {
	if defaultEventDealFactory.eventMap == nil {
		defaultEventDealFactory.eventMap = make(map[string]EventDeal.DealInterface, 5)
		defaultEventDealFactory.eventMap[eventType] = eventDeal
	} else {
		// 检查是否已存在，已存在时不再添加
		_, exist := defaultEventDealFactory.eventMap[eventType]
		if !exist {
			defaultEventDealFactory.eventMap[eventType] = eventDeal
		}
	}
}
