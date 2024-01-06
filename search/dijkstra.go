package search

import (
	"math"
	"slices"
)

// DijkstraGraph represents a simple graph data structure for
// the Dijkstra algorithm implementation.
type DijkstraGraph map[string]map[string]int

var processedVertices []string

// DijkstraSearch is the implementation of the Dijkstra
// algorithm to find the shortest path between point A
// and B in terms of costs.
func DijkstraSearch(graph DijkstraGraph) (costs map[string]int, parents map[string]string) {
	costs = graph["start"]
	// "end's" cost is initialized as "infinite"
	// because it's unknown.
	costs["end"] = math.MaxInt

	// initialize the parents hash table
	// by assigning "start" as parent
	// of its immediate children.
	parents = map[string]string{}
	for k := range graph["start"] {
		parents[k] = "start"
	}

	vertex := getLessCostlyVertex(costs)

	// proceed until all vertices are processed.
	for vertex != "" {
		cost := costs[vertex]
		neighbors := graph[vertex]

		// iterate over all current vertices' neighbors.
		for k, v := range neighbors {
			newCost := cost + v

			// if it's cheaper getting to this neighbor
			// from the current vertex, update its cost
			// and parent accordingly.
			if costs[k] > newCost {
				costs[k] = newCost
				parents[k] = vertex
			}
		}

		// make sure the same vertex will not
		// be processed again and proceed to
		// process the next "cheaper" vertex.
		processedVertices = append(processedVertices, vertex)
		vertex = getLessCostlyVertex(costs)
	}

	return costs, parents
}

func getLessCostlyVertex(costs map[string]int) string {
	// initialize lowerCost with an infinity representative.
	lowerCost := math.MaxInt
	var lessCostlyNode string

	for vertex, cost := range costs {
		// skip if it's already processed
		if slices.Contains(processedVertices, vertex) {
			continue
		}

		// if current vertex cost is less than the current
		// lower cost it should be promoted as the new
		// less costly node
		if cost < lowerCost {
			lowerCost = cost
			lessCostlyNode = vertex
		}
	}

	return lessCostlyNode
}
