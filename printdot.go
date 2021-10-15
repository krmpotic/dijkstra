package dijkstra

import "fmt"

func isPath(id1, id2 string, path []string) bool {
	for i := 0; i < len(path)-1; i++ {
		if path[i] == id1 && path[i+1] == id2 || path[i] == id2 && path[i+1] == id1 {
			return true
		}
	}
	return false
}

func (g *graph) PrintDot(path []string) {
	fmt.Println("graph {")
	for _, e := range g.edges {
		style := ""
		if path != nil && isPath(e.n[0].id, e.n[1].id, path) {
			style = ",style=dotted"
		}
		fmt.Printf("\t%s -- %s [label=%d%s]\n", e.n[0].id, e.n[1].id, e.D, style)

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
