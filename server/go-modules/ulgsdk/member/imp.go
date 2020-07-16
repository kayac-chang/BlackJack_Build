package member

// func GetBalance(cond *WalletCond) (*Wallet, error) {
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: ulg168utils.WalletPath,
// 		Cond: cond,
// 	}
// 	res, err := httprequest.Get(h.RequestUrl(), ulg168utils.DefaultTimeout)
// 	if err != nil {
// 		return nil, err
// 	}
// 	obj := &Wallet{}
// 	if err := json.Unmarshal(res, obj); err != nil {
// 		return nil, err
// 	}

// 	return obj, nil
// }
