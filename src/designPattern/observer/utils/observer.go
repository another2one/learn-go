package utils

type Observer struct {
	clients []*Client
}

func NewObserver() *Observer {
	o := new(Observer)
	o.clients = make([]*Client, 0)
	return o
}

func (observer *Observer) Notify(i int) {
	for _, c := range observer.clients {
		c.Notify(i)
	}
}

func (observer *Observer) Bind(client *Client) {
	observer.clients = append(observer.clients, client)
}

func (observer *Observer) Unbind(client *Client) {
	for i, o := range observer.clients {
		if o == client {
			observer.clients = append(observer.clients[:i], observer.clients[i+1:]...)
		}
	}
}
