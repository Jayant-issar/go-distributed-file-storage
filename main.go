package main

import (
	"fmt"
	"log"

	"github.com/Jayant-issar/go-distributed-file-storage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
	}
	tr := p2p.NewTcpTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
	fmt.Println("done hai ji")
}
