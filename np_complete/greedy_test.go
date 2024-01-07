package np_complete_test

import (
	"github.com/anlsergio/go-algorithms/np_complete"
	"github.com/emirpasic/gods/sets/hashset"
	"slices"
	"testing"
)

func TestGreedyIdealRadioStation(t *testing.T) {
	stations := make(map[string]*hashset.Set)
	stations["kone"] = hashset.New("id", "nv", "ut")
	stations["ktwo"] = hashset.New("wa", "id", "mt")
	stations["kthree"] = hashset.New("or", "nv", "ca")
	stations["kfour"] = hashset.New("nv", "ut")
	stations["kfive"] = hashset.New("ca", "az")

	states := hashset.New("mt", "wa", "or", "id", "nv", "ut", "ca", "az")

	got := np_complete.GreedyIdealRadioStation(states, stations)
	want := []string{"ktwo", "kthree", "kone", "kfive"}

	t.Logf("got stations: %v", got)

	if len(got) != len(want) {
		t.Errorf("ideal stations length don't match. Want %d, got %d", len(want), len(got))
	}

	for _, station := range want {
		if !slices.Contains(got, station) {
			t.Errorf("expected %s station to be in %v", station, got)
		}
	}
}
