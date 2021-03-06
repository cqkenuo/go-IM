package service

import (
	"go-IM/pkg/defs"
	"testing"
)

// 模拟系统消息发送给用户
func Test_MockSystemSend(t *testing.T) {
	sender := defs.Sender{
		AppId:      1,
		SenderType: defs.SenderType_ST_SYSTEM,
		SenderId:   0,
		DeviceId:   0,
	}
	messageReq := defs.SendMessageReq{
		ReceiverType:   defs.ReceiverType_RT_USER,
		ReceiverId:     54146910402904065,
		MessageType:    defs.Message_TEXT,
		MessageContent: "hello world",
		ToUserIds:      []string{},
		IsPersist:      true,
	}
	err := MessageService.Send(0, sender, messageReq)
	if err != nil {
		panic(err)
	}
}
