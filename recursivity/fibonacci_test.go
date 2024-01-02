package recursivity_test

import (
	"fmt"
	"github.com/anlsergio/go-algorithms/recursivity"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		in   uint
		want uint
	}{
		{
			in:   0,
			want: 0,
		},
		{
			in:   1,
			want: 1,
		},
		{
			in:   2,
			want: 1,
		},
		{
			in:   3,
			want: 2,
		},
		{
			in:   4,
			want: 3,
		},
		{
			in:   5,
			want: 5,
		},
		{
			in:   9,
			want: 34,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("in: %d", tc.in), func(t *testing.T) {
			got := recursivity.Fibonacci(tc.in)
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
