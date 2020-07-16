package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/gorilla/websocket"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/conf"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
)

const (
	CMDc2sLogin          code.Code = 8003150
	CMDc2sWalletTransfer code.Code = 8003151
	CMDc2sLogout         code.Code = 8003152

	CMDc2sPingPong code.Code = 8003199
)

const (
	CMDs2cLoginAck          code.Code = 8003050
	CMDs2cWalletTransferAck code.Code = 8003051
	CMDs2cLogoutAck         code.Code = 8003052
	CMDs2cGameNotFinished   code.Code = 8003053

	CMDs2cMemberInfo    code.Code = 8003095
	CMDs2cMaintainKick  code.Code = 8003096
	CMDs2cMaintainNoify code.Code = 8003097

	CMDs2cLoginRepeat code.Code = 8003098
	CMDs2cPingPongAck code.Code = 8003099
)

type C2SWsObj struct {
	CMD       code.Code `json:"cmd"`
	Token     string    `json:"token"`
	GameID    string    `json:"game_id"`
	GameToken string    `json:"game_token"`
}

type WsObj struct {
	CMD code.Code `json:"cmd"`
}

type S2CErrorAck struct {
	WsObj
	Data ErrorAck `json:"data"`
}

type ErrorAck struct {
	ErrCode    string `json:"err_code"`
	ErrMessage string `json:"err_message"`
}

const (
	CMDs2cErrorAck code.Code = 8003091

	ECServerError         string = "999-999"
	ECPingPongError       string = "199-001"
	ECLoginError          string = "150-001"
	ECWalletTransferError string = "151-001"
	ECLogoutError         string = "152-001"
)

func NewS2CErrorAck(code string, err error) frame.Frame {
	return frame.Frame{
		Command: CMDs2cErrorAck,
		Data: ErrorAck{
			ErrCode:    code,
			ErrMessage: err.Error(),
		},
	}
}

func GetCMDFromMsg(msg []byte) (*C2SWsObj, error) {
	obj := &C2SWsObj{}
	if err := json.Unmarshal(msg, obj); err != nil {
		log.Println(err)
		return nil, err
	}
	return obj, nil
}

func GetFrameFromMsg(msg []byte) (*frame.Frame, error) {
	obj := &frame.Frame{Data: &json.RawMessage{}}
	if err := json.Unmarshal(msg, obj); err != nil {
		log.Println(err)
		return nil, err
	}
	return obj, nil
}

func WsFormatEncode(i interface{}) []byte {
	obj, err := json.Marshal(i)
	if err != nil {
		return []byte(err.Error())
	}
	if !conf.Base64Enable {
		return []byte(obj)
	}
	return []byte(base64.StdEncoding.EncodeToString(obj))
}

func WsFormatDecode(msg []byte) ([]byte, error) {
	if !conf.Base64Enable {
		return msg, nil
	}
	b, err := base64.StdEncoding.DecodeString(string(msg))
	if err != nil {
		fmt.Println("err: ", err)
		return []byte{}, err
	}
	return b, nil
}

func isClosed(err error) bool {
	if err == nil {
		return false
	}

	_, o := err.(*net.OpError)
	_, c := err.(*websocket.CloseError)

	return o || c || err == websocket.ErrCloseSent
}
