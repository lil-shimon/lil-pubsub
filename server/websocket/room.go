package websocket

type Room struct {
	Clients []*Client
}

func (room *Room) AddClient(client *Client) {
	room.Clients = append(room.Clients, client)
}

func (room *Room) GetClients() []Client {
	var cs []Client

	for _, client := range room.Clients {
		cs = append(cs, *client)
	}

	return cs
}

func (room *Room) Publish(msg []byte) {
	for _, client := range room.Clients {
		client.Send(msg)
	}
}