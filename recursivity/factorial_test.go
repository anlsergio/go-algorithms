package recursivity_test

import (
	"fmt"
	"github.com/anlsergio/go-algorithms/recursivity"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		in   uint
		want uint
	}{
		{
			in:   3,
			want: 6,
		},
		{
			in:   0,
			want: 1,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("factorial of %d", tc.in), func(t *testing.T) {
			got := recursivity.Factorial(tc.in)
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
