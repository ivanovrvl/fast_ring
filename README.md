# Fast ring buffer
Fast and flexible implementation of the ring buffer algorithm in Go but simply portable on any another language.
It provides add/remove to/from the tail or head of the ring buffer by single element or by coping array of elements.
Access to any element or to any continiuos subset (slice) of the ring buffer is prodived as well.

It is based on the following abstractions:
Buffer - the ring buffer. It can be any fixed size indexed collection. Index [0..length-1] of preferred.
Ring - the logical representation of the ring buffer over Buffer.
RingRange - a descriptor of any continiuos subset (slice) of the Ring. It is consint of one or two Segment of the Buffer.
Segment - a continiuos subset (slice) of the Buffer
