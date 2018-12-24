package number

import (
	"math/rand"
)

func SeqUint8(len uint) (result []uint8) {
	var i uint

	for i = 0; i < len; i++ {
		result = append(result, Uint8())
	}

	return result
}

func Uint8() uint8 {
	return uint8(rand.Uint32()>>2)
}

func SeqUint16(len int) (result []uint16) {
	for i := 0; i < len; i++ {
		result = append(result, Uint16())
	}

	return result
}

func Uint16() uint16 {
	return uint16(rand.Uint32()>>1)
}

func SeqUint32(len int) (result []uint32) {
	for i := 0; i < len; i++ {
		result = append(result, Uint32())
	}

	return result
}

func Uint32() uint32 {
	return rand.Uint32()
}

func SeqUint64(len uint) (result []uint64) {
	var i uint

	for i = 0; i < len; i++ {
		result = append(result, Uint64())
	}

	return result
}

func Uint64() uint64 {
	return rand.Uint64()
}

func UintGenerator(len uint) chan int64 {
	ch := make(chan int64)

	go func() {
		for len > 0 {
			len--
			ch <- rand.Int63n(255)
		}
		close(ch)
	} ()

	return ch
}
