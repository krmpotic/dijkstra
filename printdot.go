package dijkstra

import "fmt"

func (g *graph) PrintDot(path []string) {
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
