package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddress := ":4000"
	tr := NewTcpTransport(listenAddress)
	assert.Equal(t, tr.listenAddress, listenAddress)

	//server
	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
