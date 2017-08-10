package utils

import (
	//	"github.com/davecgh/go-spew/spew"
	// "github.com/stretchr/testify/assert"
	"testing"
)

func init() {
}

func logIndirect2(args ...interface{}) {
	log := NewLogger().Build()
	log.Debug(args...)
}
func logIndirect(args ...interface{}) {
	logIndirect2(args...)
}

func TestDebug(t *testing.T) {
	// assert := assert.New(t)

	log := NewLogger().Build()
	log.Debug("test info")
	log.Debug("test info", "1")
	log.Debug("test info", "1", 2)
	log.Debug("test info", "1", 2, 3)
	log.Debug("test info", "1", 2, 3, 4)
	log.Debug("test info"+" foobar", "1", 2)
}

func TestLogIndirect(t *testing.T) {
	// assert := assert.New(t)
	logIndirect("test indirect")
}
