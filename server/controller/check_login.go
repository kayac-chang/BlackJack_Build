package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/ulgsdk/member"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/ulgsdk/operate"
)

func checklogin(conn *websocket.Conn) (*client, error) {
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, WsFormatEncode(NewS2CErrorAck(ECServerError, err)))
			conn.Close()
			return nil, err
		}

		msg, err := WsFormatDecode(data)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, WsFormatEncode(NewS2CErrorAck(ECServerError, err)))
			conn.Close()
			return nil, err
		}

		wsObj, err := GetCMDFromMsg(msg)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, WsFormatEncode(NewS2CErrorAck(ECServerError, err)))
			conn.Close()
			return nil, err
		}

		switch wsObj.CMD {
		case CMDc2sLogin:
			obj := &operate.C2SLogin{}
			if err := json.Unmarshal(msg, obj); err != nil {
				log.Println(err)
				conn.WriteMessage(websocket.TextMessage, WsFormatEncode(NewS2CErrorAck(ECLoginError, err)))
				continue
			}

			id := uuid.Must(uuid.NewV4(), err).String()
			// fmt.Println("ActionCheck checklogin 1 Login 1")
			// TODO 第一次進入點
			// res, err := operate.Login(&operate.AuthCond{
			// 	Token:  obj.Token,
			// 	GameID: conf.GameID,
			// 	ConnID: id,
			// })
			// if err != nil {
			// 	log.Println(err)
			// 	conn.WriteMessage(websocket.TextMessage,
			// 		WsFormatEncode(NewS2CErrorAck(ECLoginError, err)))
			// 	continue
			// }
			// w, err := member.GetBalance(&member.WalletCond{
			// 	Token:     obj.Token,
			// 	GameID:    conf.GameID,
			// 	GameToken: res.GameToken,
			// })
			// if err != nil {
			// 	log.Println(err)
			// 	conn.WriteMessage(websocket.TextMessage,
			// 		WsFormatEncode(NewS2CErrorAck(ECLoginError, err)))
			// 	continue
			// }

			// 新 api 接口
			user, errproto, err := GetUser(obj.Token)
			if errproto != nil {
				log.Println(errproto)
				conn.WriteMessage(websocket.TextMessage,
					WsFormatEncode(NewS2CErrorAck(ECLoginError, fmt.Errorf("%d : %s", errproto.GetCode(), errproto.GetMessage()))))
				continue
			} else if err != nil {
				log.Println(err)
				conn.WriteMessage(websocket.TextMessage,
					WsFormatEncode(NewS2CErrorAck(ECLoginError, err)))
				continue
			}

			conn.WriteMessage(websocket.TextMessage, WsFormatEncode(operate.NewS2CLoginAck(&operate.Auth{})))
			conn.WriteMessage(websocket.TextMessage, WsFormatEncode(member.NewS2CMemberInfo(user.UserGameInfo.Name, float64(user.UserGameInfo.GetMoney()))))

			c := newClient(conn)
			c.Token = obj.Token
			// c.GameToken = user.UserGameInfo.GameToken
			c.ConnID = id
			c.Name = user.UserGameInfo.Name
			c.Account = user.UserGameInfo.IDStr
			return c, nil
		}
	}
}
