package sockettool

import (
	"sync"
)

type Channel struct {
	mu         sync.Mutex
	id         int
	name       string
	subscriber map[ConnClient]struct{}
	subUserMap map[string]ConnClient
}

func NewChannel(name string) *Channel {
	return &Channel{
		name:       name,
		subscriber: make(map[ConnClient]struct{}),
		subUserMap: make(map[string]ConnClient),
	}
}

func (c *Channel) GetUser() []string {
	res := []string{}
	for k, _ := range c.subUserMap {
		res = append(res, k)
	}
	return res
}

func (c *Channel) GetUserConn() map[string]ConnClient {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.subUserMap
}

func (c *Channel) Send(user string, msg interface{}) {
	if client, found := c.subUserMap[user]; found {
		if client.Avaliable() {
			client.Send(msg)
		}
	}
}

func (c *Channel) Broadcast(msg interface{}) {
	if c.subscriber == nil {
		c.subscriber = make(map[ConnClient]struct{})
	}
	for client, _ := range c.subscriber {
		if client.Avaliable() {
			client.Send(msg)
		}
	}
}

func (c *Channel) Subscribe(client ConnClient) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.subscriber == nil {
		c.subscriber = make(map[ConnClient]struct{})
		c.subUserMap = make(map[string]ConnClient)
	}
	_, found := c.subscriber[client]
	if !found {
		c.subscriber[client] = struct{}{}
		c.subUserMap[client.User()] = client
	}
	client.SubscribeChannel(c.name)
}

func (c *Channel) UnSubscribe(client ConnClient) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.subscriber == nil {
		c.subscriber = make(map[ConnClient]struct{})
		c.subUserMap = make(map[string]ConnClient)
	}
	if _, found := c.subscriber[client]; found {
		delete(c.subscriber, client)
	}
	if _, found := c.subUserMap[client.User()]; found {
		delete(c.subUserMap, client.User())
	}
	client.UnSubscribeChannel(c.name)
}
