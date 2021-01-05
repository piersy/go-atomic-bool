package abool

import (
	"sync/atomic"
)

// ABool provides a simple atomic boolean.
type ABool struct {
	flag int32
}

// Get returns the current value.
func (ab *ABool) Get() bool {
	return atomic.LoadInt32(&ab.flag) == 1
}

// Set sets the value to b.
func (ab *ABool) Set(b bool) {
	var val int32
	if b {
		val = 1
	}
	atomic.StoreInt32(&ab.flag, val)
}
