package main

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/labstack/echo/v4"
	"github.com/sairash/chitosocket"
)

func websocket_example(c echo.Context) error {
	_, _, _, err := chitosocket.UpgradeConnection(c.Request(), c.Response())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	go chitosocket.StartUp()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, chitosocket.Hub.Subs[id])
	})

	chitosocket.On["message"] = func(subs **chitosocket.Subscriber, op ws.OpCode, data map[string]interface{}) {
		sub := *subs
		length_of_room := len(sub.Room)
		if length_of_room > 0 {
			room := sub.Room[length_of_room-1]
			chitosocket.Emit("message", room, op, data)
		}
	}

	chitosocket.On["add_to_room"] = func(subs **chitosocket.Subscriber, op ws.OpCode, data map[string]interface{}) {
		sub := *subs
		if data["room_id"] != nil {
			sub.AddToRoom(data["room_id"].(string))
			chitosocket.Emit("room_event", []string{data["room_id"].(string)}, op, map[string]interface{}{"user_id": data["user_id"], "message": "New user just joined the room "})
		}
	}

	e.GET("/ws/", websocket_example)
	e.Logger.Fatal(e.Start(":6969"))
}
