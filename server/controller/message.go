package controller

import (
	"encoding/base64"
	"encoding/json"

	"github.com/gorilla/websocket"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/conf"
)

var modeB64 bool = conf.Base64Enable

func readJSON(conn *websocket.Conn, v interface{}) error {
	_, r, err := conn.NextReader()
	if err != nil {
		return err
	}

	if modeB64 {
		r = base64.NewDecoder(base64.URLEncoding, r)
	}

	return json.NewDecoder(r).Decode(v)
}

func writeJSON(conn *websocket.Conn, v interface{}) error {
	w, err := conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	defer w.Close()

	if modeB64 {
		o := base64.NewEncoder(base64.URLEncoding, w)
		defer o.Close()

		return json.NewEncoder(o).Encode(v)
	}
	return json.NewEncoder(w).Encode(v)
}

func EnableBase64(enabled bool) {
	modeB64 = enabled
}
