package search_test

import (
	"github.com/anlsergio/go-algorithms/search"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	tests := []struct {
		name    string
		myName  string
		friends search.Graph
		want    string
	}{
		{
			name:    "I have no friends (nil)",
			myName:  "myself",
			friends: nil,
			want:    "",
		},
		{
			name:    "I have no friends (empty)",
			myName:  "myself",
			friends: make(search.Graph),
			want:    "",
		},
		{
			name:   "bob is the immediate friend who sells mangoes",
			myName: "myself",
			friends: search.Graph{
				"myself": []string{"claire", "bob", "peggy"},
			},
			want: "bob",
		},
		{
			name:   "caleb is the closest person who sells mangoes",
			myName: "myself",
			friends: search.Graph{
				"myself": []string{"claire", "johnny", "peggy"},
				"johnny": []string{"tom", "clarke", "caleb"},
			},
			want: "caleb",
		},
		{
			name:   "there's no mango seller amongst friends",
			myName: "myself",
			friends: search.Graph{
				"myself": []string{"claire", "rachel", "peggy"},
			},
			want: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := search.BreadthFirstSearch(tc.myName, tc.friends)
			if tc.want != got {
				t.Errorf("want %s, got %s", tc.want, got)
			}
		})
	}
}
