package sort_test

import (
	"fmt"
	"github.com/anlsergio/go-algorithms/sort"
	"reflect"
	"testing"
)

func TestQuick(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{
			in:   make([]int, 0),
			want: make([]int, 0),
		},
		{
			in:   []int{1},
			want: []int{1},
		},
		{
			in:   []int{3, 1, 2, 5, 4},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("in: %v", tc.in), func(t *testing.T) {
			got := sort.QuickSort(tc.in)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
