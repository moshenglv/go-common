package queue

type Node interface{}

type Queue struct {
	nodes []Node
}

type IQueue interface {
	New() Queue
	Push(t Node)
	Pop(t Node)
	IsEmpty() bool
	Size() int
}

func New() *Queue {
	q := Queue{}
	q.nodes = []Node{}
	return &q
}

func (q *Queue) Push(node Node) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) Pop() *Node {
	if len(q.nodes) <= 0 {
		return nil
	}
	val := q.nodes[0]
	q.nodes = q.nodes[1:]
	return &val
}

// 队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}

// 队列长度
func (q *Queue) Size() int {
	return len(q.nodes)
}
