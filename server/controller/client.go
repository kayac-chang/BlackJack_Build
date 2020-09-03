package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/conf"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/controller/protoc"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/action"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/command"
)

type IOrder interface {
	GetBet() uint64
}

type client struct {
	betData     protocol.BetData
	betOrderRes map[string]IOrder
	ConnID      string
	Account     string
	Name        string
	Balance     float64
	Token       string
	GameToken   string
	BetAmount   map[string]float64
	GameRoom    string
	GameRound   string

	ctx  context.Context
	quit context.CancelFunc

	// in is the message queue which sent to WebSocket clients.
	in chan frame.Frame

	// out is a channel which the data sent to lobby.Player.
	out chan frame.Frame

	// conn holds the WebSocket client.
	conn *websocket.Conn
}

func (c *client) serve() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case f := <-c.in:
			if f.Status == code.PlayerNewerLogin {
				writeJSON(c.conn, NewS2CLoginRepeat())
				continue
			}

			switch f.Command {
			case command.Ask:
				if data, ok := f.Data.(protocol.Data); ok {
					amount := c.betOrderRes[fmt.Sprintf("%d-%s", data.No, action.Bet)].GetBet()
					if uint64(c.Balance) < amount {
						data.Options[action.Double] = false
						data.Options[action.Split] = false
					}

					if uint64(c.Balance) < amount/2 {
						data.Options[action.Insurance] = false
					}

					f.Data = data
				}

			case command.NewRound, command.TableResult:
				if t, ok := f.Data.(protocol.Table); ok {
					roundData := strings.Split(t.Round, "-")
					day := fmt.Sprintf("%s-%s-%s", roundData[0], roundData[1], roundData[2])
					round := fmt.Sprintf("%s-%s", roundData[3], roundData[4])
					r := fmt.Sprintf("%s-%03d-%s", day, t.ID, round)
					c.GameRound = r
					c.GameRoom = strconv.Itoa(t.ID)
				}

			case command.UpdateSeat:
				if seats, ok := f.Data.([]protocol.Seat); ok {
					for _, s := range seats {
						if s.Account == c.Account {
							// ...
						}
					}
				}

			case command.GameResult:
				if result, ok := f.Data.([]protocol.BetData); ok {
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

					for _, d := range result {
						// banker := strings.Join(d.Dealer.Codes(), ",")
						ibetOrder, _ := c.betOrderRes[fmt.Sprintf("%d-%s", d.No, action.Bet)]
						betOrder := ibetOrder.(*protoc.Order)

						for k, v := range d.Action {
							key := k
							if key == action.Pay || key == action.GiveUp {
								key = action.Bet
							}
							_, found := c.betOrderRes[fmt.Sprintf("%d-%s", d.No, key)]
							if !found {
								continue
							}
							switch k {
							case action.Double:
								// log.Printf("settlement: %s, %+v, %+v\n", k, v, o)
								v.Cards = d.Action[action.Bet].Cards
							case action.Insurance:
								// log.Printf("settlement: %s, %+v, %+v\n", k, v, o)
								v.Cards = d.Dealer
							}
							betOrder.Win += uint64(v.Pay)
							// settlement(basic, k, banker, v, c.betOrderRes)
						}
						user.UserGameInfo.MoneyU += betOrder.Win
						EndOrder(c.Token, betOrder)
					}
					c.betOrderRes = make(map[string]IOrder)
					c.Balance = float64(user.UserGameInfo.GetMoney())
					c.write(NewS2CMemberInfo(user.UserGameInfo.Name, c.Balance))

					var total float64
					for _, d := range result {
						for _, p := range d.Action {
							total += p.Pay - p.Bet
						}
					}

					f.Data = protocol.GameResult{
						ID:    result[0].ID,
						Round: result[0].Round,
						Win:   total,
					}
				}
			}

			if err := writeJSON(c.conn, f); err != nil {
				log.Println("controller:", err)
				if err == websocket.ErrCloseSent { // already closed
					c.Close()
				}
			}
		}
	}
}

func (c *client) listenAndServe() {
	go c.serve()
	defer func() {
		// time.Sleep(5 * time.Second)
		close(c.out)
	}()

	for {
		req := Frame{Frame: frame.Frame{Data: &json.RawMessage{}}}

		if err := readJSON(c.conn, &req); err != nil {
			if isClosed(err) {
				return
			}
			c.write(NewS2CErrorAck(ECServerError, err))
			continue
		}

		if conf.LoginRepeatEnable && req.Command != CMDc2sLogin {
			// TODO 重複登入判斷 要移除的功能
			// res, err := operate.LoginValidate(&operate.AuthValidateCond{
			// 	Token:     c.Token,
			// 	GameToken: c.GameToken,
			// 	GameID:    conf.GameID,
			// 	ConnID:    c.ConnID,
			// })
			// if res.Status != 1 {
			// 	c.write(NewS2CLoginRepeat())
			// 	continue
			// }
		}

		ok, err := ActionCheck(c, &req)
		if err != nil {
			c.write(NewS2CErrorAck(ECServerError, err))
			continue
		}

		if ok {
			continue
		}

		d := protocol.Data{}
		if err = json.Unmarshal(*req.Data.(*json.RawMessage), &d); err != nil {
			continue
		}
		req.Data = d

		c.out <- req.Frame
	}
}

func (c *client) Receive() <-chan frame.Frame {
	return c.out
}

func (c *client) Send() chan<- frame.Frame {
	return c.in
}

func (c *client) write(f frame.Frame) {
	select {
	case <-c.ctx.Done():
	case c.in <- f:
	}
}

func (c *client) Close() {
	c.quit()
}

func newClient(conn *websocket.Conn) *client {
	if conn == nil {
		panic("controller: nil websocket connection")
	}

	ctx, quit := context.WithCancel(context.Background())
	c := client{
		in:          make(chan frame.Frame, 8),
		out:         make(chan frame.Frame),
		ctx:         ctx,
		quit:        quit,
		conn:        conn,
		betOrderRes: make(map[string]IOrder),
		BetAmount:   make(map[string]float64),
	}
	go c.listenAndServe()

	return &c
}
