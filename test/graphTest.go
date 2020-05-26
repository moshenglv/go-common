package test

import (
	graph2 "go-common/structure/graph"
)

func GraphTest() {

	graph := graph2.NewGraph(8)
	graph.AddEge(0, 1)
	graph.AddEge(0, 3)
	graph.AddEge(1, 0)
	graph.AddEge(1, 2)
	graph.AddEge(1, 4)
	graph.AddEge(2, 1)
	graph.AddEge(2, 5)
	graph.AddEge(3, 0)
	graph.AddEge(3, 4)
	graph.AddEge(4, 3)
	graph.AddEge(4, 1)
	graph.AddEge(4, 5)
	graph.AddEge(4, 6)
	graph.AddEge(5, 2)
	graph.AddEge(5, 4)
	graph.AddEge(5, 7)
	graph.AddEge(6, 4)
	graph.AddEge(6, 7)
	graph.AddEge(7, 6)
	graph.AddEge(7, 5)

	//for _, v := range graph.Adj {
	//	v.Display()
	//}

	graph.BFS(0, 7)
}
