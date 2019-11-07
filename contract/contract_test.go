package contract

import (
	"testing"
	"time"
)

// TestContract_AddNodes ...
func TestContract_AddNodes(t *testing.T) {
	c := NewContract(Key(""), Node(""))
	c.AddNodes(true, time.Now(), "/ip4/13.124.213.107/tcp/4001/ipfs/QmS9knxyQkdiFrGbzEb2FNqnq2yEzvBKx2pEGNBnXTWi7h")
}
