package chat

import (
	"context"
	"log"

	"go-chat/internal/entity"
	"go-chat/internal/pkg/ichat/socket"
	"go-chat/internal/pkg/jsonutil"
	"go-chat/internal/repository/model"
)

type TalkReadMessage struct {
	Event   string `json:"event"`
	Content struct {
		MsgIds     []int `json:"msg_id"`
		ReceiverId int   `json:"receiver_id"`
	} `json:"content"`
}

// OnReadMessage 消息已读事件
func (h *Handler) OnReadMessage(ctx context.Context, client socket.IClient, data []byte) {

	var m *TalkReadMessage
	if err := jsonutil.Decode(data, &m); err != nil {
		log.Println("Chat OnReadMessage Err: ", err)
		return
	}

	h.memberService.Db().Model(&model.TalkRecords{}).
		Where("id in ? and receiver_id = ? and is_read = 0", m.Content.MsgIds, client.Uid()).
		Update("is_read", 1)

	h.redis.Publish(ctx, entity.ImTopicChat, jsonutil.Encode(entity.MapStrAny{
		"event": entity.EventTalkRead,
		"data": jsonutil.Encode(entity.MapStrAny{
			"sender_id":   client.Uid(),
			"receiver_id": m.Content.ReceiverId,
			"ids":         m.Content.MsgIds,
		}),
	}))
}
