package sockettool

import (
	"fmt"
	"sync"
)

var HubObj *Hub

type Hub struct {
	mu      sync.Mutex
	clients map[ConnClient]bool
	// broadcast  chan interface{}
	register   chan ConnClient
	unregister chan ConnClient
	channels   map[string]*Channel
}

func NewHub() *Hub {
	if HubObj != nil {
		return HubObj
	}
	HubObj = &Hub{
		// broadcast:  make(chan interface{}),
		register:   make(chan ConnClient),
		unregister: make(chan ConnClient),
		clients:    make(map[ConnClient]bool),
		channels:   make(map[string]*Channel),
	}
	return HubObj
}

func (h *Hub) CheckUser(name string) ConnClient {
	h.mu.Lock()
	defer h.mu.Unlock()
	for conn, ok := range h.clients {
		if ok {
			if conn.User() == name {
				return conn
			}
		}
	}
	return nil
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				channels := client.GetChannels()
				for _, channel := range channels {
					c, found := h.channels[channel]
					if !found {
						continue
					}
					c.UnSubscribe(client)
				}
				delete(h.clients, client)
				client.Close()
			}
		}
	}
}

func (h *Hub) Broadcast(msg interface{}) {
	for client := range h.clients {
		if client.Avaliable() {
			client.Send(msg)
		}
	}
}

func (h *Hub) LobyBroadcast(msg interface{}) {
	for conn, _ := range h.clients {
		if conn.GetChannel() == "lobby" {
			if conn.Avaliable() {
				conn.Send(msg)
			}
		}
	}
}

func (h *Hub) BroadcastByChannel(channel string, msg interface{}) {
	c, found := h.channels[channel]
	if !found {
		return
	}
	c.Broadcast(msg)
}

func (h *Hub) SendToChannelUser(channel, user string, msg interface{}) {
	c, found := h.channels[channel]
	if !found {
		return
	}
	c.Send(user, msg)
}

func (h *Hub) Regist(c ConnClient) {
	h.register <- c
}

func (h *Hub) UnRegist(c ConnClient) {
	h.unregister <- c
}

func (h *Hub) RegistChannel(channel string) {
	h.channels[channel] = NewChannel(channel)
}

func (h *Hub) UnRegistChannel(channel string) {
	c, found := h.channels[channel]
	if !found {
		return
	}
	for client, _ := range c.subscriber {
		client.UnSubscribeChannel(channel)
	}
	delete(h.channels, channel)
}

func (h *Hub) RegistUserChannel(conn ConnClient, channel string) {
	c, found := h.channels[channel]
	if !found {
		return
	}
	c.Subscribe(conn)
}

func (h *Hub) UnRegistUserChannel(conn ConnClient, channel string) {
	c, found := h.channels[channel]
	if !found {
		return
	}
	c.UnSubscribe(conn)
}

func (h *Hub) GetChannel(channelName string) (*Channel, bool) {
	c, found := h.channels[channelName]
	return c, found
}

func (h *Hub) GetDupConn(token, gameToken string) ConnClient {
	// h.mu.Lock()
	// defer h.mu.Unlock()
	for client, _ := range h.clients {
		fmt.Println()
		if client.Token() == token &&
			client.GameToken() == gameToken {
			return client
		}
	}
	return nil
}
