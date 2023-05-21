package EventDeal

import (
	"EventTools/EventParams"
	"fmt"
)

type WaitCdsUrl struct {
}

func (w *WaitCdsUrl) TriggerEvent(eventKey string, eventType string, listenEventParam EventParams.ListenEventParams) bool {
	fmt.Println("触发等待上传Cds地址事件: ["+eventType+"] eventKey: ["+eventKey+"] eventListenParams: [", listenEventParam, "]")

	return true
}

func (w *WaitCdsUrl) RestoreEvent(eventKey string, eventType string) bool {
	fmt.Println("恢复等待上传Cds地址事件： eventKey: [", eventKey, "] eventType: [", eventType, "]")

	return true
}
