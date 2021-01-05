package abool_test

import (
	"testing"

	"github.com/piersy/go-atomic-bool/abool"
	"github.com/stretchr/testify/assert"
)

func TestABool(t *testing.T) {
	ab := &abool.ABool{}
	assert.False(t, ab.Get())
	ab.Set(true)
	assert.True(t, ab.Get())
	ab.Set(false)
	assert.False(t, ab.Get())
}
