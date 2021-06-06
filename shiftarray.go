package folkatech

import "fmt"

type Node struct {
	Value    int
	Position int
	Next     *Node
}

func (n Node) String() string {
	return fmt.Sprintf("Value is %v and the position is %v", n.Value, n.Position)
}

func ShiftArray(arr []int, shift int) []int {
	l := map[int]int{1: 2, 2: 3, 3: 6, 4: 1, 6: 9, 7: 4, 8: 7, 9: 8}
	nodes := []Node{}
	for i := range arr {
		n := Node{Position: i, Value: arr[i]}
		nodes = append(nodes, n)
	}
	for i := range nodes {
		if nodes[i].Value != 5 {
			nextIndex := l[nodes[i].Value] - 1
			nodes[i].Next = &nodes[nextIndex]
		}
	}

	for i := 0; i < shift; i++ {
		nexts := map[int]int{}
		for j := range nodes {
			if nodes[j].Value != 5 {
				nexts[nodes[j].Value] = nodes[j].Next.Position
			}
		}
		for j := range nodes {
			if nodes[j].Value != 5 {
				nodes[j].Position = nexts[nodes[j].Value]
			}
		}
	}
	res := make([]int, 9)
	for _, n := range nodes {
		res[n.Position] = n.Value
	}
	return res
}
