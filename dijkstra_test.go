package dijkstra

import (
	"testing"
)

func TestGetPath(t *testing.T) {
	type in_edge struct {
		id1, id2 string
		D        int
	}

	var in = []struct {
		start, goal string
		edges       []in_edge
	}{
		{"N1", "N5",
			[]in_edge{
				{"N1", "N2", 7},
				{"N1", "N3", 9},
				{"N1", "N6", 14},
				{"N2", "N3", 10},
				{"N2", "N4", 15},
				{"N3", "N4", 11},
				{"N3", "N6", 2},
				{"N4", "N5", 6},
				{"N5", "N6", 9},
			}},

		{"N0", "N11",
			[]in_edge{
				{"N0", "N1", 5},
				{"N1", "N2", 9},
				{"N2", "N3", 13},
				{"N2", "N0", 12},
				{"N4", "N5", 3},
				{"N5", "N6", 10},
				{"N6", "N7", 3},
				{"N6", "N4", 2},
				{"N7", "N8", 10},
				{"N7", "N5", 4},
				{"N8", "N9", 1},
				{"N10", "N8", 3},
				{"N10", "N9", 12},
				{"N10", "N11", 12},
				{"N11", "N9", 1},
			},
		},
	}
	var want = []struct {
		b int
		p []string
	}{
		{20, []string{"N5", "N6", "N3", "N1"}},
		{0, nil},
	}

	for k := range in {
		g := NewGraph()
		for _, edge := range in[k].edges {
			g.AddEdge(edge.id1, edge.id2, edge.D)
		}
		b, p := g.GetPath(in[k].start, in[k].goal)

		ok := true
		if len(p) != len(want[k].p) {
			ok = false
		}
		for i := range want[k].p {
			if p[i] != want[k].p[i] {
				ok = false
			}
		}
		if b != want[k].b || !ok {
			t.Fatalf(`%d. getPath(%q, %q) = %d, %v, want %d, %v`, k, in[k].start, in[k].goal, b, p, want[k].b, want[k].p)
		}

		g.AddEdge("HANGING_NODE_1", "HANGING_NODE_2", 9)
		b, p = g.GetPath(in[k].edges[0].id1, "HANGING_NODE_1")
		if b != 0 || p != nil {
			t.Fatalf(`getPath("N1", "X") = %d, %v, want 0, nil`, b, p)
		}
	}

}

func TestGetEdge(t *testing.T) {
	g := NewGraph()

	want := 80085
	g.AddEdge("A", "B", want)

	e := g.getEdge("A", "B")
	if e == nil {
		t.Fatalf(`getEdge("A", "B") = nil`)
	}
	if e.D != want {
		t.Fatalf(`getEdge("A", "B") = %d, want %d`, e.D, want)
	}

	e = g.getEdge("B", "A")
	if e == nil {
		t.Fatalf(`getEdge("B", "A") = nil`)
	}
	if e.D != want {
		t.Fatalf(`getEdge("B", "A") = %d, want %d`, e.D, want)
	}
}
