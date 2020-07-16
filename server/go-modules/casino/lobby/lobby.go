package lobby

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/player"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
	"go.uber.org/zap"
)

// A Lobby is a portal of a multi-room game where players are gathered.
type Lobby struct {
	mu sync.RWMutex
	in chan frame.Frame

	ctx  context.Context
	stop context.CancelFunc

	opts    Options
	rooms   map[string]Room
	players map[string]*player.Player
}

// info is a shorthand of l.opts.Logger.Info(msg, fields)
func (l *Lobby) info(msg string, fields ...zap.Field) {
	l.opts.Logger.Info(msg, fields...)
}

// debug is a shorthand of l.opts.Logger.Debug(msg, fields)
func (l *Lobby) debug(msg string, fields ...zap.Field) {
	l.opts.Logger.Debug(msg, fields...)
}

// Broadcast sends the given frame to all players in the lobby.
// If all, it asks all rooms to broadcast the message.
// It blocks until all players have been sent or a second has passed.
func (l *Lobby) Broadcast(f frame.Frame, all bool) {
	select {
	case <-l.ctx.Done():
		// lobby closed, do nothing
	default:
		l.mu.RLock()
		defer l.mu.RUnlock()

		if all {
			for _, r := range l.rooms {
				go r.Broadcast(f)
			}
		}

		var wg sync.WaitGroup
		for _, p := range l.players {
			wg.Add(1)
			go func(p *player.Player) {
				p.SendTimeout(time.Second, f)
				wg.Done()
			}(p)
		}

		wg.Wait()
	}
}

// Add appends the given room to the lobby. If the room is nil, it panics.
func (l *Lobby) Add(r Room) error {
	if r == nil {
		panic("lobby: nil Room applied")
	}

	id := r.ID()

	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.rooms[id]; ok {
		return fmt.Errorf("lobby: room %s already exist", id)
	}
	l.rooms[id] = r

	l.info("room added", zap.String("room_id", id))
	return nil
}

// Remove asks the given room to close and removes it from the lobby.
// If the room doesn't exist, it does nothing.
func (l *Lobby) Remove(room string, soft bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if r, ok := l.rooms[room]; ok {
		go r.Close(soft, nil)
		delete(l.rooms, room)
		l.info("room removed", zap.String("room_id", room))
	}
}

// wellcome sends a list of the room stats to the player.
func (l *Lobby) wellcome(p *player.Player) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	// send wellcome message, all rooms' information
	var list []interface{}
	for _, r := range l.rooms {
		list = append(list, r.Stat())
	}
	p.CommandTimeout(time.Second, l.opts.Wellcome(), code.OK, list)
}

// Wellcome puts the given player in the lobby. If a player with same
// ID already exist,
func (l *Lobby) Wellcome(p *player.Player) {
	select {
	case <-l.ctx.Done():

		// closed, service unavailable
		go p.QuitWithCommand(l.opts.Wellcome(), code.Unavailable, nil)
	default:
		l.mu.Lock()
		defer l.mu.Unlock()

		// check if duplicated
		var o *player.Player
		var r Room
		if o = l.players[p.ID()]; o != nil {
			goto dup
		}

		for _, r = range l.rooms {
			if o = r.Player(p.ID()); o != nil {
				goto dup
			}
		}

	dup:
		if o != nil {
			switch l.opts.HandleDuplicated {
			case RemoveFormer:
				if r != nil {
					r.Kick(o.ID(), code.PlayerNewerLogin, nil)
					break
				}
				go o.QuitWithCommand(l.opts.Wellcome(), code.PlayerNewerLogin, nil)
			case RejectLatter:
				msg := "player already in lobby"
				if r != nil {
					msg = "player already in room " + r.ID()
				}
				go p.QuitWithCommand(l.opts.Wellcome(), code.PlayerAlreadyLoggedIn, msg)
				return
			}
		}

		// put in the player list
		p.Output(l.in)
		l.players[p.ID()] = p
		go l.wellcome(p)

		l.info("player joined", zap.String("player_id", p.ID()))
	}
}

// move moves the player with given ID to the room.
func (l *Lobby) move(id, room string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	p := l.players[id]
	r := l.rooms[room]

	if p != nil && r != nil {
		if e, ok := r.Wellcome(p).(*frame.Err); ok {
			go p.SendTimeout(time.Second, e.Frame())
			return
		}

		delete(l.players, id)
		l.info("player moved", zap.String("player_id", p.ID()), zap.String("room_id", room))
	}
}

// kick removes the given player from the lobby, returning true if such
// player was found.
func (l *Lobby) kick(pid string, c code.Code, reason interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	p, ok := l.players[pid]
	if ok {
		go p.QuitWithCommand(l.opts.Wellcome(), c, reason)

		delete(l.players, pid)
		l.info("player removed from lobby", zap.String("player_id", pid))
	}
	return ok
}

// Kick removes the given player from the lobby, asking all the rooms
// to kick out the player.
func (l *Lobby) Kick(pid string, c code.Code, reason interface{}) {
	if l.kick(pid, c, reason) {
		return
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	for _, r := range l.rooms {
		if r.Kick(pid, c, reason) != nil {
			l.info("player removed from room", zap.String("player_id", pid), zap.String("room_id", r.ID()))
		}
	}
}

// Player returns the player and the room it belongs to by given ID.
// If such player not found, it returns nil. If the player is in the lobby,
// the returned room is nil.
func (l *Lobby) Player(pid string) (*player.Player, Room) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if p := l.players[pid]; p != nil {
		return p, nil
	}

	for _, r := range l.rooms {
		if p := r.Player(pid); p != nil {
			return p, r
		}
	}

	return nil, nil
}

// Room returns the room by given ID. If such room not found, it returns nil.
func (l *Lobby) Room(room string) Room {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.rooms[room]
}

// Open opens all the rooms in the lobby, handling the incoming frames.
// Open blocks until the lobby is closed.
func (l *Lobby) Open() {
	l.info("lobby opened")
	defer l.info("lobby closed")

	l.mu.RLock()
	for _, r := range l.rooms {
		go r.Open()
	}
	l.mu.RUnlock()

	for {
		select {
		case <-l.ctx.Done():
			return
		case f := <-l.in:
			fmt.Println("dd: ", l.opts.Move())
			l.debug("frame received", zap.Int32("command", int32(f.Command)), zap.String("player_id", f.From))
			switch f.Command {
			case l.opts.Move():
				l.move(f.From, l.opts.RoomID(f.Data))
			case l.opts.Exit(), code.PlayerExited:
				l.kick(f.From, code.PlayerExited, nil)
			}
		}
	}
}

// Close removes all the player from the lobby. Also, it calls Close(soft) of
// all rooms in the lobby. It waits until all rooms are closed.
func (l *Lobby) Close(soft bool, reason interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.stop()

	for _, p := range l.players {
		go p.QuitWithCommand(l.opts.Wellcome(), code.Unavailable, reason)
	}

	wg := sync.WaitGroup{}
	for _, r := range l.rooms {
		go func(r Room) {
			wg.Add(1)
			defer wg.Done()
			r.Close(soft, reason)
		}(r)
	}

	if soft {
		wg.Wait()
	}
}

// New returns a newly created lobby by given options. If opts is nil,
// a default Options will be applied.
func New(opts *Options) (*Lobby, error) {
	if opts == nil {
		opts = DefaultOptions()
	}

	if err := opts.Valid(); err != nil {
		return nil, err
	}

	opts.Logger = opts.Logger.WithOptions(zap.AddCallerSkip(1))
	ctx, stop := context.WithCancel(context.Background())
	return &Lobby{
		in:      make(chan frame.Frame),
		ctx:     ctx,
		stop:    stop,
		opts:    *opts,
		rooms:   make(map[string]Room),
		players: make(map[string]*player.Player),
	}, nil
}
