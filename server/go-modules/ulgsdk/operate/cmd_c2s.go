package operate

var ()

type C2SWsObj struct {
	CMD       int    `json:"cmd"`
	Token     string `json:"token"`
	GameID    string `json:"game_id"`
	GameToken string `json:"game_token"`
}

type C2SLogin struct {
	C2SWsObj
}

type C2SWalletTransfer struct {
	C2SWsObj
	Data c2sWalletTransfer `json:"data"`
}

type c2sWalletTransfer struct {
	CoinType   int     `json:"coin_type"`
	CoinAmount float64 `json:"coin_amount"`
}
