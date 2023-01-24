package main

import "fmt"

// connectedComponents returns the size of the smallest and largest
// connected components in the graph that have 2 or more nodes.
func connectedComponents(g [][]int32) []int32 {
	// Create a map to track which nodes have been visited
	visited := make(map[int32]bool)

	// Initialize the sizes of the smallest and largest components
	var smallest, largest int32 = 2 * int32(len(g)), 0

	// Iterate through the edges in the graph
	for _, e := range g {
		// If the From node has not been visited, do a depth-first search
		if !visited[e[0]] {
			// Track the size of the current component
			var size int32

			// Perform the depth-first search
			size = dfs(g, visited, e[0])

			// Update the size of the smallest component
			if size < smallest && size > 1 {
				smallest = size
			}

			// Update the size of the largest component
			if size > largest {
				largest = size
			}
		}
	}

	// Return the sizes of the smallest and largest components
	return []int32{smallest, largest}
}

// dfs performs a depth-first search from the given node
func dfs(g [][]int32, visited map[int32]bool, node int32) int32 {
	// Mark the current node as visited
	visited[node] = true

	// Initialize the size of the current component
	var size int32 = 1

	// Iterate through the edges in the graph
	for _, e := range g {
		// If the edge connects to the current node and the To node has not been visited
		if e[0] == node && !visited[e[1]] {
			// Recursively search the To node
			size += dfs(g, visited, e[1])
		}
	}

	// Return the size of the current component
	return size
}

func main() {
	// Create a sample graph
	g := [][]int32{
		{1, 6},
		{2, 7},
		{3, 8},
		{4, 9},
		{2, 6},
		{3, 7},
		{4, 8},
		{5, 9},
	}

	// Find the sizes of the smallest and largest connected components
	sizes := connectedComponents(g)
	fmt.Println("Smallest component:", sizes[0])
	fmt.Println("Largest component:", sizes[1])
}
