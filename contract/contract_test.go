package contract

import (
	"testing"
	"time"
)

// TestContract_AddNodes ...
func TestContract_AddNodes(t *testing.T) {
	c := NewContract(Key(""), Node(""))
	ts := time.Now()
	ts.IsZero()
	c.AddNodes(true)
}
