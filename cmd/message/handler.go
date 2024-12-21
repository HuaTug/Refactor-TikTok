package main

import (
	"context"

	"HuaTug.com/cmd/message/service"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/messages"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type MessageServiceImpl struct{}

// InsertMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) InsertMessage(ctx context.Context, request *messages.InsertMessageRequest) (resp *messages.InsertMessageResponse, err error) {
	// TODO: Your code here...
	resp = new(messages.InsertMessageResponse)

	err = service.NewMessageService(ctx).PushMessage(request)
	if err != nil {
		resp.Base = &base.Status{
			Code: consts.StatusBadRequest,
			Msg:  "Fail to insert message",
		}
		return resp, err
	}

	resp.Base = &base.Status{
		Code: consts.StatusOK,
		Msg:  "Success to insert message",
	}
	return resp, nil
}

// PopMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) PopMessage(ctx context.Context, request *messages.PopMessageRequest) (resp *messages.PopMessageResponse, err error) {
	// TODO: Your code here...
	resp = new(messages.PopMessageResponse)
	resp.Data = &messages.PopMessageResponseData{}

	data, err := service.NewMessageService(ctx).PopMessage(request)
	if err != nil {
		resp.Base = &base.Status{
			Code: consts.StatusBadRequest,
			Msg:  "Fail to pop message",
		}
		return resp, err
	}

	resp.Base = &base.Status{
		Code: consts.StatusOK,
		Msg:  "Success to pop message",
	}
	resp.Data = data
	return resp, nil
}
