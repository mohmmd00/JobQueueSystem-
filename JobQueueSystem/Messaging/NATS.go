package messaging

import "github.com/nats-io/nats.go"



func Connect() (*nats.Conn , error){
	return nats.Connect(nats.DefaultURL)
}


func Publish(nc *nats.Conn, subject string, data []byte) error {
	return nc.Publish(subject, data)
}

func Subscribe(nc *nats.Conn, subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return nc.Subscribe(subject, handler)
}