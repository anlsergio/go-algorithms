package search_test

import (
	"github.com/anlsergio/go-algorithms/search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	t.Run("item is 6, index should be 6", func(t *testing.T) {
		want := 6
		got := search.Binary(in, 7)

		assertIndex(t, want, got)
	})

	t.Run("item is 1, index should be 0", func(t *testing.T) {
		want := 0
		got := search.Binary(in, 1)

		assertIndex(t, want, got)
	})

	t.Run("item is 9, index should be 8", func(t *testing.T) {
		want := 0
		got := search.Binary(in, 1)

		assertIndex(t, want, got)
	})

	t.Run("item does not exist", func(t *testing.T) {
		want := -1
		got := search.Binary(in, 15)

		assertIndex(t, want, got)
	})
}

func assertIndex(t testing.TB, want int, got int) {
	t.Helper()

	if want != got {
		t.Errorf("want index %d, got %d", want, got)
	}
}
