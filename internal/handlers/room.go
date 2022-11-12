package handlers

import (
	"fmt"
	w "github.com/akshanshgusain/Video-Calling-App/pkg/webrtc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}

func RoomWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}

	_, _, room := createOrGetRoom(uuid)

	// TODO: finish later

}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}

	uuid, suuid, _ := createOrGetRoom(uuid)

	// TODO: finish later

}

func createOrGetRoom(uuid string) (string, string, *w.Room) {

}

func RoomViewerWebSocket(c *websocket.Conn) {

}

func RoomViewerConn(c *websocket.Conn, p *w.Peers) {

}

type WebsocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
