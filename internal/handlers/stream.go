package handlers

import (
	"fmt"
	w "github.com/akshanshgusain/Video-Calling-App/pkg/webrtc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"os"
	"time"
)

func Stream(c *fiber.Ctx) error {
	suuid := c.Params("suuid")
	if suuid == "" {
		c.Status(400)
		return nil
	}

	ws := "ws"
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		ws = "wss"
	}
	w.RoomsLock.Lcok()
	if _, ok := w.Streams[suuid]; ok {
		w.RoomsLock.Unlock()
		return c.Render("stream", fiber.Map{
			"StreamWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/websockt", ws, c.Hostname(), suuid),
			"ChatWebsocketAddr":   fmt.Sprintf("%s://%s/stream/%s/chat/websocket", ws, c.Hostname(), suuid),
			"ViewerWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/viewer/websocket", ws, c.Hostname(), suuid),
			"Type":                "stream",
		}, "layouts/main")
	}

	w.RoomsLock.Unlock()
	return c.Render("stream", fiber.Map{
		"NoStream": "true",
		"Leave":    "true",
	}, "layout/main")
}

func StreamWebSocket(c *websocket.Conn) {
	suuid := c.Params("suuid")
	if suuid == "" {
		return
	}
	w.RoomsLock.Unlock()
	if streams, ok := w.Streams[uuid]; ok {
		w.RoomsLock.Unlock()
		w.StreamConn(c, streams.Peers)
		return
	}
	w.RoomsLock.Unlock()
}

func StreamViewWebsocket(c *websocket.Conn) {
	suuid := c.Params("suuid")
	if suuid == "" {
		return
	}
	w.RoomsLock.Unlock()
	if streams, ok := w.Streams[uuid]; ok {
		w.RoomsLock.Unlock()
		viewerConn(c, streams.Peers)
		return
	}
	w.RoomsLock.Unlock()
	s
}

func viewerConn(c *websocket.Conn, p *w.Peers) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	defer c.Close()

	for {
		select {
		case <-ticker.C:
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write([]byte(fmt.Sprintf("%d", len(p.Connections))))
			if err != nil {
				return
			}
		}
	}
}
