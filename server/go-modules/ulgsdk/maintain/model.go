package maintain

import "time"

type Maintain struct {
	ID            int       `json:"id"`
	GameID        string    `json:"game_id"`
	Status        int       `json:"status"`
	NotifyTime    time.Time `json:"notify_time"`
	KickTime      time.Time `json:"kick_time"`
	CloseTime     time.Time `json:"close_time"`
	NotifyMessage string    `json:"notify_message"`
	KickMessage   string    `json:"kick_message"`
}
