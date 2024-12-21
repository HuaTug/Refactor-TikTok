package service

import (
	"context"

	"HuaTug.com/cmd/message/dal/db"
	"HuaTug.com/cmd/message/service/convert"

	"HuaTug.com/kitex_gen/messages"
	"HuaTug.com/pkg/errno"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{
		ctx: ctx,
	}
}

func (s *MessageService) PushMessage(request *messages.InsertMessageRequest) error {
	if err := db.InsertMessage(request.Message.FromUid, request.Message.ToUid, request.Message.Content); err != nil {
		return errno.MysqlErr
	}
	return nil
}

func (s *MessageService) PopMessage(request *messages.PopMessageRequest) (*messages.PopMessageResponseData, error) {
	if data, err := db.PopMessage(request.Uid); err != nil {
		return nil, errno.MysqlErr
	} else {
		return &messages.PopMessageResponseData{
			Items: *convert.DBRespToKitexGen(data),
		}, nil
	}
}
