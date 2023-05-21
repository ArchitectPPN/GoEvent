package main

import (
	"EventTools/EventDealTypeConst"
	"EventTools/EventParams"
	"EventTools/EventToolFactory"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	listenAndTrigger()
}

func listenAndTrigger() {
	event := new(EventToolFactory.Factory)
	eventTools, err := event.CreatEventTools()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

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
