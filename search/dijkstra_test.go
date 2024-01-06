package search_test

import (
	"github.com/anlsergio/go-algorithms/search"
	"testing"
)

func TestDijkstraSearch(t *testing.T) {
	graph := make(search.DijkstraGraph)
	// populate "start" neighbors
	graph["start"] = map[string]int{}
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	// populate "a" neighbors
	graph["a"] = map[string]int{}
	graph["a"]["end"] = 1

	// populate "b" neighbors
	graph["b"] = map[string]int{}
	graph["b"]["a"] = 3
	graph["b"]["end"] = 5

	// "end" has no neighbor assigned
	graph["end"] = map[string]int{}

	costs, parents := search.DijkstraSearch(graph)

	t.Run("final cost should be 6", func(t *testing.T) {
		want := 6
		got := costs["end"]

		if want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	})

	t.Run("parent of 'end' should be 'a'", func(t *testing.T) {
		want := "a"
		got := parents["end"]

		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("parent of 'a' should be 'b'", func(t *testing.T) {
		want := "b"
		got := parents["a"]

		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("parent of 'b' should be 'start'", func(t *testing.T) {
		want := "start"
		got := parents["b"]

		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}
