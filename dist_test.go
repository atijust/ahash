package ahash

import (
	"math/rand"
	"testing"
)

func expectedDistance(hash1 uint64, hash2 uint64) int {
	d := 0
	x := hash1 ^ hash2
	for x != 0 {
		if x&1 == 1 {
			d++
		}
		x >>= 1
	}
	return d
}

func rand64() uint64 {
	return uint64(rand.Int31())<<32 | uint64(rand.Int31())
}

func TestDistance(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		h1, h2 := rand64(), rand64()
		act, exp := Distance(h1, h2), expectedDistance(h1, h2)
		if act != exp {
			t.Errorf("0x%xと0x%xのハミング距離; 期待される値:%d 実際の値:%d", h1, h2, exp, act)
		}
	}
}

func BenchmarkDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Distance(0xffffffffffffffff, 0x1111111111111111)
	}
}
