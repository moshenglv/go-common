package graph

import (
	"fmt"
	"go-common/structure/linkedList"
	"go-common/structure/queue"
)

type Graph struct {
	V   int                      //顶点个数
	Adj []*linkedList.LinkedList //邻接表
}

func NewGraph(v int) *Graph {
	adj := make([]*linkedList.LinkedList, 0)
	for i := 0; i < v; i++ {
		adj = append(adj, linkedList.NewLinkedList())
	}
	return &Graph{
		V:   v,
		Adj: adj,
	}
}

func (graph *Graph) AddEge(s, t int) {
	graph.Adj[s].Append(t)
	graph.Adj[t].Append(s)
}

/**
广度优先,s->t的最短路径
*/
func (graph *Graph) BFS(s, t int) {
	if s == t {
		return
	}
	visited := make(map[int]bool)
	prev := make([]int, 0)
	for i := 0; i < graph.V; i++ {
		prev = append(prev, -1)
	}

	queue := queue.New()
	queue.Push(s)
	visited[s] = true
	for !queue.IsEmpty() {
		w, _ := (*queue.Pop()).(int)
		for i := 0; i < graph.Adj[w].Size(); i++ {
			q := graph.Adj[w].Get(i).Value.(int)
			if !visited[q] {
				//fmt.Print(q, " ")
				prev[q] = w
				if q == t {
					fmt.Println(prev)
					print(prev, s, t)
					return
				}
				visited[q] = true
				queue.Push(q)
			}

		}
	}
}

func print(prev []int, s, t int) {
	if prev[t] != -1 && t != s {
		print(prev, s, prev[t])
	}
	fmt.Print(t, " ")
}

/**
深度优先
*/
func (graph *Graph) DFS() {

}
