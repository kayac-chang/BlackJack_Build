package orderframe

import (
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/ulgsdk/order"
)

//OrderItemFill 填入OrderItem轉換成PayoutOrderItem的必要資訊
func OrderItemFill(oir *order.OrderItemRes, pRate float64, pWin float64, pResult string) {
	oir.Rate = &pRate
	oir.Win = &pWin
	oir.Result = &pResult
}

//OrderItemPayoutByOrderItemRes 使用 OrderItemRes 進行ItemPayout結算 不傳入結果
// func OrderItemPayoutByOrderItemRes(pOr *order.OrderRes, oir *order.OrderItemRes) error {
// 	poi := &order.PayoutOrderItem{
// 		Basic:    pOr.Basic,
// 		Rate:     *oir.Rate,
// 		Win:      *oir.Win,
// 		Result:   *oir.Result,
// 		PayoutAt: timetool.GetNowByUTC(),
// 	}
// 	nres, err := order.OrderItemPayout(pOr.UUID, oir.UUID, poi)

// 	if err != nil {
// 		return err
// 	}
// 	oir = nres
// 	return nil

// }

//OrderPayoutByOrderRes 使用 OrderRes 進行結算 而不是 PayoutOrder
//pOr 要結算的物件
// func OrderPayoutByOrderRes(pOr *order.OrderRes) error {
// 	nOrderPay := &order.PayoutOrder{
// 		pOr.Basic,
// 		*pOr.Result,
// 		*pOr.Summary,
// 		time.Now(),
// 	}
// 	nOrderRes, err := order.OrderPayout(pOr.UUID, nOrderPay)
// 	if err != nil {
// 		return err
// 	}
// 	pOr = nOrderRes
// 	return nil
// }

// //SupperOrderPayoutByOrderRes 輸入 OrderRes , pResult 來完成整個結算
// func SupperOrderPayoutByOrderRes(pOr *order.OrderRes) error {
// 	for _, obj := range pOr.OrderItems {
// 		err := OrderItemPayoutByOrderItemRes(pOr, obj)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	err := OrderPayoutByOrderRes(pOr)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //SupperOrderPayoutByOrderResV2 輸入 OrderRes , pResult 來完成整個結算 會初始化pOr的WIN
// func SupperOrderPayoutByOrderResV2(pOr *order.OrderRes) error {
// 	tmpwin := 0.0
// 	pOr.Win = &tmpwin
// 	for _, obj := range pOr.OrderItems {
// 		tmpwin += *obj.Win
// 		err := OrderItemPayoutByOrderItemRes(pOr, obj)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	err := OrderPayoutByOrderRes(pOr)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
