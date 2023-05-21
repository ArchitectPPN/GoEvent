package main

import (
	"EventTools/EventDealTypeConst"
	"EventTools/EventParams"
	"EventTools/EventToolFactory"
	"github.com/google/uuid"
)

func main() {
	listenAndTrigger()
}

func listenAndTrigger() {
	event := new(EventToolFactory.Factory)
	eventTools := event.CreatEventTools()

	eventKey := uuid.New().String()
	// 设置监听事件的参数
	listenEventParams := EventParams.ListenEventParams{
		MsgId:     uuid.New().String(),
		EventKey:  eventKey,
		EventType: EventDealTypeConst.WaitCdsUrlConst,
	}

	// 添加监听事件
	eventTools.ListenEvent(EventDealTypeConst.WaitCdsUrlConst, eventKey, listenEventParams)

	// 触发监听事件
	eventTools.TriggerEvent(EventDealTypeConst.WaitCdsUrlConst, eventKey, listenEventParams)
}
