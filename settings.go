package rabbit

const (
	Host         = "localhost"
	UserName     = "guest"
	Password     = "guest"
	RoutingKey   = "onyore"
	ExchangeName = "amq.direct"
	FullHost     = "amqp://" + UserName + ":" + Password + "@localhost"
)
