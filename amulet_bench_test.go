package amulet

import (
	"math/rand"
	"testing"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func BenchmarkIsAmulet(b *testing.B) {
	input := "DON'T WORRY."
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsAmulet(input)
	}
}

func BenchmarkIsAmuletRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsAmulet(randString(20))
	}
}

func BenchmarkIsAmuletParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			IsAmulet(randString(20))
		}
	})
}
