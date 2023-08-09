package main

import (
	"log"
	"tree-traversal/node"
)

func main() {

	myTree := node.New("Root",
		node.New("child1",
			node.New("Another child", nil, nil),
			node.New("More children", nil, nil),
		),
		node.New("child2", nil, nil),
	)

	preOrderSearch(myTree)

}

func walk(current *node.Node, trail []*node.Node) {
	if current == nil {
		return
	}

	trail = append(trail, current)

	log.Printf("Visiting: %s", current.Data)
	log.Printf("Trail:")
	for _, x := range trail {
		log.Printf("  %s", x.Data)
	}

	walk(current.Left, trail)
	walk(current.Right, trail)
}

func preOrderSearch(head *node.Node) {
	var trail []*node.Node
	walk(head, trail)
}
