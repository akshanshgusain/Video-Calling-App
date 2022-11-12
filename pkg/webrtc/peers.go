package webrtc

import (
	"sync"
)

type Peers struct {
}

func (p *Peers) DispatchKeyFrame() {
	ListLock sync.RWMutex
	Connections []PeerConnnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}
