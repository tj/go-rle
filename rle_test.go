package rle_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tj/go-rle"
)

func TestInt64(t *testing.T) {
	{
		var nums []int64
		b := rle.Int64(nums)
		v, err := rle.ParseInt64(b)
		assert.NoError(t, err)
		assert.Equal(t, nums, v)
	}

	{
		nums := []int64{1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2}
		b := rle.Int64(nums)
		v, err := rle.ParseInt64(b)
		assert.NoError(t, err)
		assert.Equal(t, nums, v)
	}

	{
		nums := []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		b := rle.Int64(nums)
		v, err := rle.ParseInt64(b)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(b), "should be two bytes")
		assert.Equal(t, nums, v)
	}
}

func TestInt64Cardinality(t *testing.T) {
	{
		nums := []int64{1, 1, 1, 1, 1, 1, 0, 0, 0, 2}
		b := rle.Int64(nums)
		v, err := rle.Int64Card(b)
		assert.NoError(t, err)
		assert.Equal(t, map[int64]uint64{1: 6, 0: 3, 2: 1}, v)
	}
}

func BenchmarkInt64(b *testing.B) {
	nums := []int64{1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2}
	for i := 0; i < b.N; i++ {
		rle.Int64(nums)
	}
}

func BenchmarkParseInt64(b *testing.B) {
	nums := []int64{1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2}
	buf := rle.Int64(nums)
	for i := 0; i < b.N; i++ {
		rle.ParseInt64(buf)
	}
}
