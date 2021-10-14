package main

import (
	"fmt"
)

type graph struct {
	nodes map[string]*node
	edges []*edge
}
type edge struct {
	n [2]*node
	D int
}
type node struct {
	id         string
	neigh      []*node
	B          int
	from       *node
	once, done bool
}

func NewGraph() *graph {
	nodes := make(map[string]*node)
	return &graph{nodes, nil}
}

func (g *graph) getEdge(id1, id2 string) *edge {
	for _, e := range g.edges {
		if e.n[0].id == id1 && e.n[1].id == id2 || e.n[1].id == id1 && e.n[0].id == id2 {
			return e
		}
	}
	return nil
}

func (g *graph) addEdge(id1, id2 string, D int) {
	if g.getEdge(id1, id2) != nil {
		return
	}

	e := new(edge)
	e.D = D
	g.edges = append(g.edges, e)

	if n, ok := g.nodes[id1]; ok {
		e.n[0] = n
	} else {
		e.n[0] = new(node)
		g.nodes[id1] = e.n[0]
		g.nodes[id1].id = id1
	}

	if n, ok := g.nodes[id2]; ok {
		e.n[1] = n
	} else {
		e.n[1] = new(node)
		g.nodes[id2] = e.n[1]
		g.nodes[id2].id = id2
	}

	e.n[0].neigh = append(e.n[0].neigh, e.n[1])
	e.n[1].neigh = append(e.n[1].neigh, e.n[0])
}

func (g *graph) getPath(start, end string) (int, []string) {
	var f []*node
	S, E := g.nodes[start], g.nodes[end]
	S.once = true
	f = append(f, S)

	for {
		var u []*node
		for _, a := range f {
			for _, b := range a.neigh {
				if b.done {
					continue
				}
				u = append(u, b) // part of the next front
				e := g.getEdge(a.id, b.id)
				if !b.once || a.B+e.D < b.B {
					b.once = true
					b.B = a.B + e.D
					b.from = a
				}
			}
			a.done = true
		}
		if E.done || u == nil {
			if !E.once {
				return 0, nil
			}
			var path []string
			for x := E; x != nil; x = x.from {
				path = append(path, x.id)
			}
			return E.B, path
		}
		f = u
	}
}

func (g *graph) printDot(path []string) {
	fmt.Println("graph {")
	for _, e := range g.edges {
		fmt.Printf("\t%s -- %s [label=%d]\n", e.n[0].id, e.n[1].id, e.D)
	}
	if path != nil {
		fmt.Printf("\tsubgraph {\n\t\t")
		for i, n := range path {
			fmt.Printf("%s", n)
			if i != len(path)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Println(" [style=filled, fillcolor=yellow]")
		fmt.Println("\t}")
	}
	fmt.Println("}")
}

func (g *graph) printRaw() {
	fmt.Println("Nodes:")
	for k, v := range g.nodes {
		fmt.Println(k, ":", v)
	}
	fmt.Println("Edges:")
	for _, v := range g.edges {
		fmt.Println(v)
	}
}

func main() {
	g := NewGraph()
	g.addEdge("N1", "N2", 7)
	g.addEdge("N1", "N3", 9)
	g.addEdge("N1", "N6", 14)
	g.addEdge("N2", "N3", 10)
	g.addEdge("N2", "N4", 15)
	g.addEdge("N3", "N4", 11)
	g.addEdge("N3", "N6", 2)
	g.addEdge("N4", "N5", 6)
	g.addEdge("N5", "N6", 9)
	_, p := g.getPath("N1", "N5")

	g.printDot(p)
}
