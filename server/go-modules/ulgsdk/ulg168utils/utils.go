package ulg168utils

import (
	"net/url"
	"time"
)

const (
	DefaultTimeout = 20 * time.Second
)

const (
// prefix              = "/v1/apis/ulg168"
// AuthPath            = prefix + "/auth"
// AuthValidatePath    = prefix + "/auth/validate"
// ExchangePath        = prefix + "/exchange"
// CheckoutPath        = prefix + "/checkout"
// WalletPath          = prefix + "/wallet/balance"
// SettlementPath      = prefix + "/game/:game_id/settlement"
// MaintainPath        = prefix + "/game/:game_id/maintain"
// MaintainDonePath    = prefix + "/game/:game_id/maintain/done"
// ULG168LoginPath     = prefix + "/login"
// MemberInfoPath      = prefix + "/userinfo"
// SearchOrderPath     = prefix + "/orders"
// GetOrderItemPath    = prefix + "/order/:uuid/item"
// ListOrderPath       = prefix + "/orders/list"
// CreateOrderPath     = prefix + "/orderdetail"
// CreateOrderItemPath = prefix + "/order/:uuid/order_item"
// PayoutOrderPath     = prefix + "/order/:uuid/payout"
// PayoutOrderItemPath = prefix + "/order/:uuid/item/:item_uuid/payout"
// AdminValidatePath   = prefix + "/admin/validate"
)

type Acion struct {
	Host string
	Path string
	Cond
}

func (h *Acion) URL() string {
	return h.Host + h.Path
}

func (h *Acion) RequestUrl() string {
	if h.Cond == nil {
		return h.Host + h.Path
	}
	return h.Host + h.Path + "?" + h.Cond.Body().Encode()
}

type Cond interface {
	Body() url.Values
}

type WsObj struct {
	CMD int `json:"cmd"`
}
