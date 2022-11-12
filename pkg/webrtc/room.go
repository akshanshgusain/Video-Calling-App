package webrtc

import (
	"github.com/akshanshgusain/Video-Calling-App/pkg/chat"
	"github.com/gofiber/websocket/v2"
	"log"
	"sync"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

func RoomConn(c *websocket.Conn, p *Peers) {
	var config webrtc.Configuration
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Printf(err)
		return
	}

	newPeer := PeerConnectionState{
		PeerConnection: peerConnection,
		WebSocket:      &ThreadSafeWriter{},
		Conn:           c,
		Mutex:          sync.Mutex{},
	}
}
