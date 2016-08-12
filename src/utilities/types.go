package utilities

import "errors"

// HTTPProtocol = HTTP | HTTPS
type Protocol int
const (
	HTTP Protocol = iota
	HTTPS
	ZMQ
)

type Behavior int
const (
	PRIVATE Behavior = iota
	REQUEST
	REPLY
	PUBLISH
	SUBSCRIBE
	PUSH
	PULL
)

func (b Behavior) Dual () Behavior {
	switch b {
	case PRIVATE:
		return PRIVATE
	case REQUEST:
		return REPLY
	case REPLY:
		return REQUEST
	case PUBLISH:
		return SUBSCRIBE
	case SUBSCRIBE:
		return PUBLISH
	case PUSH:
		return PULL
	case PULL:
		return PUSH
	default:
		panic("Invalid behavior")
	}
}

type Endpoint struct {
	Protocol Protocol
	Behavior Behavior
}

type Labelling map[string]Endpoint

func (peer Peer) Dual () (dualPeer Peer) {
	dualPeer.Behavior = peer.Behavior.Dual()
	dualPeer.Protocol = peer.Protocol
	for id, subPeer := range peer.Channels {
		dualPeer.Channels[id] = subPeer.Dual()
	}
	return dualPeer
}



func (p0 Peer) Combine (p1 Peer) (p2 Peer, err error) {

}

func (p0 Peer) Connect (p1 Peer) (p2 Peer, err error) {
	if p0.Behavior == PRIVATE || p1.Behavior == PRIVATE {
		return nil, errors.New(P)
	} else {
		matching_protocols := p0.Protocol == p1.Protocol
		dual_behaviors := p0.Behavior == p1.Behavior.dual()
		return matching_protocols && dual_behaviors
	}
}

func (p0 Peer) Merge (p1 Peer) Peer {

}


// String functions

func (p HTTPProtocol) String () string {
	if p == HTTP {
		return "http"
	} else {
		return "https"
	}
}
