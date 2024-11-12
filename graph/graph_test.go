package graph

import (
	"reflect"
	"slices"
	"testing"
)

func TestGraph_AddNode(t *testing.T) {
	type args struct {
		node        int
		connections []int
	}
	tests := []struct {
		name      string
		graph     Graph
		args      args
		wantNodes map[int][]int
		wantErr   bool
	}{
		{
			name: "Should add a node along with all connections",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{node: 9, connections: []int{1, 5}},
			wantNodes: map[int][]int{
				0: {4, 6, 8},
				1: {2, 5, 8, 9},
				2: {1, 3},
				3: {2, 7},
				4: {0, 5},
				5: {1, 4, 9},
				6: {0, 7},
				7: {3, 6},
				8: {0, 1},
				9: {1, 5},
			},
		},
		{
			name: "Should return error if the node already exists",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{node: 8, connections: []int{1, 2}},
			wantNodes: map[int][]int{
				0: {4, 6, 8},
				1: {2, 5, 8},
				2: {1, 3},
				3: {2, 7},
				4: {0, 5},
				5: {1, 4},
				6: {0, 7},
				7: {3, 6},
				8: {0, 1},
			},
			wantErr: true,
		},
		{
			name: "Should return error if a connection is to a non-existent node",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{node: 9, connections: []int{1, 10}},
			wantNodes: map[int][]int{
				0: {4, 6, 8},
				1: {2, 5, 8},
				2: {1, 3},
				3: {2, 7},
				4: {0, 5},
				5: {1, 4},
				6: {0, 7},
				7: {3, 6},
				8: {0, 1},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.graph.AddNode(tt.args.node, tt.args.connections); (err != nil) != tt.wantErr {
				t.Errorf("Graph.AddNode() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Sort before comparing
			for node := range tt.graph.nodes {
				slices.Sort(tt.graph.nodes[node])
			}
			for node := range tt.wantNodes {
				slices.Sort(tt.wantNodes[node])
			}

			if !reflect.DeepEqual(tt.graph.nodes, tt.wantNodes) {
				t.Errorf("Graph.AddNode() Graph.nodes = %v, want %v", tt.graph.nodes, tt.wantNodes)
			}
		})
	}
}

func TestGraph_RemoveNode(t *testing.T) {
	type args struct {
		targetNode int
	}
	tests := []struct {
		name      string
		graph     Graph
		args      args
		wantNodes map[int][]int
		wantErr   bool
	}{
		{
			name: "Should remove a node and all its edges",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{targetNode: 0},
			wantNodes: map[int][]int{
				1: {2, 5, 8},
				2: {1, 3},
				3: {2, 7},
				4: {5},
				5: {1, 4},
				6: {7},
				7: {3, 6},
				8: {1},
			},
			wantErr: false,
		},
		{
			name: "Should return error if the does not exist",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{targetNode: 9},
			wantNodes: map[int][]int{
				0: {4, 6, 8},
				1: {2, 5, 8},
				2: {1, 3},
				3: {2, 7},
				4: {0, 5},
				5: {1, 4},
				6: {0, 7},
				7: {3, 6},
				8: {0, 1},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.graph.RemoveNode(tt.args.targetNode); (err != nil) != tt.wantErr {
				t.Errorf("Graph.RemoveNode() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Sort before comparing
			for node := range tt.graph.nodes {
				slices.Sort(tt.graph.nodes[node])
			}
			for node := range tt.wantNodes {
				slices.Sort(tt.wantNodes[node])
			}

			if !reflect.DeepEqual(tt.graph.nodes, tt.wantNodes) {
				t.Errorf("Graph.RemoveNode() Graph.nodes = %v, want %v", tt.graph.nodes, tt.wantNodes)
			}
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name      string
		graph     Graph
		args      args
		wantNodes map[int][]int
		wantErr   bool
	}{
		{
			name: "Should add an edge between two nodes",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{a: 0, b: 2},
			wantNodes: map[int][]int{
				0: {2, 4, 6, 8},
				1: {2, 5, 8},
				2: {0, 1, 3},
				3: {2, 7},
				4: {0, 5},
				5: {1, 4},
				6: {0, 7},
				7: {3, 6},
				8: {0, 1},
			},
			wantErr: false,
		},
		{
			name: "Should return error if node a does not exist",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{a: 9, b: 0},
			wantNodes: map[int][]int{
				0: {4, 6, 8},
				1: {2, 5, 8},
				2: {1, 3},
				3: {2, 7},
				4: {0, 5},
				5: {1, 4},
				6: {0, 7},
				7: {3, 6},
				8: {0, 1},
			},
			wantErr: true,
		},
		{
			name: "Should return error if node b does not exist",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{a: 0, b: 9},
			wantNodes: map[int][]int{
				0: {4, 6, 8},
				1: {2, 5, 8},
				2: {1, 3},
				3: {2, 7},
				4: {0, 5},
				5: {1, 4},
				6: {0, 7},
				7: {3, 6},
				8: {0, 1},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.graph.AddEdge(tt.args.a, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Graph.AddEdge() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Sort before comparing
			for node := range tt.graph.nodes {
				slices.Sort(tt.graph.nodes[node])
			}
			for node := range tt.wantNodes {
				slices.Sort(tt.wantNodes[node])
			}

			if !reflect.DeepEqual(tt.graph.nodes, tt.wantNodes) {
				t.Errorf("Graph.AddEdge() Graph.nodes = %v, want %v", tt.graph.nodes, tt.wantNodes)
			}
		})
	}
}

func TestGraph_RemoveEdge(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name      string
		graph     Graph
		args      args
		wantNodes map[int][]int
		wantErr   bool
	}{
		{
			name: "Should remove the edge between two nodes",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{a: 0, b: 6},
			wantNodes: map[int][]int{
				0: {4, 8},
				1: {2, 5, 8},
				2: {1, 3},
				3: {2, 7},
				4: {0, 5},
				5: {1, 4},
				6: {7},
				7: {3, 6},
				8: {0, 1},
			},
			wantErr: false,
		},
		{
			name: "Should return error if node a does not exist",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{a: 9, b: 0},
			wantNodes: map[int][]int{
				0: {4, 6, 8},
				1: {2, 5, 8},
				2: {1, 3},
				3: {2, 7},
				4: {0, 5},
				5: {1, 4},
				6: {0, 7},
				7: {3, 6},
				8: {0, 1},
			},
			wantErr: true,
		},
		{
			name: "Should return error if node b does not exist",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{a: 0, b: 9},
			wantNodes: map[int][]int{
				0: {4, 6, 8},
				1: {2, 5, 8},
				2: {1, 3},
				3: {2, 7},
				4: {0, 5},
				5: {1, 4},
				6: {0, 7},
				7: {3, 6},
				8: {0, 1},
			},
			wantErr: true,
		},
		{
			name: "Should return error if the edge does not exist",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{a: 0, b: 2},
			wantNodes: map[int][]int{
				0: {4, 6, 8},
				1: {2, 5, 8},
				2: {1, 3},
				3: {2, 7},
				4: {0, 5},
				5: {1, 4},
				6: {0, 7},
				7: {3, 6},
				8: {0, 1},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.graph.RemoveEdge(tt.args.a, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Graph.RemoveEdge() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Sort before comparing
			for node := range tt.graph.nodes {
				slices.Sort(tt.graph.nodes[node])
			}
			for node := range tt.wantNodes {
				slices.Sort(tt.wantNodes[node])
			}

			if !reflect.DeepEqual(tt.graph.nodes, tt.wantNodes) {
				t.Errorf("Graph.RemoveEdge() Graph.nodes = %v, want %v", tt.graph.nodes, tt.wantNodes)
			}
		})
	}
}

func TestGraph_FindShortestPath(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		graph   Graph
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "Should find shortest path",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args:    args{a: 0, b: 3},
			want:    []int{0, 6, 7, 3},
			wantErr: false,
		},
		{
			name: "Should return error if starting node does not exist",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args:    args{a: 9, b: 3},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should return error if node b does not exist",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args:    args{a: 0, b: 9},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should return error if a path does not exist",
			graph: Graph{
				nodes: map[int][]int{
					0:  {4, 6, 8},
					1:  {2, 5, 8},
					2:  {1, 3},
					3:  {2, 7},
					4:  {0, 5},
					5:  {1, 4},
					6:  {0, 7},
					7:  {3, 6},
					8:  {0, 1},
					9:  {10},
					10: {9},
				},
			},
			args:    args{a: 0, b: 9},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.graph.FindShortestPath(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Graph.FindShortestPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.FindShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_PathExists(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		graph Graph
		args  args
		want  bool
	}{
		{
			name: "Should return true if a path exists",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			args: args{a: 0, b: 3},
			want: true,
		},
		{
			name: "Should return false if a does not exist",
			graph: Graph{
				nodes: map[int][]int{
					0:  {4, 6, 8},
					1:  {2, 5, 8},
					2:  {1, 3},
					3:  {2, 7},
					4:  {0, 5},
					5:  {1, 4},
					6:  {0, 7},
					7:  {3, 6},
					8:  {0, 1},
					9:  {10},
					10: {9},
				},
			},
			args: args{a: 0, b: 9},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.graph.PathExists(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Graph.PathExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_CycleExists(t *testing.T) {
	tests := []struct {
		name  string
		graph Graph
		want  bool
	}{
		{
			name: "Should return true if a cycle exists",
			graph: Graph{
				nodes: map[int][]int{
					0: {4, 6, 8},
					1: {2, 5, 8},
					2: {1, 3},
					3: {2, 7},
					4: {0, 5},
					5: {1, 4},
					6: {0, 7},
					7: {3, 6},
					8: {0, 1},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.graph.CycleExists(); got != tt.want {
				t.Errorf("Graph.CycleExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
