package fast_ring

import (
	"errors"
)

type Ring struct { // Ring algorith provider
	Size          int  // Buffer size
	CheckOverflow bool // Check buffer owerflow otherwise move tail
	isNotEmpty    bool
	head          int
	tail          int
}

type ringRangeSegment struct {
	Start  int // Absolute start point
	RStart int // Relative start point
	Length int // Segment length
}

type RingRange [2]ringRangeSegment // Range of the ring descriptor

// Length of the ring content 0..Size-1
func (r *Ring) Length() int {
	tmp := r.head - r.tail
	if tmp < 0 {
		return tmp + r.Size
	} else if (tmp == 0) && r.isNotEmpty {
		return r.Size
	}
	return tmp
}

// Check has content
func (r *Ring) IsEmpty() bool {
	return !r.isNotEmpty
}

// Check ring is full
func (r *Ring) IsFull() bool {
	return r.isNotEmpty && (r.head == r.tail)
}

// First element of then tail
func (r *Ring) Tail() int {
	return r.tail
}

// First empty element after the head
func (r *Ring) Head() int {
	return r.head
}

// Elements from head toward to the tail
func (r *Ring) GetFromHead(i int) int {
	return r.Shift(r.head, -i-1)
}

// Elements from tail toward to the head
func (r *Ring) FromTail(i int) int {
	return r.Shift(r.tail, i)
}

// Get content range adjacent to the head
func (r *Ring) GetRangeFromHead(len int) (res RingRange) {
	res, _ = r.GetRange(r.Shift(r.head, -len), len)
	return
}

// Get content range from tail
func (r *Ring) GetRangeFromTail(len int) (res RingRange) {
	res, _ = r.GetRange(r.tail, len)
	return
}

// move point inside ring in range -Size..+Size from the start
func (r *Ring) Shift(start int, shift int) (res int) {
	res = start + shift
	if res < 0 {
		res += r.Size
	} else {
		if res > r.Size {
			res -= r.Size
		}
	}
	return
}

// Get range from the absolute start of a given length and the next point after one
func (r *Ring) GetRange(start int, len int) (res RingRange, next int) {
	if len > r.Size {
		panic(errors.New("Out of ring buffer"))
	}
	res[0].Start = start
	res[0].Length = len
	//res[0].Length = 0
	next = start + len
	if next >= r.Size {
		next -= r.Size
		res[0].Length = r.Size - start
		res[1].Length = len - res[0].Length
		//res[1].Start = 0
	}
	//res[0].RStart = 0
	res[1].RStart = res[0].Length
	return
}

func (r *Ring) AddToHead() (i int) {
	i = r.head
	if r.isNotEmpty && (r.head == r.tail) { // IsFull
		if r.CheckOverflow {
			panic(errors.New("Ring buffer overflow"))
		}
		r.head++
		if r.head >= r.Size {
			r.head -= r.Size
		}
		r.tail = r.head
	} else {
		r.head++
		if r.head >= r.Size {
			r.head -= r.Size
		}
	}
	r.isNotEmpty = true
	return
}

func (r *Ring) AddToTail() int {
	if r.isNotEmpty && (r.head == r.tail) { // IsFull
		if r.CheckOverflow {
			panic(errors.New("Ring buffer overflow"))
		}
		r.tail--
		if r.tail < 0 {
			r.tail += r.Size
		}
		r.head = r.tail
	} else {
		r.tail--
		if r.tail < 0 {
			r.tail += r.Size
		}
	}
	r.isNotEmpty = true
	return r.tail
}

func (r *Ring) RemoveFromHead() int {
	if !r.isNotEmpty {
		panic(errors.New("Out of ring buffer"))
	}
	r.head -= 1
	if r.head < 0 {
		r.head += r.Size
	}
	if r.tail == r.head {
		r.isNotEmpty = false
	}
	return r.head
}

func (r *Ring) RemoveFromTail() (i int) {
	if !r.isNotEmpty {
		panic(errors.New("Out of ring buffer"))
	}
	i = r.tail
	r.tail += 1
	if r.tail >= r.Size {
		r.tail -= r.Size
	}
	if r.tail == r.head {
		r.isNotEmpty = false
	}
	return
}

func (r *Ring) AddRangeToHead(len int) (res RingRange) {
	res, next := r.GetRange(r.head, len)
	if len == 0 {
		return
	}
	if len > r.Size-r.Length() {
		if r.CheckOverflow {
			panic(errors.New("Ring buffer overflow"))
		}
		r.tail = next
	}
	r.head = next
	r.isNotEmpty = true
	return
}

func (r *Ring) RemoveRangeFromHead(len int) (res RingRange) {
	d := r.Length() - len
	if d < 0 {
		panic(errors.New("Out of ring buffer"))
	}
	r.head = r.Shift(r.head, -len)
	res, _ = r.GetRange(r.head, len)
	if len == 0 {
		return
	}
	r.isNotEmpty = d > 0
	return
}

func (r *Ring) AddRangeToTail(len int) (res RingRange) {
	start := r.Shift(r.tail, -len)
	res, _ = r.GetRange(start, len)
	if len == 0 {
		return
	}
	if len > r.Size-r.Length() {
		if r.CheckOverflow {
			panic(errors.New("Ring buffer overflow"))
		}
		r.head = start
	}
	r.tail = start
	r.isNotEmpty = true
	return
}

func (r *Ring) RemoveRangeFromTail(len int) (res RingRange) {
	d := r.Length() - len
	if d < 0 {
		panic(errors.New("Out of ring buffer"))
	}
	res, r.tail = r.GetRange(r.tail, len)
	if len == 0 {
		return
	}
	r.isNotEmpty = d > 0
	return
}
