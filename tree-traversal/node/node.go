package node

type Node struct {
	Data  string
	Left  *Node
	Right *Node
}

func New(data string, left *Node, right *Node) *Node {
	return &Node{
		Data:  data,
		Left:  left,
		Right: right,
	}
}
