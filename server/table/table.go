package table

import (
	"context"
	"database/sql"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/lobby"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/player"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/round"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/poker"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/command"
	"go.uber.org/zap"
)

const (
	// ResponseTime is the default limit in seconds a player should answer.
	ResponseTime = 10

	// BetTime is the default limit in seconds a player can bet.
	BetTime = 20
)

var _ lobby.Room = &Table{}

type Table struct {
	id      int
	round   *round.Round
	maxBet  float64
	minBet  float64
	state   []int
	shoe    poker.Set
	dealer  *dealerSeat
	seats   []*seat
	players map[string]*player.Player
	history []interface{}

	ch     chan frame.Frame
	act    chan protocol.Data
	mu     sync.RWMutex
	stop   bool
	db     *sql.DB
	logger *zap.Logger
	lobby  *lobby.Lobby
}

func (t *Table) Broadcast(f frame.Frame) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var wg sync.WaitGroup
	ctx, q := context.WithTimeout(context.Background(), 2*time.Second)
	defer q()

	for _, p := range t.players {
		wg.Add(1)
		go func(p *player.Player) {
			if err := p.SendTx(ctx, f); err != nil {
				t.logger.Debug("broadcast failed",
					zap.Int("room_id", t.id),
					zap.String("player_id", p.ID()),
					zap.Error(err))
			}
			wg.Done()
		}(p)
	}

	wg.Wait()
}

func (t *Table) broadcast(cmd code.Code, data interface{}) {
	t.Broadcast(frame.New(cmd, code.OK, data))
}

func (t *Table) setState(state code.Code, no, pile int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.state = []int{int(state)}
	if state == command.GameStarted {
		t.state = append(t.state, no, pile)
	}

	t.logger.Debug("state changed",
		zap.Int("room_id", t.id),
		zap.Int("state", t.state[0]),
		zap.Ints("turn", []int{no, pile}))
}

func (t *Table) clean() {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, s := range t.seats {
		if s.quited && s.player != nil {
			id := s.player.ID()
			delete(t.players, id)
			t.logger.Info("player removed", zap.Int("room_id", t.id), zap.String("player_id", id))
		}
		s.clean()
	}

	t.dealer.clean()
}

func (t *Table) Wellcome(p *player.Player) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	id := p.ID()
	if _, ok := t.players[id]; ok {
		msg := "player already exist in table" + strconv.Itoa(t.id)
		return frame.Error(command.TableResult, code.AlreadyExists, msg)
	}

	t.players[id] = p
	p.Output(t.ch)

	// update table status to p
	go func() { p.CommandTimeout(time.Second, command.TableResult, code.OK, t.Detail()) }()
	return nil
}

func (t *Table) closed() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.stop
}

func (t *Table) Close(soft bool, reason interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	// ...
	t.stop = true
}

func (t *Table) Kick(id string, c code.Code, reason interface{}) *player.Player {
	var seatChanged bool
	defer func() {
		if seatChanged {
			f := frame.New(command.LobbyUpdate, code.OK, t.Stat())
			t.lobby.Broadcast(f, false)
			t.broadcast(command.UpdateSeat, t.Seats())
		}
	}()

	t.mu.Lock()
	defer t.mu.Unlock()

	p, ok := t.players[id]
	if !ok {
		return nil
	}

	rm := true
	for _, s := range t.seats {
		if s.player != nil && s.player.ID() == id {
			seatChanged = s.quit() || seatChanged
			if s.quited { // if player has game not finished, don't remove
				rm = false
			}
		}
	}

	if rm {
		delete(t.players, id)
		t.logger.Info("player removed", zap.Int("room_id", t.id), zap.String("player_id", id))
	}

	return p
}

func (t *Table) changeSeat(id string, no int, cmd code.Code) {
	if no < 0 || no >= len(t.seats) {
		// should not happen, ignore
		return
	}

	var c code.Code
	defer func() {
		if c == code.OK {
			f := frame.New(command.LobbyUpdate, code.OK, t.Stat())
			t.lobby.Broadcast(f, false)
			t.broadcast(command.UpdateSeat, t.Seats())
		}
	}()

	t.mu.Lock()
	defer t.mu.Unlock()

	p, ok := t.players[id]
	if !ok {
		return
	}
	s := t.seats[no]

	// case have seat
	if cmd == command.TakeSeat {
		if s.player != nil {
			c = protocol.SeatOccupied
		} else {
			s.player = t.players[id]
		}
		goto result
	}

	// case leave seat
	if s.player == nil || s.player.ID() != id {
		c = protocol.BadRequest
		goto result
	}
	if s.quit() {
		c = code.OK
	}

result:
	p.CommandTimeout(time.Second, command.SeatResult, c, &protocol.Data{ID: t.id, No: no})
}

func (t *Table) Open() {
	t.logger.Info("table opened", zap.Int("room_id", t.id))
	go t.Play()
	defer t.logger.Info("table closed", zap.Int("room_id", t.id))

	for {
		req, ok := <-t.ch
		if !ok {
			return
		}

		t.logger.Debug("frame received",
			zap.Int("room_id", t.id),
			zap.Int32("command", int32(req.Command)),
			zap.String("from", req.From))

		switch req.Command {
		case command.BackToLobby:
			if p := t.Kick(req.From, code.OK, nil); p != nil {
				t.lobby.Wellcome(p)
			}

		case command.Exit, code.PlayerExited:
			t.Kick(req.From, code.PlayerExited, nil)
		}

		data, ok := req.Data.(protocol.Data)
		if !ok {
			continue
		}

		switch req.Command {

		case command.TakeSeat, command.LeaveSeat:
			t.changeSeat(req.From, data.No, req.Command)

		case command.Bet:
			if data.ID != t.id {
				continue
			}
			t.setBet(data.Bets, req.From)

		case command.Action:
			go func() { t.act <- data }()
		}
	}
}

func New(opts Options) *Table {
	t := &Table{
		id:      opts.ID,
		round:   round.New(opts.LastShoe, 1),
		maxBet:  opts.MaxBet,
		minBet:  opts.MinBet,
		dealer:  &dealerSeat{*newPile(), nil, false},
		players: make(map[string]*player.Player),
		ch:      make(chan frame.Frame),
		act:     make(chan protocol.Data),
		db:      opts.Database,
		logger:  opts.Logger,
		lobby:   opts.Lobby,
	}

	for i := 0; i < opts.SeatsNum; i++ {
		t.seats = append(t.seats, newSeat())
	}

	return t
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
