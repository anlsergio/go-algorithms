package recursion_test

import (
	"fmt"
	"github.com/anlsergio/go-algorithms/recursion"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   make([]int, 0),
			want: 0,
		},
		{
			in:   []int{1},
			want: 1,
		},
		{
			in:   []int{1, 2, 3},
			want: 6,
		},
		{
			in:   []int{50, 0, 1, 5},
			want: 56,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("in: %v", tc.in), func(t *testing.T) {
			got := recursion.Sum(tc.in)
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func TestCountItems(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   make([]int, 0),
			want: 0,
		},
		{
			in:   []int{1, 2, 3},
			want: 3,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("in: %v", tc.in), func(t *testing.T) {
			got := recursion.CountItems(tc.in)
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func TestBiggestNumber(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   make([]int, 0),
			want: 0,
		},
		{
			in:   []int{5},
			want: 5,
		},
		{
			in:   []int{2, 5, 9, 1},
			want: 9,
		},
		{
			in:   []int{2, 5, 9, 15},
			want: 15,
		},
		{
			in:   []int{20, 5, 9, 15},
			want: 20,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("in: %v", tc.in), func(t *testing.T) {
			got := recursion.BiggestNumber(tc.in)
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want bool
	}{
		{
			name: "item does not exist",
			in:   10,
			want: false,
		},
		{
			name: "item exists",
			in:   7,
			want: true,
		},
	}

	iterable := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := recursion.BinarySearch(iterable, tc.in)
			if tc.want != got {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}

	t.Run("empty iterable", func(t *testing.T) {
		found := recursion.BinarySearch(make([]int, 0), 1)
		if found {
			t.Error("expected 'found' to be false")
		}
	})
}
