package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInc(t *testing.T) {
	assert.EqualValues(t, 0, traceLevel)
	trace("")
	assert.EqualValues(t, 1, traceLevel)
	for i := 0; i < 10; i++ {
		trace("")
	}
	assert.EqualValues(t, 11, traceLevel)
	for i := 0; i < 5; i++ {
		untrace("")
	}
	assert.EqualValues(t, 6, traceLevel)
}
