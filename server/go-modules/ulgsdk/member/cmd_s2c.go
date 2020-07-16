package member

import "gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/ulgsdk/ulg168utils"

type S2CMemberInfo struct {
	ulg168utils.WsObj
	Data MemberInfo `json:"data"`
}

type MemberInfo struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func NewS2CMemberInfo(name string, balance float64) S2CMemberInfo {
	obj := S2CMemberInfo{}
	obj.CMD = ulg168utils.CMDs2cMemberInfo
	obj.Data = MemberInfo{
		Name:    name,
		Balance: balance,
	}
	return obj
}
