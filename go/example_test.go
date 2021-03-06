package fast_ring

import (
	"fmt"
	"testing"
)

func exampleByteRingBuffer() {

	// An example how to use the byte ring buffer

	// Initialization
	ring := Ring{ // the ring provider
		Size:          100,
		CheckOverflow: false, // will override tail if needed
	}
	ringBuffer := make([]byte, ring.Size) // the ring buffer

	// Usage

	// add one byte to the head
	ringBuffer[ring.AddToHead()] = 1
	fmt.Println("Len", ring.Length())

	// add one byte to the tail
	ringBuffer[ring.AddToTail()] = 100
	fmt.Println("Len", ring.Length())

	// add bytes to head
	var toWrite1 = [8]byte{2, 3, 4, 5, 6, 7, 8, 9}
	// using loop over segments
	for _, seg := range ring.AddRangeToHead(len(toWrite1)) {
		copy(ringBuffer[seg.Start:seg.Start+seg.Length], toWrite1[seg.RStart:seg.RStart+seg.Length])
	}

	// add bytes to the tail
	var toWrite2 = [3]byte{103, 102, 101}
	// without loop
	seg := ring.AddRangeToTail(len(toWrite2))
	copy(ringBuffer[seg[0].Start:seg[0].Start+seg[0].Length], toWrite2[0:seg[0].Length])
	if seg[1].Length > 0 {
		copy(ringBuffer[seg[1].Start:seg[1].Start+seg[1].Length], toWrite2[seg[1].Length:])
	}

	// iterate over all elements from the tail to head
	for _, seg := range ring.GetRangeFromTail(ring.Length()) {
		for i := 0; i < seg.Length; i++ {
			fmt.Print(" ", ringBuffer[seg.Start+i])
		}
	}
	fmt.Println()

	// read bytes from the tail
	var toRead [10]byte
	for _, seg := range ring.RemoveRangeFromTail(len(toRead)) {
		copy(toRead[seg.RStart:seg.RStart+seg.Length], ringBuffer[seg.Start:seg.Start+seg.Length])
	}
	fmt.Println(toRead)
	fmt.Println("Len", ring.Length())

	// read one byte from the tail
	fmt.Println(ringBuffer[ring.RemoveFromTail()])
	// read one byte from the head
	fmt.Println(ringBuffer[ring.RemoveFromHead()])
	// read one more from the head
	fmt.Println(ringBuffer[ring.RemoveFromHead()])

	fmt.Println("Len", ring.Length())

}
func TestByteRingBuffer(t *testing.T) {
	exampleByteRingBuffer()
}
