package p2p

// Peer is an interface that repsersents a remote node.
type Peer interface {
}

// Transport is anythng that handles the communication
// between the nodes in a network is can be in the form, TCP, UDP, websockets...
type Transport interface {
	ListenAndAccept() error
}
