package operate

import "gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/ulgsdk/ulg168utils"

type WsObj struct {
	CMD int `json:"cmd"`
}

type C2SLogout struct {
	C2SWsObj
}

type S2CLoginAck struct {
	WsObj
	Data   Auth   `json:"data"`
	ErrMsg string `json:"error_msg"`
}

func NewS2CLoginAck(res *Auth) S2CLoginAck {
	obj := S2CLoginAck{}
	obj.CMD = ulg168utils.CMDs2cLoginAck
	obj.Data = *res
	return obj
}

type S2CWalletTransferAck struct {
	WsObj
	Data   Exchange `json:"data"`
	ErrMsg string   `json:"error_msg"`
}

func NewS2CWalletTransferAck(res *Exchange) S2CWalletTransferAck {
	obj := S2CWalletTransferAck{}
	obj.CMD = ulg168utils.CMDs2cWalletTransferAck
	obj.Data = *res
	return obj
}

type S2CLogoutAck struct {
	WsObj
	Data   Checkout `json:"data"`
	ErrMsg string   `json:"error_msg"`
}

func NewS2CLogoutAck(res *Checkout) S2CLogoutAck {
	obj := S2CLogoutAck{}
	obj.CMD = ulg168utils.CMDs2cLogoutAck
	obj.Data = *res
	return obj
}

type SuccessAck struct {
	Success bool `json:"success"`
}
