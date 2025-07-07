package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCP Peer repersents the remote node over a TCP established connection
type TCPPeer struct {
	//conn is the underlying connection of the peer
	conn net.Conn

	//if we dial and retrive a connection => outbound == true
	//if we accept and retrive a connection => outbound == false
	outBound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outBound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listner       net.Listener
	handshakeFunc HandshakeFunc

	decoder Decoder
	mu      sync.RWMutex
	peers   map[net.Addr]Peer
}

func NewTcpTransport(listnerAddr string) *TCPTransport {
	return &TCPTransport{
		handshakeFunc: NOPHandshakeFunc,
		listenAddress: listnerAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listner, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listner.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		fmt.Printf("New incoming connection %+v\n", conn)

		go t.handleConn(conn)

	}
}

type Temp struct {
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.handshakeFunc(peer); err != nil {

	}

	//read loop
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s/n", err)
		}
	}
	fmt.Printf("New incoming connection %v\n", peer)
}
