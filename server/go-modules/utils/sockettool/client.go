package sockettool

type ConnClient interface {
	User() string
	Token() string
	GameToken() string
	Reader()
	Writer()
	Send(interface{})
	Close()
	Avaliable() bool
	SubscribeChannel(channel string)
	UnSubscribeChannel(channel string)
	GetChannels() []string
	GetChannel() string
	AddOrder(channel string, res interface{})
	GetOrderMapping(channel string) interface{}
	GetOrders(channel string) interface{}
	GetAllOrders() interface{}
	FlushOrder(channel string)
}
