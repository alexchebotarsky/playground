package graph

import (
	"fmt"
	"slices"
)

type Graph struct {
	nodes map[int][]int
}

// New creates an instance of Graph provided nodes, while also ensuring their
// relations are mutual and valid.
func New(nodes map[int][]int) (*Graph, error) {
	g := Graph{
		nodes: nodes,
	}

	// Check that each node's relation is mutual
	for node, neighbors := range g.nodes {
		for _, neighbor := range neighbors {
			if !slices.Contains(g.nodes[neighbor], node) {
				return nil, fmt.Errorf("provided nodes are invalid, expected node %d to contain node %d", neighbor, node)
			}
		}
	}

	return &g, nil
}

// AddNode adds a new node to the graph provided its id and connections, if
// error occurs during the process, the error is returned and no changes are
// made to the graph.
func (g *Graph) AddNode(nodeID int, connections []int) error {
	if g.nodes[nodeID] != nil {
		return fmt.Errorf("node %d already exists", nodeID)
	}

	for _, connection := range connections {
		if g.nodes[connection] == nil {
			return fmt.Errorf("unable to connect new node to non-existent node %d", connection)
		}
	}

	g.nodes[nodeID] = connections

	for _, connection := range connections {
		g.nodes[connection] = append(g.nodes[connection], nodeID)
	}

	return nil
}

// RemoveNode removes a node from the graph provided its id, only returns error
// if the node already does not exist, indicating that no removing operation was
// needed to be performed.
func (g *Graph) RemoveNode(nodeID int) error {
	if g.nodes[nodeID] == nil {
		return fmt.Errorf("node %d does not exist", nodeID)
	}

	delete(g.nodes, nodeID)

	for node := range g.nodes {
		g.nodes[node] = slices.DeleteFunc(g.nodes[node], func(neighbor int) bool {
			return neighbor == nodeID
		})
	}

	return nil
}

// AddEdge adds an edge to the graph between node `a` and node `b`, if error
// occurs during the process, the error is returned and no changes are made to
// the graph.
func (g *Graph) AddEdge(a, b int) error {
	if g.nodes[a] == nil {
		return fmt.Errorf("node %d does not exist", a)
	}
	if g.nodes[b] == nil {
		return fmt.Errorf("node %d does not exist", b)
	}

	g.nodes[a] = append(g.nodes[a], b)
	g.nodes[b] = append(g.nodes[b], a)

	return nil
}

// RemoveEdge removes an edge from the graph provided node `a` and node `b`,
// only returns error if the edge already does not exist, indicating that no
// removing operation was needed to be performed.
func (g *Graph) RemoveEdge(a, b int) error {
	if g.nodes[a] == nil {
		return fmt.Errorf("node %d does not exist", a)
	}
	if g.nodes[b] == nil {
		return fmt.Errorf("node %d does not exist", b)
	}
	if !slices.Contains(g.nodes[a], b) || !slices.Contains(g.nodes[b], a) {
		return fmt.Errorf("connection between node %d and node %d does not exist", a, b)
	}

	g.nodes[a] = slices.DeleteFunc(g.nodes[a], func(neighbor int) bool {
		return neighbor == b
	})

	g.nodes[b] = slices.DeleteFunc(g.nodes[b], func(neighbor int) bool {
		return neighbor == a
	})

	return nil
}

type LookupNode struct {
	ID     int
	Parent *LookupNode
}

// FindShortestPath returns the shortest path between node `a` and node `b`. If
// multiple same length paths exist, it will return one of them without any
// defined logic.
func (g *Graph) FindShortestPath(a, b int) ([]int, error) {
	if g.nodes[b] == nil {
		return nil, fmt.Errorf("end node %d does not exist in the graph", b)
	}

	lookupNodes := []LookupNode{
		{
			ID:     a,
			Parent: nil,
		},
	}

	pointer := 0
	for {
		// If there are no more nodes to check, we've reached the dead end
		if len(lookupNodes) <= pointer {
			return nil, fmt.Errorf("a path between node %d and node %d does not exist", a, b)
		}

		node := lookupNodes[pointer]

		neighbors, ok := g.nodes[node.ID]
		if !ok {
			return nil, fmt.Errorf("node %d does not exist in the graph", node.ID)
		}

		for _, neighbor := range neighbors {
			// This is where we came from
			if node.Parent != nil && neighbor == node.Parent.ID {
				continue
			}

			// This is a loop
			if neighbor == a {
				continue
			}

			// Reached the target
			if neighbor == b {
				path := []int{b}
				for {
					path = append(path, node.ID)

					// Reached the beginning of the path
					if node.Parent == nil {
						break
					}

					// Move to the parent node
					node = *node.Parent
				}

				slices.Reverse(path)

				return path, nil
			}

			// Add the node for the lookup
			lookupNodes = append(lookupNodes, LookupNode{
				ID:     neighbor,
				Parent: &lookupNodes[pointer],
			})
		}

		// Did not find the target yet, going to the next node
		pointer++
	}
}

// PathExists checks whether a path between node `a` and node `b` exist.
func (g *Graph) PathExists(a, b int) bool {
	path, err := g.FindShortestPath(a, b)
	return len(path) > 0 && err == nil
}

// CycleExists checks whether there is any cycle in the graph.
func (g *Graph) CycleExists() bool {
	for node := range g.nodes {
		if g.PathExists(node, node) {
			return true
		}
	}

	return false
}
