package EventToolFactory

import (
	"EventTools/EventDealFactory"
	"EventTools/EventTool"
	"fmt"
	"github.com/go-redis/redis"
)

type Factory struct {
}

func (etf *Factory) CreatEventTools() (*EventTool.EventTool, error) {
	tools := new(EventTool.EventTool)

	rds, err := etf.getRedisConnect()
	if err != nil {
		return tools, err
	}

	// 设置redis连接
	tools.SetRedisConnect(rds)
	// 设置处理
	tools.SetEventDealFactoryInterface(&EventDealFactory.DefaultEventDealFactory{})

	return tools, err
}

// getRedisConnect 获取redis连接
func (etf *Factory) getRedisConnect() (rds *redis.Client, err error) {
	rds = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	res, err := rds.Ping().Result()
	if err != nil {
		fmt.Println("连接redis 失败, err:", err)
		return rds, err
	}

	fmt.Println("连接redis 信息", res)

	return rds, err
}
