package network

import (
	"context"

	bsmsg "github.com/ipfs/go-ipfs/exchange/bitswap/message"

	ifconnmgr "gx/ipfs/QmXuucFcuvAWYAJfhHV2h4BYreHEAsLSsiquosiXeuduTN/go-libp2p-interface-connmgr"
	protocol "gx/ipfs/QmZNkThpqfVXs9GNbexPrfBbXSLNYeKrE7jwFM2oqHbyqN/go-libp2p-protocol"
	cid "gx/ipfs/QmapdYm1b22Frv3k17fqrBYTFRxwiaVJkB299Mfn33edeB/go-cid"
	peer "gx/ipfs/QmdVrMn1LhB4ybb8hMVaMLXnA8XRSewMnK6YqXKXoTcRvN/go-libp2p-peer"
)

var (
	// These two are equivalent, legacy
	ProtocolBitswapOne    protocol.ID = "/ipfs/bitswap/1.0.0"
	ProtocolBitswapNoVers protocol.ID = "/ipfs/bitswap"

	ProtocolBitswap protocol.ID = "/ipfs/bitswap/1.1.0"
)

// BitSwapNetwork provides network connectivity for BitSwap sessions
type BitSwapNetwork interface {

	// SendMessage sends a BitSwap message to a peer.
	SendMessage(
		context.Context,
		peer.ID,
		bsmsg.BitSwapMessage) error

	// SetDelegate registers the Reciver to handle messages received from the
	// network.
	SetDelegate(Receiver)

	ConnectTo(context.Context, peer.ID) error

	NewMessageSender(context.Context, peer.ID) (MessageSender, error)

	ConnectionManager() ifconnmgr.ConnManager

	Routing
}

type MessageSender interface {
	SendMsg(context.Context, bsmsg.BitSwapMessage) error
	Close() error
	Reset() error
}

// Implement Receiver to receive messages from the BitSwapNetwork
type Receiver interface {
	ReceiveMessage(
		ctx context.Context,
		sender peer.ID,
		incoming bsmsg.BitSwapMessage)

	ReceiveError(error)

	// Connected/Disconnected warns bitswap about peer connections
	PeerConnected(peer.ID)
	PeerDisconnected(peer.ID)
}

type Routing interface {
	// FindProvidersAsync returns a channel of providers for the given key
	FindProvidersAsync(context.Context, *cid.Cid, int) <-chan peer.ID

	// Provide provides the key to the network
	Provide(context.Context, *cid.Cid) error
}
