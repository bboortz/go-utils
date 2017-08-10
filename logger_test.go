package utils

import (
	//	"github.com/davecgh/go-spew/spew"
	// "github.com/stretchr/testify/assert"
	"testing"
)

func init() {
}

func TestDebug(t *testing.T) {
	// assert := assert.New(t)

	log := NewLogger().Build()
	log.Debug("test info")
}
