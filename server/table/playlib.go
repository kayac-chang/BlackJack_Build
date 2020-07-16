package table

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/conf"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/player"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/poker"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/action"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/command"
	"go.uber.org/zap"
)

func (t *Table) ask(p *player.Player, data protocol.Data, def string) string {
	if data.Expire < 1 { // 0, -1, ...
		data.Expire = ResponseTime
	}

	exp := time.Duration(data.Expire+1) * time.Second
	ctx, quit := context.WithTimeout(context.Background(), exp)
	defer quit()
	go p.CommandTx(ctx, command.Ask, code.OK, data)

	select {
	case <-ctx.Done():
		p.CommandTimeout(time.Second, command.ActionResult, protocol.Timeout, nil)
		return def
	case act := <-t.act:
		if data.Options[act.Action] {
			p.CommandTimeout(time.Second, command.ActionResult, code.OK, act)
			return act.Action
		}

		p.CommandTimeout(time.Second, command.ActionResult, protocol.BadRequest, nil)
		return def
	}
}

func (t *Table) openBet(second int) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	ctx, quit := context.WithCancel(context.Background())
	defer quit()
	defer time.Sleep(time.Second) // wait for delayed request

	for ; second > 0; second-- {
		<-ticker.C

		m := map[string][]protocol.BetData{}

		t.mu.RLock()
		for i, s := range t.seats {
			if s.player != nil {
				id := s.player.ID()
				m[id] = append(m[id], protocol.BetData{No: i})
			}
		}

		data := protocol.Data{
			ID:     t.id,
			Round:  t.round.String(),
			MaxBet: t.maxBet,
			MinBet: t.minBet,
			Expire: second,
		}

		for _, p := range t.players {
			go func(data protocol.Data, p *player.Player) {
				data.Bets = m[p.ID()]
				p.CommandTx(ctx, command.PleaseBet, code.OK, data)
			}(data, p)
		}
		t.mu.RUnlock()
	}
}

func (t *Table) setBet(bs []protocol.BetData, pid string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	p, ok := t.players[pid]
	if !ok {
		return
	}

	if code.Code(t.state[0]) != command.NewRound {
		p.CommandTimeout(time.Second, command.BetResult, protocol.Timeout, bs)
		return
	}

	for _, b := range bs {
		if b.No < 0 || b.No >= len(t.seats) {
			p.CommandTimeout(time.Second, command.BetResult, protocol.BadRequest, b)
			return
		}

		s := t.seats[b.No]
		if s.player == nil || p.ID() != s.player.ID() {
			p.CommandTimeout(time.Second, command.BetResult, protocol.NotFound, b)
			return
		}

		if b.Bet > t.maxBet || b.Bet < t.minBet {
			p.CommandTimeout(time.Second, command.BetResult, protocol.BetOutOfRange, b)
			return
		}

		s.piles[0].Action = action.Bet
		s.piles[0].Bet = b.Bet
	}

	go func() {
		p.CommandTimeout(time.Second, command.BetResult, code.OK, bs)
		t.broadcast(command.UpdateSeat, t.Seats())
	}()
}

func (t *Table) updateState(state code.Code, no, i int, wait time.Duration) {
	t.setState(state, no, i)
	t.broadcast(state, protocol.TableUpdate{State: t.state})
	time.Sleep(wait)
}

func (t *Table) newRound() (protocol.TableStat, protocol.TableUpdate) {
	t.mu.Lock()

	// if shoe cards less than 2 decks, clear out history, restart shoe
	if len(t.shoe) < 2*poker.Count {
		t.history = []interface{}{}
		t.shoe = poker.Decks(6, 0).Shuffle()
		t.round.NextShoe()
	} else {
		t.round.NextRound()
	}

	t.mu.Unlock()
	t.setState(command.NewRound, 0, 0)

	rnd := t.round.String()
	t.logger.Debug("new round", zap.Int("room_id", t.id), zap.String("round", rnd))

	return t.stat(), protocol.TableUpdate{
		Round:   rnd,
		ShoeNum: len(t.shoe),
		Seats:   t.Seats(),
	}
}

func (t *Table) firstDeal() {
	var (
		count int
		ps    []protocol.Deal
	)

	t.mu.Lock()
	for i := 0; i < 2; i++ {
		for no, s := range t.seats {
			if s.hasBet() {
				count++
				c, _ := t.shoe.Draw()
				s.piles[0].hit(c)
				ps = append(ps, protocol.Deal{No: no, Card: c, Points: s.piles[0].points()})
			}
		}

		if i == 0 {
			count++
			c, _ := t.shoe.Draw()
			t.dealer.hit(c)
			ps = append(ps, protocol.Deal{No: -1, Card: c, Points: t.dealer.points()})
		}
	}
	t.mu.Unlock()

	t.broadcast(command.FirstDeal, ps)
	time.Sleep(time.Duration(count) * time.Second)
}

func (t *Table) deal(no, i int) {
	time.Sleep(400 * time.Millisecond)
	defer time.Sleep(1600 * time.Millisecond)
	var p *pile

	t.mu.Lock()

	c, _ := t.shoe.Draw()
	switch {

	// dealer
	case no == -1:
		p = &t.dealer.pile

		// switch card once if dealer bust
		if rand.Intn(4) < conf.RTPctrl && append(t.dealer.pile.Cards, c).PointsBlackjack() > 21 {
			t.shoe.Add(c)
			c, _ = t.shoe.Draw()
			t.logger.Debug("dealer card switched", zap.Int("room_id", t.id))
		}
		t.dealer.hit(c)

	// doubled pile
	case len(t.seats[no].piles) == 2 && t.seats[no].piles[1].Action == action.Double:

		// switch card once if points > 18
		if append(t.seats[no].piles[0].Cards, c).PointsBlackjack() > 18 {
			t.shoe.Add(c)
			c, _ = t.shoe.Draw()
			t.logger.Debug("player card switched", zap.Int("room_id", t.id))
		}
		fallthrough

	default:
		p = t.seats[no].piles[i]
		if p.hit(c); p.points() > 20 {
			p.stay()
		}
	}

	t.mu.Unlock()

	t.broadcast(command.UpdateDeal, protocol.Deal{
		ShoeNum: len(t.shoe),
		Card:    c,
		No:      no,
		Pile:    i,
		Cards:   p.Cards,
		Points:  p.points(),
	})
}

func (t *Table) Settle() {
	defer func() {
		stat := t.Stat()
		t.lobby.Broadcast(frame.New(command.LobbyUpdate, code.OK, stat), false)
	}()

	m := make(map[*player.Player][]protocol.BetData)
	t.mu.Lock()
	defer t.mu.Unlock()

	dPt := t.dealer.points()
	dBj := t.dealer.blackjack()
	for no, s := range t.seats {
		if !s.hasBet() {
			continue
		}
		s.settle(dPt, dBj)

		act := make(map[string]protocol.Pile)
		bet := s.piles[0]
		if bet.Result == protocol.Paid || bet.Result == protocol.GaveUp {
			s := action.Pay
			if bet.Result == protocol.GaveUp {
				s = action.GiveUp
			}

			act[s] = protocol.Pile{
				Action: s,
				Bet:    bet.Bet,
				Cards:  bet.Cards,
				Pay:    bet.Pay,
				Result: bet.Result,
			}
			bet.Pay = bet.Bet
		}
		act[action.Bet] = protocol.Pile(*bet)

		if s.insurance != 0 {
			act[action.Insurance] = protocol.Pile{
				Bet: s.insurance,
				Pay: s.pay,
			}
		}
		if p := s.piles[1]; p.Bet != 0 { // double or split
			act[p.Action] = protocol.Pile(*p)
		}

		m[s.player] = append(m[s.player], protocol.BetData{
			ID:     t.id,
			Round:  t.round.String(),
			Dealer: t.dealer.Cards,
			No:     no,
			Action: act,
		})
	}

	// push game result
	for p, data := range m {
		p.CommandTimeout(time.Second, command.GameResult, code.OK, data)
	}

	// push history
	if dBj {
		t.history = append([]interface{}{"BJ"}, t.history...)
	} else {
		t.history = append([]interface{}{dPt}, t.history...)
	}
}

const (
	insertHistories = `
		INSERT INTO histories (room_id, end_at, round_code, dealer_cards)
		VALUES (?, ?, ?, ?)`

	insertResults = `
		INSERT INTO seat_results (
			history_id, seat_no,
			pile_0_bet, pile_0_pay, pile_0_cards,
			pile_1_bet, pile_1_pay, pile_1_cards,
			insurance, ins_pay
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`  // 10 values; use sql.Null<Value> if null

	updateLast = `
		UPDATE rooms SET last_round = ? WHERE id = ?;`
)

func (t *Table) Save() error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	t.mu.RLock()
	defer t.mu.RUnlock()

	b, err := json.Marshal(t.dealer.Cards)
	if err != nil {
		tx.Rollback()
		return err
	}

	r, err := tx.Exec(insertHistories, t.id, time.Now(), t.round.String(), b)
	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := r.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	for no, s := range t.seats {
		if !s.hasBet() {
			continue
		}

		args := []interface{}{id, no}
		for _, p := range s.piles {
			cards, _ := json.Marshal(p.Cards)
			args = append(args, p.Bet, p.Pay, cards)
		}
		args = append(args, s.insurance, s.pay)

		_, err = tx.Exec(insertResults, args...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	_, err = tx.Exec(updateLast, t.round.String(), t.id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
