package benchmarks

import (
	"testing"

	"golang.org/x/exp/rand"
)

const inputSize = 20

func BenchmarkFactorialRecursive(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		factorialRecursive(inputSize)
	}
}

func BenchmarkFactorialIterative(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		factorialIterative(inputSize)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ世界")

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
	testStrings := randSeq(10, 10)

	for i := 0; i < b.N; i++ {
		stringConcat(testStrings)
	}
}

func BenchmarkStringConcat_Builder(b *testing.B) {
	testStrings := randSeq(10, 10)
	for i := 0; i < b.N; i++ {
		stringConcatWithBuilder(testStrings)
	}
}
