package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		got := Sum(numbers)
		want := 55

		if got != want {
			t.Errorf("got %d, want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 1, 1}, []int{2, 2, 2})
	want := []int{3, 6}

	assertSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 1, 1}, []int{2, 2, 2}, []int{3, 3, 3})
		want := []int{2, 4, 6}

		assertSums(t, got, want)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{}, []int{})
		want := []int{0, 0, 0}

		assertSums(t, got, want)
	})

	t.Run("sum slices with no tails", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{}, []int{})
		want := []int{0, 0, 0}

		assertSums(t, got, want)
	})

	t.Run("sum different slices", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{2, 2}, []int{3, 1, 2})
		want := []int{0, 2, 3}

		assertSums(t, got, want)
	})
}

func assertSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
