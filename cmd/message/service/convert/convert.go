package convert

import (
	"HuaTug.com/cmd/message/dal/db"
	"HuaTug.com/kitex_gen/messages"
)

func DBRespToKitexGen(items *[]db.Message) *[]*messages.MessageInfo {
	result := make([]*messages.MessageInfo, 0)
	for _, item := range *items {
		result = append(result, &messages.MessageInfo{
			FromUid: item.FromUserId,
			ToUid:   item.ToUserId,
			Content: item.Content,
		})
	}
	return &result
}
