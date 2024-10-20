package benchmarks

import (
	"strings"
)

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func stringConcat(strs []string) string {

	outputStr := ""
	for _, str := range strs {
		outputStr += str
	}

	return outputStr
}

func stringConcatWithBuilder(strs []string) string {

	var b strings.Builder
	for _, str := range strs {
		b.WriteString(str)
	}

	return b.String()
}
