package player

import (
	"context"
	"sync"
	"time"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
)

// A Player represents a player in a game.
type Player struct {
	id, name string
	client   Client

	// Quit makes the player leave.
	Quit context.CancelFunc
	ctx  context.Context

	in  chan frame.Frame
	out chan<- frame.Frame
	mu  sync.RWMutex
}

// Send sends the given frame to the client.
func (p *Player) Send(f frame.Frame) {
	select {
	case <-p.ctx.Done():
		return
	case p.in <- f:
		// already sent, do nothing
	}
}

// SendTx sends the given frame to the client, returning error if the
// ctx done before sending the frame.
func (p *Player) SendTx(ctx context.Context, f frame.Frame) error {
	select {
	case <-p.ctx.Done():
		return p.ctx.Err()
	case <-ctx.Done():
		return ctx.Err()
	case p.in <- f:
		return nil
	}
}

// SendTimeout sends the given frame to the client, returning error if
// the time exceeded.
func (p *Player) SendTimeout(dur time.Duration, f frame.Frame) error {
	ctx, q := context.WithTimeout(context.Background(), dur)
	defer q()

	return p.SendTx(ctx, f)
}

// Command is a shorthand of p.Send(frame.New(cmd, err, data)).
func (p *Player) Command(cmd, err code.Code, data interface{}) {
	p.Send(frame.New(cmd, err, data))
}

// CommandTx is a shorthand of p.Send(ctx, frame.New(cmd, err, data)).
func (p *Player) CommandTx(ctx context.Context, cmd, err code.Code, data interface{}) error {
	return p.SendTx(ctx, frame.New(cmd, err, data))
}

// CommandTimeout is a shorthand of p.SendTimeout(dur, frame.New(cmd, err, data)).
func (p *Player) CommandTimeout(dur time.Duration, cmd, err code.Code, data interface{}) error {
	return p.SendTimeout(dur, frame.New(cmd, err, data))
}

// Output sets the channel where a player sends messages.
func (p *Player) Output(out chan<- frame.Frame) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.out = out
}

// ID returns the player's ID.
func (p *Player) ID() string {
	return p.id
}

// Name returns the player's name.
func (p *Player) Name() string {
	return p.name
}

// QuitWithCommand makes the player leave after sending given command.
func (p *Player) QuitWithCommand(cmd, err code.Code, data interface{}) {
	p.CommandTimeout(time.Second, cmd, err, data)
	time.Sleep(time.Second)
	p.Quit()
}

// emit sends the given frame to the player's output channel.
func (p *Player) emit(f frame.Frame) {
	f.From = p.id

	p.mu.RLock()
	defer p.mu.RUnlock()

	p.out <- f
}

// listen sends the frames received from the client to the output channel.
// It blocks until the client closed.
func (p *Player) listen() {
	for {
		select {
		case <-p.ctx.Done():
			return
		case f, ok := <-p.client.Receive():
			if !ok { // client exit
				p.emit(frame.New(code.PlayerExited, code.OK, nil))
				return
			}
			p.emit(f)
		}
	}
}

// Play handles the messages received from the client or the server.
// It blocks until Quit() is called.
func (p *Player) Play() {
	go p.listen()

	for {
		select {
		case <-p.ctx.Done():
			p.client.Close()
			return
		case f := <-p.in:
			p.client.Send() <- f
		}
	}
}

// New returns a pointer to a newly created Player with given id and names.
func New(id, name string, client Client) *Player {
	if id == "" || client == nil {
		panic("player: nil ID or client applied")
	}

	if name == "" {
		name = "???"
	}
	ctx, quit := context.WithCancel(context.Background())

	return &Player{
		id:     id,
		name:   name,
		client: client,
		ctx:    ctx,
		Quit:   quit,
		in:     make(chan frame.Frame),
		out:    make(chan frame.Frame),
	}
}
