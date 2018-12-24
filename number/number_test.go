package number

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	loop           = 100
	minUnsignedInt = 0
	maxUint8       = 2<<8 - 1
	maxUint16      = 2<<16 - 1
	maxUint32      = 2<<31 - 1
	maxUint64      = 2<<63 - 1
)

func TestSequenceUint8(t *testing.T) {
	var result []uint8

	result = SeqUint8(loop)

	for _, element := range result {
		inRange(t, minUnsignedInt, maxUint8, uint64(element))
	}
}

func TestUint8(t *testing.T) {
	inRange(t, minUnsignedInt, maxUint8, uint64(Uint8()))
}

func TestUint16(t *testing.T) {
	var result []uint16

	for i := 0; i < loop; i++ {
		result = SeqUint16(i)
		for _, element := range result {
			inRange(t, minUnsignedInt, maxUint16, uint64(element))
		}
	}
}

func TestUint32(t *testing.T) {
	var result []uint32

	for i := 0; i < loop; i++ {
		result = SeqUint32(i)
		for _, element := range result {
			inRange(t, minUnsignedInt, maxUint32, uint64(element))
		}
	}
}

func TestUint64(t *testing.T) {
	var result []uint64
	var i uint

	for i = 0; i < loop; i++ {
		result = SeqUint64(i)
		for _, element := range result {
			inRange(t, minUnsignedInt, maxUint64, element)
		}
	}
}

func TestUintGenerator(t *testing.T) {
	for element := range UintGenerator(loop) {
		inRange(t, minUnsignedInt, maxUint64, uint64(element))
	}
}

func inRange(t *testing.T, min uint64, max uint64, value uint64) {
	assert.True(t, min <= value, fmt.Sprintf("Value %d smaller than minimum %d", value, min))
	assert.True(t, max >= value, fmt.Sprintf("Value %d greater than maximum %d", value, max))
}
