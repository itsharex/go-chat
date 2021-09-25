package cache

import (
	"context"
	"fmt"

	"go-chat/connect"
)

type WsClient struct {
	Redis *connect.Redis
}

// Set 设置客户端与用户绑定关系
// channel  渠道分组
// uuid     客户端连接ID
// id       用户ID
func (w *WsClient) Set(ctx context.Context, channel string, uuid string, id int) {
	flag := fmt.Sprintf("ws:channel:%s:client", channel)
	w.Redis.Client.HSet(ctx, flag, uuid, id)

	flag = fmt.Sprintf("ws:channel:%s:user:%d", channel, id)
	w.Redis.Client.SAdd(ctx, flag, uuid)
}

// Del 删除客户端与用户绑定关系
// channel  渠道分组
// uuid     客户端连接ID
func (w *WsClient) Del(ctx context.Context, channel string, uuid string) {
	flag := fmt.Sprintf("ws:channel:%s", channel)

	id, _ := w.Redis.Client.HGet(ctx, flag, uuid).Result()

	w.Redis.Client.HDel(ctx, flag, uuid)

	flag = fmt.Sprintf("ws:channel:%s:user:%s", channel, id)
	w.Redis.Client.SRem(ctx, flag, uuid)
}

// IsOnline 判断客户端是否在线[当前机器]
// channel  渠道分组
// id       用户ID
func (w *WsClient) IsOnline(ctx context.Context, channel string, id string) bool {
	flag := fmt.Sprintf("ws:channel:%s:user:%s", channel, id)

	val, err := w.Redis.Client.SCard(ctx, flag).Result()
	if err != nil {
		return false
	}

	return val > 0
}

// IsOnlineAll 判断客户端是否在线[所有部署机器]
// channel  渠道分组
// id       用户ID
func (w *WsClient) IsOnlineAll(ctx context.Context, channel string, id string) bool {
	flag := fmt.Sprintf("ws:channel:%s:user:%s", channel, id)
	val, err := w.Redis.Client.SCard(ctx, flag).Result()
	if err != nil {
		return false
	}

	return val > 0
}