package webrtc

import (
	"github.com/akshanshgusain/Video-Calling-App/pkg/chat"
	"sync"
)

type Room struct {
	Peers *Peers
	Hub *chat.Hub
}
type Peers struct {
}

func (p *Peers) DispatchKeyFrame() {
	ListLock sync.RWMutex
	Connections []PeerConnnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}
