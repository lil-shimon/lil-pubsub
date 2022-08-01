package websocket

type Room struct {
	Subscriptions []*Subscription
}

type Subscription struct {
	Topic string
	Client *Client
}

func (room *Room) AddSubscription(subscription *Subscription) {
	room.Subscriptions = append(room.Subscriptions, subscription)
}

func (room *Room) GetSubscription(topic string) []Subscription {
	var subs []Subscription

	for _, sub := range room.Subscriptions {
		if sub.Topic == topic {
			subs = append(subs, *sub)
		}
	}

	return subs
}

func (room *Room) Publish(msg []byte, topic string) {
	subs := room.GetSubscription(topic)
	for _, sub := range subs {
		err := sub.Client.Send(msg)
		if err != nil {
			return
		}
	}
}