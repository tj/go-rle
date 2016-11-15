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
		v, err := rle.Int64Values(b)
		assert.NoError(t, err)
		assert.Equal(t, nums, v)
	}

	{
		nums := []int64{1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2}
		b := rle.Int64(nums)
		v, err := rle.Int64Values(b)
		assert.NoError(t, err)
		assert.Equal(t, nums, v)
	}

	{
		nums := []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		b := rle.Int64(nums)
		v, err := rle.Int64Values(b)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(b), "should be two bytes")
		assert.Equal(t, nums, v)
	}
}

func TestInt64Scan(t *testing.T) {
	in := []int64{0, 0, 0, 0, 0, 3, 3, 4, 1, 1, 1}
	out := []int64{}

	s := rle.Int64Scan(rle.Int64(in))

	for s.Next() {
		out = append(out, s.Value)
	}

	assert.Equal(t, in, out)
	assert.NoError(t, s.Err())
}

func TestInt64Cardinality(t *testing.T) {
	nums := []int64{1, 1, 1, 1, 1, 1, 0, 0, 0, 2}
	b := rle.Int64(nums)
	v, err := rle.Int64Card(b)
	assert.NoError(t, err)
	assert.Equal(t, map[int64]uint64{1: 6, 0: 3, 2: 1}, v)
}

func BenchmarkInt64(b *testing.B) {
	nums100 := make([]int64, 100e3)
	nums500 := make([]int64, 500e3)
	nums1000 := make([]int64, 1e6)

	b.Run("100k", func(b *testing.B) {
		b.SetBytes(100e3 * 8)
		for i := 0; i < b.N; i++ {
			rle.Int64(nums100)
		}
	})

	b.Run("500k", func(b *testing.B) {
		b.SetBytes(500e3 * 8)
		for i := 0; i < b.N; i++ {
			rle.Int64(nums500)
		}
	})

	b.Run("1M", func(b *testing.B) {
		b.SetBytes(1e6 * 8)
		for i := 0; i < b.N; i++ {
			rle.Int64(nums1000)
		}
	})
}

func BenchmarkInt64Values(b *testing.B) {
	nums100 := rle.Int64(make([]int64, 100e3))
	nums500 := rle.Int64(make([]int64, 500e3))
	nums1000 := rle.Int64(make([]int64, 1e6))

	b.Run("100k", func(b *testing.B) {
		b.SetBytes(100e3 * 8)
		for i := 0; i < b.N; i++ {
			rle.Int64Values(nums100)
		}
	})

	b.Run("500k", func(b *testing.B) {
		b.SetBytes(500e3 * 8)
		for i := 0; i < b.N; i++ {
			rle.Int64Values(nums500)
		}
	})

	b.Run("1M", func(b *testing.B) {
		b.SetBytes(1e6 * 8)
		for i := 0; i < b.N; i++ {
			rle.Int64Values(nums1000)
		}
	})
}

func BenchmarkInt64Scan(b *testing.B) {
	nums100 := rle.Int64(make([]int64, 100e3))
	nums500 := rle.Int64(make([]int64, 500e3))
	nums1000 := rle.Int64(make([]int64, 1e6))

	b.Run("100k", func(b *testing.B) {
		b.SetBytes(100e3 * 8)
		for i := 0; i < b.N; i++ {
			s := rle.Int64Scan(nums100)
			for s.Next() {

			}
		}
	})

	b.Run("500k", func(b *testing.B) {
		b.SetBytes(500e3 * 8)
		for i := 0; i < b.N; i++ {
			s := rle.Int64Scan(nums500)
			for s.Next() {

			}
		}
	})

	b.Run("1M", func(b *testing.B) {
		b.SetBytes(1e6 * 8)
		for i := 0; i < b.N; i++ {
			s := rle.Int64Scan(nums1000)
			for s.Next() {

			}
		}
	})
}

func BenchmarkInt64Card(b *testing.B) {
	nums := []int64{1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2}
	buf := rle.Int64(nums)
	for i := 0; i < b.N; i++ {
		rle.Int64Card(buf)
	}
}
