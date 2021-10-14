package dijkstra

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

func (g *graph) AddEdge(id1, id2 string, D int) {
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

func (g *graph) GetPath(start, end string) (int, []string) {
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

