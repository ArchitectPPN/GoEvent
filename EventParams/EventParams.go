package EventParams

import "time"

type ListenEventParams struct {
	MsgId          string `json:"msg_id"`                     // 消息ID
	TimeStamp      int64  `json:"time_stamp,omitempty"`       // 时间
	EventType      string `json:"event_type"`                 // 事件类型
	EventKey       string `json:"event_key"`                  // 事件Key
	CancelEventKey string `json:"cancel_event_key,omitempty"` // 取消事件Key
	ListenStatus   int    `json:"listen_status"`              // 监听状态
	Expire         int64  `json:"expire,omitempty"`           // 过期时间
	Params         string `json:"params"`                     // 监听参数
}

type ListenParam struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

type TriggerEvent struct {
	MsgId         string    // 消息ID
	TimeStamp     time.Time // 时间
	Params        string    // 监听参数
	EventKey      string    // 事件Key
	TriggerStatus int       // 触发状态
}
