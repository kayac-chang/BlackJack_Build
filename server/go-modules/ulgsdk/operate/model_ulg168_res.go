package operate

import "gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/utils/timetool"

type AuthValidate struct {
	Status int `json:"status"`
}

type Auth struct {
	AuthStatus    int             `json:"auth_status"` //0: 已登入, 1: 重新登入
	Result        int             `json:"result"`
	Status        int             `json:"status"`
	UserID        int             `json:"user_id"`
	UserName      string          `json:"user_name"`
	AccountName   string          `json:"account_name"`
	GameToken     string          `json:"game_token"`
	UserCoinQuota []userCoinQuota `json:"user_coin_quota"`
	GameInfo      []gameInfo      `json:"game_info"`
}

type userCoinQuota struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type gameInfo struct {
	Type   string  `json:"type"`
	Status int     `json:"status"`
	Rate   float64 `json:"rate"`
	Sort   int     `json:"sort"`
}

type Exchange struct {
	Result        int             `json:"result"`
	GameCoin      float64         `json:"game_coin"`
	Status        int             `json:"status"`
	UserCoinQuota []userCoinQuota `json:"user_coin_quota"`
	GameInfo      []gameInfo      `json:"game_info"`
}

type Checkout struct {
	Result                int    `json:"result"`   //0失敗(需轉回登入頁面) 1 成功
	ErrorMsg              string `json:"errorMsg"` //result = 0, 會給錯誤訊息
	CheckoutUserCoinQuota `json:"user_coin_quota"`
	AmountCoin            `json:"amount_coin"`
}

type CheckoutUserCoinQuota struct {
	Coin1Out     *float64           `json:"coin1_out"`
	Coin2Out     *float64           `json:"coin2_out"`
	Coin3Out     *float64           `json:"coin3_out"`
	Coin4Out     *float64           `json:"coin4_out"`
	Betting      float64            `json:"betting"`
	Win          float64            `json:"win"`
	Lost         float64            `json:"lost"`
	OutboundTime timetool.Timestamp `json:"outbound_time"`
	Status       int                `json:"status"` //0 空紀錄 1 已換匯 2 結算
}

type AmountCoin struct {
	Coin1 float64 `json:"coin1"`
	Coin2 float64 `json:"coin2"`
	Coin3 float64 `json:"coin3"`
	Coin4 float64 `json:"coin4"`
}
