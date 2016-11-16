// Package rle implements run-length encoding for various types (just ints at the moment).
package rle

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Int64Decoder is what it sounds like.
type Int64Decoder struct {
	Value int64
	Run   int64
	buf   *bytes.Buffer
	err   error
}

// NewInt64Decoder returns an int64 decoder.
func NewInt64Decoder(buf []byte) *Int64Decoder {
	return &Int64Decoder{
		buf: bytes.NewBuffer(buf),
	}
}

// Next returns true if a value was scanned.
func (d *Int64Decoder) Next() bool {
	if d.Run > 1 {
		d.Run--
		return true
	}

	num, err := binary.ReadVarint(d.buf)
	if err == io.EOF {
		return false
	}

	if err != nil {
		d.err = err
		return false
	}

	run, err := binary.ReadVarint(d.buf)
	if err == io.EOF {
		d.err = io.ErrUnexpectedEOF
		return false
	}

	if err != nil {
		d.err = err
		return false
	}

	d.Value = num
	d.Run = run

	return true
}

// Err returns any error which ocurred during decoding.
func (d *Int64Decoder) Err() error {
	return d.err
}

// EncodeInt64 encoded run.
func EncodeInt64(nums []int64) []byte {
	size := len(nums)

	if size == 0 {
		return nil
	}

	var b = make([]byte, 8)
	var buf bytes.Buffer
	var cur = nums[0]
	var run int64

	for i := 0; i < size; i++ {
		num := nums[i]

		if num != cur {
			n := binary.PutVarint(b, cur)
			buf.Write(b[:n])
			n = binary.PutVarint(b, run)
			buf.Write(b[:n])
			cur = num
			run = 0
		}

		run++
	}

	n := binary.PutVarint(b, cur)
	buf.Write(b[:n])
	n = binary.PutVarint(b, run)
	buf.Write(b[:n])

	return buf.Bytes()
}

// DecodeInt64 encoded run.
func DecodeInt64(buf []byte) (v []int64, err error) {
	s := NewInt64Decoder(buf)

	for s.Next() {
		v = append(v, s.Value)
	}

	return v, s.Err()
}

// DecodeInt64Card returns a map of value cardinality.
func DecodeInt64Card(buf []byte) (v map[int64]uint64, err error) {
	d := NewInt64Decoder(buf)
	v = make(map[int64]uint64)

	for d.Next() {
		v[d.Value]++
	}

	return v, d.Err()
}
