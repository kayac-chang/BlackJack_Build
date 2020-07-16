package lobby

import (
	"errors"
	"fmt"
	"strconv"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
	"go.uber.org/zap"
)

var defaultCommands = CommandSet{
	Wellcome: 0,
	Move:     1,
	Exit:     2,
}

func defaultRoomIDer(v interface{}) string {
	switch id := v.(type) {
	case string:
		return id
	case int:
		return strconv.Itoa(id)
	case fmt.Stringer:
		return id.String()
	default:
		return ""
	}
}

// A CommandSet is a set of predefined commands. Which can be
// customized in different games.
type CommandSet struct {
	Wellcome code.Code
	Move     code.Code
	Exit     code.Code
}

// Options defines the configuration fields of a lobby which can be
// customized in different games.
type Options struct {
	Commands CommandSet
	Logger   *zap.Logger

	// RoomID is a function which returns the ID of a room by given value.
	RoomID func(interface{}) string

	// HandleDuplicated is the policy to handle duplicated logged-in players.
	HandleDuplicated Policy
}

// Wellcome is a shorthand of o.Commands.Wellcome.
func (o *Options) Wellcome() code.Code {
	return o.Commands.Wellcome
}

// Move is a shorthand of o.Commands.Move.
func (o *Options) Move() code.Code {
	return o.Commands.Move
}

// Exit is a shorthand of o.Commands.Exit.
func (o *Options) Exit() code.Code {
	return o.Commands.Exit
}

// Valid returns if o is OK. If not, it returns an error.
func (o *Options) Valid() error {
	cmd := o.Commands
	if cmd.Wellcome == cmd.Move ||
		cmd.Wellcome == cmd.Exit ||
		cmd.Move == cmd.Exit {
		return fmt.Errorf("lobby: same command values: %d, %d, %d",
			cmd.Wellcome, cmd.Move, cmd.Exit)
	}

	if o.Logger == nil {
		return errors.New("lobby: nil logger in options")
	}

	if o.RoomID == nil {
		return errors.New("lobby: nil roomID converter")
	}

	if _, ok := o.HandleDuplicated.(policy); !ok {
		return errors.New("lobby: unknown policy")
	}

	return nil
}

// DefaultOptions returns options with predefined settings.
func DefaultOptions() *Options {
	z, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return &Options{
		Commands:         defaultCommands,
		Logger:           z,
		RoomID:           defaultRoomIDer,
		HandleDuplicated: RejectLatter,
	}
}
