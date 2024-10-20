package benchmarks

import (
	"testing"

	"golang.org/x/exp/rand"
)

const factorialValue = 20

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ世界")

func BenchmarkFactorialRecursive(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		factorialRecursive(factorialValue)
	}
}

func BenchmarkFactorialIterative(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		factorialIterative(factorialValue)
	}
}

func randSeq(size int, n int) []string {
	b := make([]rune, size)
	strs := []string{}

	for i := 0; i < n; i++ {
		for j := range b {
			b[j] = letters[rand.Intn(len(letters))]
		}
		strs = append(strs, string(b))
	}

	return strs
}

func BenchmarkStringConcat_Plus(b *testing.B) {
	b.ReportAllocs()
	testStrings := randSeq(10, 10)

	for i := 0; i < b.N; i++ {
		stringConcat(testStrings)
	}
}

func BenchmarkStringConcat_Builder(b *testing.B) {
	b.ReportAllocs()
	testStrings := randSeq(10, 10)
	for i := 0; i < b.N; i++ {
		stringConcatWithBuilder(testStrings)
	}
}

func BenchmarkSmallSliceUsage(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		smallSliceUsage()
	}
}

func BenchmarkBetterSliceUsage(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		betterSliceUsage()
	}
}
