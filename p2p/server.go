package p2p

import (
	"context"

	"github.com/libp2p/go-libp2p-peer"

	"github.com/gogo/protobuf/proto"
)

type TopicRegistry interface {
	SubscribeTopic(string, proto.Message, TopicHandler)
	UnsubscribeTopic(string)
}

type PeerManager interface {
	Discovery()
	clearPeer()
	updatePeer() int
}

type Server interface {
	Start() error
	Stop()

	// server interfaces
	PeerManager
	TopicRegistry

	// sender interfaces
	Send(context.Context, proto.Message, string, peer.ID) error
	Publish(context.Context, proto.Message, string) error

	// for test
	getHost() Host
	setContext(context.Context)
}
