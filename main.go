package main

import (
	"fmt"
	"log"

	"github.com/Jayant-issar/go-distributed-file-storage/p2p"
)

func main() {
	tr := p2p.NewTcpTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
	fmt.Println("done hai ji")
}
