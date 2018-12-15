package number

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const loop = 100

func TestUint8(t *testing.T) {
	var result []uint8

	for i := 0; i < loop; i++ {
		result = Uint8(i)
		for _, element := range result {
			inRange(t, 0, 2<<8, int64(element))
		}
	}
}

func TestUint32(t *testing.T) {
	var result []uint32

	for i := 0; i < loop; i++ {
		result = Uint32(i)
		for _, element := range result {
			inRange(t, 0, 2<<31, int64(element))
		}
	}
}

func inRange(t *testing.T, min int64, max int64, value int64) {
	assert.True(t, min <= value, fmt.Sprintf("Value %d less than minimum %d", value, min))
	assert.True(t, max >= value, fmt.Sprintf("Value %d greater than maximum %d", value, max))
}