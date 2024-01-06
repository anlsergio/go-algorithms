package search

import "slices"

// FriendsGraph represents a simple graph data structure for
// the BFS algorithm implementation.
type FriendsGraph map[string][]string

// BreadthFirstSearch (BFS) uses graphs to get the shortest path
// between two objects, not necessarily the faster (see the Dijkstra Algorithm).
//
// In this example, we are searching for the closest friend or friend
// of a friend of mine who is a mango seller.
//
// Whenever the closest friend who is a mango seller is found his name will be returned.
// If none is found, it will return an empty string.
//
// Note: the mango seller is anyone whose name ends with the letter "b".
//
// Execution time: O(V+E) (total number of Vertices + Edges)
func BreadthFirstSearch(myName string, friends FriendsGraph) string {
	var searchQueue []string
	// initialize the queue if all my immediate friends
	searchQueue = append(searchQueue, friends[myName]...)

	// so that we don't check the same person twice
	// in case it's a friend in common
	var alreadyChecked []string

	for len(searchQueue) > 0 {
		// get the first person from the queue
		person := searchQueue[0]
		// since there's no corresponding "pop" function,
		// the item corresponding to the index must be deleted manually.
		searchQueue = searchQueue[1:]

		if !slices.Contains(alreadyChecked, person) {
			if isMangoSeller(person) {
				return person
			}
			// if the person is not a seller
			// append all the immediate friends of person
			// to the end of the queue.
			searchQueue = append(searchQueue, friends[person]...)
			alreadyChecked = append(alreadyChecked, person)
		}
	}

	return ""
}

func isMangoSeller(person string) bool {
	return person[len(person)-1] == 'b'
}
