package table

import (
	"time"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/poker"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/action"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/command"
	"go.uber.org/zap"
)

func (t *Table) Play() {
	t.logger.Info("table start playing", zap.Int("room_id", t.id))
	defer t.logger.Info("table stop playing", zap.Int("room_id", t.id))

	for {
		t.clean()

		// stop game if closed
		if t.closed() {
			// kick out all players in seat
			// ...
			return
		}

		// new round and broadcast, wait for 0.8 second
		stat, upd := t.newRound()
		t.broadcast(command.NewRound, protocol.Table{stat, upd})
		t.lobby.Broadcast(frame.New(command.LobbyUpdate, code.OK, stat), false)
		time.Sleep(800 * time.Millisecond)

		// open for bet
		t.openBet(BetTime)

		// stop betting, wait 2 seconds
		t.updateState(command.GameStarted, -1, -1, 2*time.Second)

		// deal cards
		t.firstDeal()

		dA := t.dealer.hasAce()  // dealer has Ace
		dT := t.dealer.hasTens() // dealer has 10/J/Q/K

	SeatLoop:
		for no, s := range t.seats {
			if !s.hasBet() {
				continue
			}

			// next one if player is BJ, dealer doesn't have A/10/J/Q/K
			bj := s.piles[0].blackjack()
			if bj && !(dA || dT) {
				continue
			}

			pt := s.piles[0].points()             // player points
			pP := s.piles[0].pair()               // player got Pair
			pA := s.piles[0].hasAce()             // player has Ace
			pD := pt == 9 || pt == 10 || pt == 11 // player got 9, 10 or 11 points
			fA := true                            // first action

			// update whose turn
			t.setState(command.GameStarted, no, 0)
			t.broadcast(command.UpdateTurn, protocol.Data{ID: t.id, No: no})

			for s.piles[0].Result == 0 || fA {
				opts := action.Options{
					action.Pay:       bj && (dA || dT), // player got BJ, dealer has A/10/J/Q/K, pay first?
					action.Insurance: dA && fA,         // dealer has Ace, buy insurance?
					action.Double:    pD,               // double?
					action.Split:     pP,               // player got Pair, split?
					action.Hit:       pt < 21,
					action.Stay:      true,
					action.GiveUp:    fA && !bj, // player didn't get BJ, give up in first action?
				}
				act := t.ask(s.player, protocol.Data{
					No:      no,
					Expire:  ResponseTime,
					Options: opts,
				}, action.Stay)
				fA = false // insurance and giving up is only available in fist action

				switch act {
				case action.Pay:
					s.piles[0].Result = protocol.Paid
					s.piles[0].Pay = 2 * s.piles[0].Bet

					// player BJ, next one
					continue SeatLoop

				case action.Insurance:
					s.insurance = s.piles[0].Bet / 2
					t.broadcast(command.UpdateAction, protocol.Data{
						ID:       t.id,
						No:       no,
						TotalBet: s.piles[0].Bet + s.piles[1].Bet,
						Action:   act,
						Result: []protocol.Pile{
							protocol.Pile(*s.piles[0]),
							protocol.Pile(*s.piles[1]),
						},
						Insurance: s.insurance,
					})
					if bj {
						continue SeatLoop
					}

				case action.Double:
					s.piles[1].Bet = s.piles[0].Bet
					s.piles[1].Action = action.Double
					t.broadcast(command.UpdateAction, protocol.Data{
						ID:       t.id,
						No:       no,
						TotalBet: s.piles[0].Bet + s.piles[1].Bet,
						Action:   act,
						Result: []protocol.Pile{
							protocol.Pile(*s.piles[0]),
							protocol.Pile(*s.piles[1]),
						}})
					time.Sleep(time.Second / 2)

					t.deal(no, 0)
					s.piles[0].stay()

					// player doubled, next one
					continue SeatLoop

				case action.Split:
					c, _ := s.piles[0].Cards.Draw()
					s.piles[1] = &pile{
						Cards:  poker.Set{c},
						Bet:    s.piles[0].Bet,
						Action: action.Split,
					}
					t.broadcast(command.UpdateAction, protocol.Data{
						ID:       t.id,
						No:       no,
						TotalBet: s.piles[0].Bet + s.piles[1].Bet,
						Action:   act,
						Result: []protocol.Pile{
							protocol.Pile(*s.piles[0]),
							protocol.Pile(*s.piles[1]),
						}})
					time.Sleep(time.Second / 2)

					t.deal(no, 0)
					t.deal(no, 1)

					// if it was a pair of Aces, next one
					if pA && pP {
						s.piles[0].stay()
						s.piles[1].stay()
						continue SeatLoop
					}

					pD = false
					pP = false

				case action.Hit:
					t.deal(no, 0)
					pD = false
					pP = false

				case action.Stay:
					s.piles[0].stay()

				case action.GiveUp:
					s.piles[0].Result = protocol.GaveUp
					s.piles[0].Pay = s.piles[0].Bet / 2
					continue SeatLoop
				}
			}

			if s.piles[1].Bet == 0 {
				continue
			}

			// if split:
			t.setState(command.GameStarted, no, 1)
			t.broadcast(command.UpdateTurn, protocol.Data{ID: t.id, No: no, Pile: 1})

			for s.piles[1].Result == 0 {
				act := t.ask(s.player, protocol.Data{
					No:     no,
					Pile:   1,
					Expire: ResponseTime,
					Options: action.Options{
						action.Pay:       false,
						action.Insurance: false,
						action.Double:    false,
						action.Split:     false,
						action.Hit:       true,
						action.Stay:      true,
						action.GiveUp:    false,
					},
				}, action.Stay)

				if act == action.Hit {
					t.deal(no, 1)
					continue
				}

				s.piles[1].stay()
			}
		}

		// dealer's turn
		time.Sleep(time.Second / 2)
		t.broadcast(command.UpdateTurn, protocol.Data{ID: t.id, No: -1})
		t.deal(-1, 0)

		if t.dealer.blackjack() {
			goto settlement
		}

		for pts := t.dealer.points(); pts < 17; pts = t.dealer.points() {
			t.deal(-1, 0)
		}

	settlement:
		t.Settle()
		if err := t.Save(); err != nil {
			t.logger.Error("could not save result", zap.Int("room_id", t.id), zap.Error(err))
		}

		t.setState(command.GameSettled, 0, 0)
		t.broadcast(command.GameSettled, protocol.TableUpdate{Seats: t.Seats()})
		time.Sleep(3600 * time.Millisecond) // wait 3.6 seconds
	}
}
