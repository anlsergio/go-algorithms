package np_complete

// import "github.com/emirpasic/gods/sets" because up to this
// moment Go doesn't officially support Sets argh!
import (
	"github.com/emirpasic/gods/sets/hashset"
)

type Set map[string]struct{}

// GreedyIdealRadioStation implements an algorithm to find
// the closes to ideal solution to choose between an amount
// of radio stations that can cover all the states of America
// represented by the states set in an efficient way.
//
// Efficiency here means that this algorithm will try to cover
// all states with by using the minimum amount of radio stations
// as possible.
func GreedyIdealRadioStation(statesToBeCovered *hashset.Set, stations map[string]*hashset.Set) (idealStations []string) {
	// limit the iteration by the total number
	// of stations provided since it's not possible
	// to have more ideal stations than what's available.
	// and a Set doesn't allow duplicates.
	for i := 0; i < len(stations); i++ {
		// if all states are already covered there's
		// no reason to move forward.
		if statesToBeCovered.Size() == 0 {
			break
		}

		// represents the station that covers
		// more states that are not already covered
		var bestStation string
		coveredStates := hashset.New()

		// for each iteration select the station
		// that covers the bigger amount of states
		// that are not already covered.
		for station, statesByStation := range stations {
			// get the intersection of "statesToBeCovered"
			// and "statesByStation" so that we know how many
			// states in this station are states that still
			// need to be covered.
			coveredStatesCandidate := statesToBeCovered.Intersection(statesByStation)

			// if the station being checked covers
			// more states than the current best station
			// promote it as the new best station
			if coveredStatesCandidate.Size() > coveredStates.Size() {
				bestStation = station
				coveredStates = coveredStatesCandidate
			}
		}

		// it only makes sense updating references
		// if there's a new best station.
		if bestStation != "" {
			// remove states that are already covered by this station
			// to signalize they don't need to be covered by another station.
			statesToBeCovered = statesToBeCovered.Difference(coveredStates)
			// append the best station between current candidates
			idealStations = append(idealStations, bestStation)
		}
	}

	return idealStations
}
