package handlers

import (
	"crypto/sha256"
	"fmt"
	"github.com/akshanshgusain/Video-Calling-App/pkg/chat"
	w "github.com/akshanshgusain/Video-Calling-App/pkg/webrtc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
	"os"
	"time"
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
	w.RoomConn(c, room.Peers)

}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}

	ws := "ws"
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		ws = "wss"
	}
	uuid, suuid, _ := createOrGetRoom(uuid)

	return c.Render("peer", fiber.Map{
		"RoomWebsocketAddr":   fmt.Sprintf("%s://%s/room/%s/websocket", ws, c.Hostname(), uuid),
		"RoomLink":            fmt.Sprintf("%s://%s/room/%s/", c.Protocol(), c.Hostname(), uuid),
		"ChatWebsocketAddr":   fmt.Sprintf("%s://%s/room/%s/chat/websocket", ws, c.Hostname(), uuid),
		"ViewerWebsocketAddr": fmt.Sprintf("%s://%s/room/%s/viewer/websocket", ws, c.Hostname(), uuid),
		"StreamLink":          fmt.Sprintf("%s://%s/stream/%s/", c.Protocol(), c.Hostname(), suuid),
		"Type":                "room",
	}, "layouts/main")

}

func createOrGetRoom(uuid string) (string, string, *w.Room) {
	w.RoomsLock.Lock()
	defer w.RoomsLock.Unlock()
	h := sha256.New()
	h.Write([]byte(uuid))

	suuid := fmt.Sprintf("%x", h.Sum(nil))
	if room := w.Rooms[uuid]; room != nil {
		if _, ok := w.Streams[suuid]; !ok {
			w.Streams[suuid] = room
		}
		return uuid, suuid, room
	}
	hub := chat.NewHub()
	p := &w.Peers()
	p.TrackLocals = make(map[string]*webrtc.TrackLocalStaticRTP)
	room := &w.Room{
		Peers: p,
		Hub:   hub,
	}
	w.Rooms[uuid] = room
	w.Streams[suuid] = room
	go hub.Run()
	return uuid, suuid, room
}

func RoomViewerWebSocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}
	w.RoomLock.Lock()
	if peer, ok := w.Rooms[uuid]; ok {
		w.RoomLock.Unlock()
		roomViewerConn(c, peer.Peers)
		return

	}
	w.RoomLock.Unlock()
}
func roomViewerConn(c *websocket.Conn, p *w.Peers) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {

		}
	}(c)

	for {
		select {
		case <-ticker.C:
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write([]byte(fmt.Sprintf("%d", len(p.Connections))))
		}
	}
}

type WebsocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
