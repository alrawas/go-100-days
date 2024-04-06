package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 9)
	expected := "aaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// The benchmark showed that the strings.Repeat is way faster than doing a for loop in Repeat()
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("p", 10)
	fmt.Println(result)
	// Output: pppppppppp
}
