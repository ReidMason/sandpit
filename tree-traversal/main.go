package main

import (
	"log"
	"tree-traversal/node"
)

func main() {

	myTree := node.New("Root",
		node.New("child1", nil, nil),
		node.New("child2", nil, nil),
	)

	log.Println(myTree.Left)
}
