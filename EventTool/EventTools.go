package EventTool

import (
	"EventTools/EventDealFactory"
	"EventTools/EventParams"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type EventTool struct {
	rdsDb            *redis.Client
	eventDealFactory EventDealFactory.BaseEventDealFactoryInterface
}

// ListenEvent 监听事件
func (e *EventTool) ListenEvent(eventType string, eventKey string, params EventParams.ListenEventParams) {

	paramsStr, _ := json.Marshal(params)

	redisEventKey := eventType + ":" + eventKey
	e.rdsDb.Set(redisEventKey, string(paramsStr), 0)
}

// TriggerEvent 触发事件
func (e *EventTool) TriggerEvent(eventType string, eventKey string, params EventParams.ListenEventParams) {
	redisEventKey := eventType + ":" + eventKey

	listenParam, _ := e.rdsDb.Get(redisEventKey).Result()
	if listenParam == "" {
		fmt.Println(redisEventKey + ": 监听事件不存在")
		return
	}

	var listenParams EventParams.ListenEventParams
	_ = json.Unmarshal([]byte(listenParam), &listenParams)

	eventObj, err := e.eventDealFactory.CreateEventDeal(eventType)
	if err != nil {
		return
	}

	dealRes := eventObj.TriggerEvent(eventKey, eventType, listenParams)
	if dealRes {
		fmt.Println(redisEventKey + ": 处理成功!")
		// 处理成功
		e.rdsDb.Del(redisEventKey)
		return
	}

	// 处理失败
	fmt.Println(redisEventKey + ": 处理失败!")
}

// CancelEvent 触发事件
func (e *EventTool) CancelEvent() {

}

// RestoreEvent 触发事件
func (e *EventTool) RestoreEvent() {

}

// SetEventDealFactoryInterface 设置事件处理工厂
func (e *EventTool) SetEventDealFactoryInterface(eventDealFactory EventDealFactory.BaseEventDealFactoryInterface) {
	e.eventDealFactory = eventDealFactory
}

// SetRedisConnect 设置redis连接
func (e *EventTool) SetRedisConnect(rds *redis.Client) {
	e.rdsDb = rds
}
