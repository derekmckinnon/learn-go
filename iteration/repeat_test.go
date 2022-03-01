package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("d", 3)
	fmt.Println(repeated)
	// Output: ddd
}

const benchmarkCharacter = "a"
const benchmarkCount = 5

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(benchmarkCharacter, benchmarkCount)
	}
}

func BenchmarkStdRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat(benchmarkCharacter, benchmarkCount)
	}
}
