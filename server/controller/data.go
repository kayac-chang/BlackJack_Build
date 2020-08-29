package controller

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/ulgsdk/operate"
)

type authRequest struct {
	Token      string `json:"login_token"`
	GameID     string `json:"game_id"`
	GameToken  string `json:"game_token,omitempty"`
	CoinType   string `json:"coin_type,omitempty"`
	CoinAmount int    `json:"coin_amount,omitempty"`
}

type quota struct {
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}

type gameInfo struct {
	Type   string  `json:"type"`
	Status int     `json:"status"`
	Rate   float64 `json:"rate"`
	Sort   int     `json:"sort"`
}

type authResult struct {
	Result  int        `json:"result"`
	Status  int        `json:"status"`
	UID     int        `json:"user_id"`
	Name    string     `json:"user_name"`
	Account string     `json:"account_name"`
	Token   string     `json:"game_token"`
	Quota   []quota    `json:"user_coin_quota"`
	Info    []gameInfo `json:"game_info"`
	ConnID  uuid.UUID  `json:"-"`
}

type c2sframe struct {
	Command   code.Code   `json:"cmd"`
	GameID    string      `json:"game_id"`
	Token     string      `json:"token"`
	GameToken string      `json:"game_token"`
	Data      interface{} `json:"data"`
}

type C2SLogin struct {
	C2SWsObj
}

type C2SWalletTransfer struct {
	C2SWsObj
	Data WalletTransfer `json:"data"`
}

type WalletTransfer struct {
	CoinType   int     `json:"coin_type"`
	CoinAmount float64 `json:"coin_amount"`
}

type C2SLogout struct {
	C2SWsObj
}

type S2CLoginAck struct {
	WsObj
	Data   operate.Auth `json:"data"`
	ErrMsg string       `json:"error_msg"`
}

func NewS2CLoginAck() frame.Frame { //(res *operate.Auth) frame.Frame {
	obj := frame.Frame{}
	obj.Command = CMDs2cLoginAck
	// obj.Data = *res
	return obj
}

type S2CGameNotFinished struct {
	WsObj
}

func NewS2CGameNotFinished() frame.Frame {
	obj := frame.Frame{}
	obj.Command = CMDs2cGameNotFinished
	return obj
}

type S2CWalletTransferAck struct {
	WsObj
	Data   operate.Exchange `json:"data"`
	ErrMsg string           `json:"error_msg"`
}

func NewS2CWalletTransferAck(res *operate.Exchange) frame.Frame {
	obj := frame.Frame{}
	obj.Command = CMDs2cWalletTransferAck
	obj.Data = *res
	return obj
}

type S2CLogoutAck struct {
	WsObj
	Data   operate.Checkout `json:"data"`
	ErrMsg string           `json:"error_msg"`
}

func NewS2CLogoutAck(res *operate.Checkout) frame.Frame {
	obj := frame.Frame{}
	obj.Command = CMDs2cLogoutAck
	obj.Data = *res
	return obj
}

type S2CMemberInfo struct {
	WsObj
	Data MemberInfo `json:"data"`
}

type MemberInfo struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func NewS2CMemberInfo(name string, balance float64) frame.Frame {
	obj := frame.Frame{}
	obj.Command = CMDs2cMemberInfo
	obj.Data = MemberInfo{
		Name:    name,
		Balance: balance,
	}
	fmt.Println("MemberInfo:", obj)
	return obj
}

type S2CMaintain struct {
	WsObj
	Data Maintain `json:"data"`
}

type Maintain struct {
	Time    int    `json:"time"`
	Message string `json:"message"`
}

func NewS2CMaintainNoify(t int, message string) frame.Frame {
	obj := frame.Frame{}
	obj.Command = CMDs2cMaintainNoify
	obj.Data = Maintain{
		Time:    t,
		Message: message,
	}
	return obj
}

func NewS2CMaintainKick(message string) frame.Frame {
	obj := frame.Frame{}
	obj.Command = CMDs2cMaintainKick
	obj.Data = Maintain{
		Time:    0,
		Message: message,
	}
	return obj
}

type S2CLoginRepeat struct {
	WsObj
}

func NewS2CLoginRepeat() frame.Frame {
	obj := frame.Frame{}
	obj.Command = CMDs2cLoginRepeat
	return obj
}

type S2CPingPongAck struct {
	WsObj
	Data PingPongAck `json:"data"`
}

type PingPongAck struct {
	Ping string `json:"ping"`
}

func NewS2CPingPongAck() frame.Frame {
	return frame.New(CMDs2cPingPongAck, code.OK, PingPongAck{Ping: "pong"})
}
