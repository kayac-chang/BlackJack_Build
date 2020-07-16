package order

import (
	"net/http"
)

// func OrderDetailCreate(o *CreateOrder) (*OrderRes, error) {
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: ulg168utils.CreateOrderPath,
// 		Cond: o,
// 	}
// 	res, err := httprequest.PostForm(h.URL(), ulg168utils.DefaultTimeout, h.Cond.Body())
// 	if err != nil {
// 		return nil, err
// 	}
// 	obj := &Order{}
// 	if err := json.Unmarshal(res, obj); err != nil {
// 		return nil, err
// 	}
// 	return obj.Order, nil
// }

// func OrderItemCreate(orderUUID string, oi *CreateOrderItem) (*OrderItemRes, error) {
// 	tmp := strings.Replace(ulg168utils.CreateOrderItemPath, ":uuid", orderUUID, -1)
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: tmp,
// 		Cond: oi,
// 	}
// 	res, err := httprequest.PostForm(h.URL(), ulg168utils.DefaultTimeout, h.Cond.Body())
// 	if err != nil {
// 		return nil, err
// 	}
// 	obj := &OrderItemRes{}
// 	if err := json.Unmarshal(res, obj); err != nil {
// 		return nil, err
// 	}
// 	return obj, nil
// }

// func OrderDeatailPayout(o *OrderRes) (*OrderRes, error) {
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: strings.Replace(PayoutOrderPath, ":uuid", o.UUID, -1),
// 		Cond: o,
// 	}
// 	res, err := httprequest.PostForm(h.URL(), ulg168utils.DefaultTimeout, h.Cond.Body())
// 	if err != nil {
// 		return nil, err
// 	}
// 	obj := &OrderRes{}
// 	if err := json.Unmarshal(res, obj); err != nil {
// 		return nil, err
// 	}
// 	return obj, nil
// }

// func OrderItemPayout(orderUUID, itemUUID string, oi *PayoutOrderItem) (*OrderItemRes, error) {
// 	tmp := strings.Replace(ulg168utils.PayoutOrderItemPath, ":uuid", orderUUID, -1)
// 	tmp = strings.Replace(tmp, ":item_uuid", itemUUID, -1)
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: tmp,
// 		Cond: oi,
// 	}
// 	res, err := httprequest.Put(h.URL(), ulg168utils.DefaultTimeout, h.Cond.Body())
// 	if err != nil {
// 		return nil, err
// 	}
// 	obj := &OrderItemRes{}
// 	if err := json.Unmarshal(res, obj); err != nil {
// 		return nil, err
// 	}
// 	return obj, nil
// }

// func OrderPayout(orderUUID string, o *PayoutOrder) (*OrderRes, error) {
// 	tmp := strings.Replace(ulg168utils.PayoutOrderPath, ":uuid", orderUUID, -1)
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: tmp,
// 		Cond: o,
// 	}
// 	res, err := httprequest.Put(h.URL(), ulg168utils.DefaultTimeout, h.Cond.Body())
// 	if err != nil {
// 		return nil, err
// 	}
// 	obj := &OrderRes{}
// 	if err := json.Unmarshal(res, obj); err != nil {
// 		return nil, err
// 	}
// 	return obj, nil
// }

// func LogoutCheck(cond *ListOrderCond) (bool, error) {
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: ulg168utils.ListOrderPath,
// 		Cond: cond,
// 	}
// 	res, err := httprequest.Get(h.RequestUrl(), ulg168utils.DefaultTimeout)
// 	if err != nil {
// 		return false, err
// 	}
// 	obj := []*OrderRes{}
// 	if err := json.Unmarshal(res, &obj); err != nil {
// 		return false, err
// 	}
// 	if len(obj) > 0 {
// 		return false, nil
// 	}
// 	return true, nil
// }

// func SearchOrder(c *gin.Context) {
// 	director := func(req *http.Request) {
// 		urls := strings.Split(ulg168utils.Conf.APIHost, "://")
// 		schema, host := urls[0], urls[1]
// 		req.Host = host
// 		req.Method = "GET"
// 		req.URL.Host = req.Host
// 		req.URL.Scheme = schema
// 		req.URL.Path = ulg168utils.SearchOrderPath
// 	}
// 	proxy := &httputil.ReverseProxy{Director: director}
// 	proxy.ModifyResponse = addCustomHeader
// 	proxy.ServeHTTP(c.Writer, c.Request)
// }

// func GetOrderDetail(c *gin.Context, uuid string) {
// 	path := strings.Replace(ulg168utils.GetOrderItemPath, ":uuid", uuid, -1)
// 	director := func(req *http.Request) {
// 		urls := strings.Split(ulg168utils.Conf.APIHost, "://")
// 		schema, host := urls[0], urls[1]
// 		req.Host = host
// 		req.Method = "GET"
// 		req.URL.Host = req.Host
// 		req.URL.Scheme = schema
// 		req.URL.Path = path
// 	}
// 	proxy := &httputil.ReverseProxy{Director: director}
// 	proxy.ModifyResponse = addCustomHeader
// 	proxy.ServeHTTP(c.Writer, c.Request)
// }

func addCustomHeader(r *http.Response) error {
	r.Header["Access-Control-Allow-Origin"] = []string{}
	return nil
}

//OrderItemPayoutByOrderItemRes 使用 OrderItemRes 進行ItemPayout結算
// func OrderItemPayoutByOrderItemRes(pOr *OrderRes, oir *OrderItemRes, pRate float64, pWin float64, pResult string) error {
// 	poi := &PayoutOrderItem{
// 		Basic:    pOr.Basic,
// 		Rate:     pRate,
// 		Win:      pWin,
// 		Result:   pResult,
// 		PayoutAt: timetool.GetNowByUTC(),
// 	}
// 	nres, err := OrderItemPayout(pOr.UUID, oir.UUID, poi)
// 	if err != nil {
// 		return err
// 	}
// 	oir = nres
// 	return nil
// }

// //OrderPayoutByOrderRes 使用 OrderRes 進行結算 而不是 PayoutOrder
// //pOr 要結算的物件
// func OrderPayoutByOrderRes(pOr *OrderRes) error {
// 	nOrderPay := &PayoutOrder{
// 		pOr.Basic,
// 		*pOr.Result,
// 		*pOr.Summary,
// 		timetool.GetDateByUTC(),
// 	}
// 	nOrderRes, err := OrderPayout(pOr.UUID, nOrderPay)
// 	if err != nil {
// 		return err
// 	}
// 	pOr = nOrderRes
// 	return nil
// }
