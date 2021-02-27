# Fast ring buffer
Fast and flexible implementation of the ring buffer algorithm in Go but simply portable to any another language.
It provides add/remove to/from the tail or head of the ring buffer by single element or by coping array of elements.
Access to any element or to any continiuos subset (slice) of the ring buffer is prodived as well.

It is based on the following abstractions:
* Buffer - the ring buffer. It can be any fixed size indexed collection. Index [0..length-1] of preferred.
* Ring - a logical representation of the ring buffer over Buffer having Content between tail anf head.
* Range - a continiuos subset (slice) of the Content. It is consint of one or two Segment of the Buffer.
* Segment - a continiuos subset (slice) of the Buffer

```go

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

	// add bytes to the head
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

```
