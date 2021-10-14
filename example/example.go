package main

import "github.com/krmpotic/dijkstra"

func main() {
	g := dijkstra.NewGraph()
	g.AddEdge("N1", "N2", 7)
	g.AddEdge("N1", "N3", 9)
	g.AddEdge("N1", "N6", 14)
	g.AddEdge("N2", "N3", 10)
	g.AddEdge("N2", "N4", 15)
	g.AddEdge("N3", "N4", 11)
	g.AddEdge("N3", "N6", 2)
	g.AddEdge("N4", "N5", 6)
	g.AddEdge("N5", "N6", 9)
	_, p := g.GetPath("N1", "N5")
	g.PrintDot(p)
}
