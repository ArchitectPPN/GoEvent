package EventDeal

import (
	"EventTools/EventParams"
	"fmt"
)

type WaitModelForProduction struct {
}

func (w *WaitModelForProduction) TriggerEvent(eventKey string, eventType string, listenEventParam EventParams.ListenEventParams) bool {
	fmt.Println("触发等待来等模型事件: ["+eventType+"] eventKey: ["+eventKey+"] eventListenParams: [", listenEventParam, "]")

	return true
}

func (w *WaitModelForProduction) RestoreEvent(eventKey string, eventType string) bool {
	fmt.Println("恢复等待来等模型事件： eventKey: [", eventKey, "] eventType: [", eventType, "]")

	return true
}
