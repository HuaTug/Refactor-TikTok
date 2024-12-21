package chat

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/messages"
	"HuaTug.com/kitex_gen/users"
	jwt "HuaTug.com/pkg"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
)

type ChatService struct {
	ctx  context.Context
	c    *app.RequestContext
	conn *websocket.Conn
}

type _user struct {
	username string
	conn     *websocket.Conn
	rsa      *utils.RsaService
}

var userMap = make(map[int64]*_user)

func NewChatService(ctx context.Context, c *app.RequestContext, conn *websocket.Conn) *ChatService {
	return &ChatService{
		ctx:  ctx,
		c:    c,
		conn: conn,
	}
}

func (service ChatService) Login() error {
	v, err := jwt.ConvertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		hlog.Info(err)
		return err
	}
	uid := utils.Transfer(v)
	user, err := rpc.UserClient.GetUserInfo(service.ctx, &users.GetUserInfoRequest{
		UserId: uid,
	})
	if err != nil {
		hlog.Info(err)
		return err
	}

	r := utils.NewRsaService()
	// pub := service.c.GetHeader("rsa_public_key")
	// if err := r.Build(pub); err != nil {
	// 	hlog.Info(err)
	// 	return err
	// }
	// 生成密钥对
	priv, pub, err := r.GenerateKeyPair(2048)
	if err != nil {
		log.Fatalf("Error generating key pair: %v", err)
	}

	// 保存公钥和私钥
	_, err = r.SavePrivateKey("private_key"+fmt.Sprint(uid)+".pem", priv)
	if err != nil {
		log.Fatalf("Error saving private key: %v", err)
	}
	_, err = r.SavePublicKey("public_key"+fmt.Sprint(uid)+".pem", pub)
	if err != nil {
		log.Fatalf("Error saving public key: %v", err)
	}
	priKey, _ := os.ReadFile("private_key" + fmt.Sprint(uid) + ".pem")
	pubKey, _ := os.ReadFile("public_key" + fmt.Sprint(uid) + ".pem")
	hlog.Info(priKey, pubKey)
	r.Build(pub, priv)
	userMap[uid] = &_user{conn: service.conn, username: user.User.UserName, rsa: r}

	// 将私钥和公钥传输给客户端
	msg := "Login Success"
	if err := service.conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		hlog.Info(err)
		return errno.ServiceErr
	}
	return nil
}

func (service ChatService) Logout() error {
	v, err := jwt.ConvertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return err
	}
	uid := utils.Transfer(v)
	userMap[uid] = nil
	return nil
}

func (service ChatService) SendMessage(msg Message) error {
	v, err := jwt.ConvertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return err
	}
	fromId := utils.Transfer(v)
	t := service.c.Query("to_user_id")
	toId, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		hlog.Info(err)
		return err
	}
	user, err := rpc.UserClient.GetUserInfo(service.ctx, &users.GetUserInfoRequest{
		UserId: toId,
	})
	if err != nil {
		hlog.Info(err)
		return err
	}
	if user == nil {
		return errors.New("user not exist")
	}

	toConn := userMap[toId]
	fromConn := userMap[fromId]
	hlog.Info(toConn, ",", fromConn)
	switch toConn {
	case nil: // 用户离线，存储消息到数据库
		{
			plainText := []byte(msg.Text)
			if err := rpc.InsertMessage(context.Background(), &messages.InsertMessageRequest{
				Message: &messages.MessageInfo{
					FromUid: fmt.Sprint(fromId),
					ToUid:   fmt.Sprint(toId),
					Content: string(plainText),
				},
			}); err != nil {
				hlog.Info(err)
				return errors.New("insert message failed")
			}
		}
	default: // 用户在线，直接发送加密消息
		hlog.Info(msg.Text)
		plainText, err := fromConn.rsa.EncryptMessage(toConn.rsa.ClientRsaKey, msg.Text)
		if err != nil {
			hlog.Info(err)
			return errno.ServiceErr
		}

		ciphertext, err := toConn.rsa.DecryptMessage(toConn.rsa.ServerRsaKey, plainText)
		if err != nil {
			hlog.Info(err)
			return errno.ServiceErr
		}
		hlog.Info(ciphertext)
		err = toConn.conn.WriteMessage(websocket.TextMessage, []byte(ciphertext))
		if err != nil {
			hlog.Info(err)
			return errno.ServiceErr
		}
	}
	return nil
}

func (service ChatService) ReadOfflineMessage() error {
	v, err := jwt.ConvertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return errno.TokenInvailedErr
	}
	uid := utils.Transfer(v)
	data, err := rpc.PopMessage(context.Background(), &messages.PopMessageRequest{
		Uid: fmt.Sprint(uid),
	})
	if err != nil {
		hlog.Info(err)
		return errno.ServiceErr
	}
	//toConn := userMap[uid]
	for _, item := range (*data).Items {
		// 解密离线消息
		ciphertext := []byte(item.Content)
		// plainText, err := toConn.rsa.DecryptMessage(toConn.rsa.ServerRsaKey, ciphertext)
		// if err != nil {
		// 	hlog.Info(err)
		// 	return errno.ServiceErr
		// }

		// // 加密并发送消息
		// ciphertext, err = toConn.rsa.EncryptMessage(toConn.rsa.ClientRsaKey, userinfoAppend(plainText, item.FromUid))
		// if err != nil {
		// 	hlog.Info(err)
		// 	return errno.ServiceErr
		// }
		err = service.conn.WriteMessage(websocket.TextMessage, []byte(ciphertext))
		if err != nil {
			hlog.Info(err)
			return errno.ServiceErr
		}
	}
	return nil
}

func userinfoAppend(rawText string, from string) string {
	return (utils.ConvertTimestampToString(time.Now().Unix()) + ` [` + from + `]: ` + rawText)
}
