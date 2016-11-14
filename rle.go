// Package rle implements run-length encoding for various types (just ints at the moment).
package rle

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Int64Scanner is what it sounds like.
type Int64Scanner struct {
	Value int64
	Run   int64
	buf   *bytes.Buffer
	err   error
}

// Next returns true if a value was scanned.
func (s *Int64Scanner) Next() bool {
	if s.Run > 1 {
		s.Run--
		return true
	}

	num, err := binary.ReadVarint(s.buf)
	if err == io.EOF {
		return false
	}

	if err != nil {
		s.err = err
		return false
	}

	run, err := binary.ReadVarint(s.buf)
	if err == io.EOF {
		s.err = io.ErrUnexpectedEOF
		return false
	}

	if err != nil {
		s.err = err
		return false
	}

	s.Value = num
	s.Run = run

	return true
}

// Err returns any error which ocurred during scanning.
func (s *Int64Scanner) Err() error {
	return s.err
}

// Int64 encoded run.
func Int64(nums []int64) []byte {
	if len(nums) == 0 {
		return nil
	}

	var b = make([]byte, 8)
	var buf bytes.Buffer
	var cur = nums[0]
	var run int64

	for _, num := range nums {
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

// Int64Scan returns an int64 scanner.
func Int64Scan(buf []byte) *Int64Scanner {
	return &Int64Scanner{
		buf: bytes.NewBuffer(buf),
	}
}

// Int64Values encoded run.
func Int64Values(buf []byte) (v []int64, err error) {
	s := Int64Scan(buf)

	for s.Next() {
		v = append(v, s.Value)
	}

	return v, s.Err()
}

// Int64Card returns a map of value cardinality.
func Int64Card(buf []byte) (v map[int64]uint64, err error) {
	s := Int64Scan(buf)
	v = make(map[int64]uint64)

	for s.Next() {
		v[s.Value]++
	}

	return v, s.Err()
}
