package dijkstra

import (
	"testing"
)

func TestGetPath(t *testing.T) {
	var want = struct{ b int; p []string } { 20, []string{"N5", "N6", "N3", "N1"} }
	g := NewGraph()
	g.AddEdge("N1", "N2", 7)
	g.AddEdge("N1", "N3", 9)
	g.AddEdge("N1", "N6", 14)
	g.AddEdge("N2", "N3", 10)
	g.AddEdge("N2", "N4", 15)
	g.AddEdge("N3", "N4", 11)
	g.AddEdge("N3", "N6", 2)
	g.AddEdge("N4", "N5", 6)
	g.AddEdge("N5", "N6", 9)
	b, p := g.GetPath("N1", "N5")

	ok := true
	for i := range want.p {
		if p[i] != want.p[i] {
			ok = false
		}
	}
	if b != want.b || !ok {
		t.Fatalf(`getPath("N1", "N5") = %d, ..., want %d, ...`, b, want.b)
	}

}
