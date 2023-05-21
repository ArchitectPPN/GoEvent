package EventTool

import "EventTools/EventParams"

type EventToolsInterface interface {
	// ListenEvent 监听事件
	ListenEvent(eventType string, eventKey string, params EventParams.ListenEventParams)

	// TriggerEvent 触发事件
	TriggerEvent(eventType string, eventKey string, params EventParams.ListenEventParams)

	// CancelEvent 取消监听
	CancelEvent()

	// RestoreEvent 恢复监听
	RestoreEvent()
}
