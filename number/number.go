package number

import "math/rand"

func Uint8(len int) (result []uint8) {
	for i := 0; i < len; i++ {
		result = append(result, uint8(rand.Uint32()>>2))
	}

	return result
}

func Uint16(len int) (result []uint16) {
	for i := 0; i < len; i++ {
		result = append(result, uint16(rand.Uint32()>>1))
	}

	return result
}

func Uint32(len int) (result []uint32) {
	for i := 0; i < len; i++ {
		result = append(result, rand.Uint32())
	}

	return result
}

func Uint64(len int) (result []uint64) {
	for i := 0; i < len; i++ {
		result = append(result, rand.Uint64())
	}

	return result
}
