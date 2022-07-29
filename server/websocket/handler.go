package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{}
var rooms = Room{}

func ServeWs(c echo.Context) error {

	topic := c.Param("topic")

	log.Printf("topic: %s", topic)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		c.Logger().Error(err)
	}

	defer ws.Close()

	client := &Client{Ws: ws}

	rooms.AddSubscription(&Subscription{Topic: topic, Client: client})

	for {
		_, msg, err := ws.ReadMessage()

		if err != nil {
			c.Logger().Error(err)
			break
		}

		rooms.Publish(msg, topic)
	}

	return nil
}