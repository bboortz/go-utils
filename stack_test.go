package utils

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	assert := assert.New(t)

	var result []byte = Stack()
	assert.NotNil(result)
}

func TestPrintStack(t *testing.T) {
	PrintStack()
}

func TestTrace(t *testing.T) {
	Trace()
}

func TestGetCurrentMethodName(t *testing.T) {
	assert := assert.New(t)

	var result string = GetCurrentMethodName()
	assert.NotNil(result)
	assert.Equal("_/app.TestGetCurrentMethodName", result)
}

func TestGetCallingMethodName(t *testing.T) {
	assert := assert.New(t)

	var result string = GetCallingMethodName()
	assert.NotNil(result)
	assert.Equal("testing.tRunner", result)
}
