package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"

	uuid "github.com/satori/go.uuid"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/controller/protoc"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/action"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/command"
)

func ActionCheck(c *client, req *Frame) (bool, error) {
	ok := false

	user, errproto, err := GetUser(c.Token)
	if errproto != nil {
		log.Println(errproto)
		c.write(NewS2CErrorAck(ECLoginError, fmt.Errorf("%d : %s", errproto.GetCode(), errproto.GetMessage())))
		return ok, err
	} else if err != nil {
		log.Println(err)
		c.write(NewS2CErrorAck(ECLoginError, err))
		return ok, err
	}

	switch req.Command {
	case CMDc2sLogin:
		ok = true
		id := uuid.Must(uuid.NewV4(), err).String()
		// fmt.Println("ActionCheck Login 1")
		// TODO
		// res, err := operate.Login(&operate.AuthCond{
		// 	Token:  req.Token,
		// 	GameID: conf.GameID,
		// 	ConnID: id,
		// })
		// if err != nil {
		// 	log.Println(err)
		// 	c.write(NewS2CErrorAck(ECLoginError, err))
		// 	return ok, err
		// }

		c.Token = req.Token
		// c.GameToken = res.GameToken
		c.ConnID = id

		c.write(NewS2CLoginAck())
		c.Balance = float64(user.UserGameInfo.GetMoney())
		c.write(NewS2CMemberInfo(user.UserGameInfo.Name, c.Balance))
		return ok, nil

	case command.Bet:
		// fmt.Println("ActionCheck Bet 1")
		obj := &protocol.Data{}
		if err := json.Unmarshal(*req.Data.(*json.RawMessage), obj); err != nil {
			log.Println(err)
			return ok, err
		}

		var betAmount float64
		for _, item := range obj.Bets {
			betAmount += item.Bet
		}
		if betAmount > c.Balance {
			return ok, errors.New("Balance error")
		}

		var syncWait sync.WaitGroup
		//var ok bool
		var err error
		// TODO 新注單
		for _, item := range obj.Bets {
			syncWait.Add(1)
			go func(goItem protocol.BetData) {
				defer syncWait.Done()
				order, errproto, apierr := NewOrder(user.UserServerInfo.Token, user.UserGameInfo.IDStr, int64(goItem.Bet))

				if errproto != nil {
					err = errors.New(fmt.Sprint("%d : %s", errproto.GetCode(), errproto.GetMessage()))
					return
				} else if err != nil {
					err = apierr
					return
				}

				c.betOrderRes[fmt.Sprintf("%d-%s", goItem.No, action.Bet)] = order
				c.BetAmount[fmt.Sprintf("%d-%s", goItem.No, action.Bet)] = goItem.Bet
				user.UserGameInfo.SumMoney(int64(goItem.Bet * -1))
			}(item)
		}
		syncWait.Wait()

		c.Balance = float64(user.UserGameInfo.GetMoney())
		c.write(NewS2CMemberInfo(user.UserGameInfo.Name, c.Balance))
		return ok, nil

	case command.Action:
		// fmt.Println("ActionCheck Action 1")
		obj := &protocol.Data{}
		if err := json.Unmarshal(*req.Data.(*json.RawMessage), obj); err != nil {
			log.Println(err)
			c.write(NewS2CErrorAck(ECServerError, err))
			return ok, err
		}

		iorder := c.betOrderRes[fmt.Sprintf("%d-%s", obj.No, action.Bet)]
		order := iorder.(*protoc.Order)
		switch obj.Action {
		case action.Split:
			// TODO 新增子單
			// fmt.Println("ActionCheck Action 1 Split 1")
			subOrder, errproto, err := NewSubOrder(user.UserServerInfo.Token, order, int64(order.Bet))
			if errproto != nil {
				return ok, errors.New(fmt.Sprint("%d : %s", errproto.GetCode(), errproto.GetMessage()))
			} else if err != nil {
				return ok, err
			}
			c.betOrderRes[fmt.Sprintf("%d-%s", obj.No, obj.Action)] = subOrder
		case action.Double:
			// TODO 新增子單
			// fmt.Println("ActionCheck Action 1 Double 1")
			subOrder, errproto, err := NewSubOrder(user.UserServerInfo.Token, order, int64(order.Bet))
			if errproto != nil {
				return ok, errors.New(fmt.Sprint("%d : %s", errproto.GetCode(), errproto.GetMessage()))
			} else if err != nil {
				return ok, err
			}

			user.UserGameInfo.SumMoney(int64(order.Bet) * -1)
			c.betOrderRes[fmt.Sprintf("%d-%s", obj.No, obj.Action)] = subOrder
		case action.Insurance:
			// TODO 新增子單
			// fmt.Println("ActionCheck Action 1 Insurance 1")
			subOrder, errproto, err := NewSubOrder(user.UserServerInfo.Token, order, int64(order.Bet/2))
			if errproto != nil {
				return ok, errors.New(fmt.Sprint("%d : %s", errproto.GetCode(), errproto.GetMessage()))
			} else if err != nil {
				return ok, err
			}
			user.UserGameInfo.SumMoney(int64(order.Bet/2) * -1)
			c.betOrderRes[fmt.Sprintf("%d-%s", obj.No, obj.Action)] = subOrder
		case action.Pay:
			// TODO 新增子單
			// fmt.Println("ActionCheck Action 1 Pay 1")
			subOrder, errproto, err := NewSubOrder(user.UserServerInfo.Token, order, 0)
			if errproto != nil {
				return ok, errors.New(fmt.Sprint("%d : %s", errproto.GetCode(), errproto.GetMessage()))
			} else if err != nil {
				return ok, err
			}
			user.UserGameInfo.SumMoney(int64(order.Bet * 2))
			c.betOrderRes[fmt.Sprintf("%d-%s", obj.No, obj.Action)] = subOrder
		case action.GiveUp: // 投降
			// TODO 新增子單
			// fmt.Println("ActionCheck Action 1 GiveUp 1")
			subOrder, errproto, err := NewSubOrder(user.UserServerInfo.Token, order, 0)
			if errproto != nil {
				return ok, errors.New(fmt.Sprint("%d : %s", errproto.GetCode(), errproto.GetMessage()))
			} else if err != nil {
				return ok, err
			}
			c.betOrderRes[fmt.Sprintf("%d-%s", obj.No, obj.Action)] = subOrder
		}
		c.Balance = float64(user.UserGameInfo.GetMoney())
		c.write(NewS2CMemberInfo(user.UserGameInfo.Name, c.Balance))
	}
	return ok, nil
}

func SendMemberInfo(c *client) {
	// fmt.Println("ActionCheck GetMemberInfo 1")
	// TODO 取得玩家資料
	user, errproto, err := GetUser(c.Token)
	if errproto != nil {
		log.Println(errproto)
		c.write(NewS2CErrorAck(ECWalletTransferError, fmt.Errorf("%d : %s", errproto.GetCode(), errproto.GetMessage())))
		return
	} else if err != nil {
		log.Println(err)
		c.write(NewS2CErrorAck(ECWalletTransferError, err))
		return
	}
	c.Balance = float64(user.UserGameInfo.GetMoney())
	c.write(NewS2CMemberInfo(user.UserGameInfo.Name, c.Balance))
}
