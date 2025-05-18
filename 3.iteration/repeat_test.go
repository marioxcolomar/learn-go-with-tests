package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat one", func(t *testing.T) {
		repeated := Repeat("one", 0)
		expected := "oneoneoneoneone"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat provided amount of times", func(t *testing.T) {
		repeated := Repeat("three", 3)
		expected := "threethreethree"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat correctly when a negative number is provided", func(t *testing.T) {
		repeated := Repeat("five", -1)
		expected := "fivefivefivefivefive"
		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t *testing.T, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}
