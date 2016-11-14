// Package rle implements run-length encoding for various types (just ints at the moment).
package rle

import (
	"bytes"
	"encoding/binary"
	"io"
)

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

// ParseInt64 encoded run.
func ParseInt64(buffer []byte) (v []int64, err error) {
	if len(buffer) == 0 {
		return nil, nil
	}

	var buf = bytes.NewBuffer(buffer)

	for {
		num, err := binary.ReadVarint(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		run, err := binary.ReadVarint(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		for i := 0; i < int(run); i++ {
			v = append(v, num)
		}
	}

	return v, nil
}
