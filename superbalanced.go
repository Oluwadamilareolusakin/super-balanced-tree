package main

import (
  "fmt"
  "github.com/oluwadamilareolusakin/datastructures/tree"
  "github.com/oluwadamilareolusakin/datastructures/slice"
)

func assertTrue(value bool, desc string) {
  if value {
    fmt.Println(desc, ": PASS")
    return
  }

  fmt.Println(desc, ": FAIL")
}

func assertFalse(value bool, desc string) {
  if !value {
    fmt.Println(desc, ": PASS")
    return
  }

  fmt.Println(desc, ": FAIL")
}

type NodeWithDepth struct {
  Node *tree.Tree
  Depth int
}

func isBalanced(t *tree.Tree) (result bool) {
  result = true

  depths := []int{}

  nodes := []NodeWithDepth{{Node: t, Depth: 0}}
  depthsFound := map[int]bool{}
  var currentNodePair NodeWithDepth

  for len(nodes) > 0 {
    currentNodePair, nodes = nodes[0], nodes[1:]

    node, depth := currentNodePair.Node, currentNodePair.Depth

    if node.Left == nil && node.Right == nil && !depthsFound[depth]{
      depths = append(depths, depth)
    }

    if node.Left != nil {
      nodes = append(nodes, NodeWithDepth{Node: node.Left, Depth: depth + 1})
    }

    if node.Right != nil {
      nodes = append(nodes, NodeWithDepth{Node: node.Right, Depth: depth + 1})
    }
  }

  max := slice.Max(depths)
  min := slice.Min(depths)

  if max - min > 1 {
    result = false
  }

  return
}

func bothSubTreesSuperBalanced() {
  desc := "both subtrees superbalanced"

  tree := tree.New(1)
  tree.InsertLeft(5)
  right := tree.InsertRight(9)

  rightLeft := right.InsertLeft(8)
  right.InsertRight(5)
  rightLeft.InsertLeft(7)

  result := isBalanced(tree)

  assertFalse(result, desc)
}

func fullTree() {
  desc := "Full tree"

  tree := tree.New(5)

  left := tree.InsertLeft(8)
  left.InsertLeft(1)
  left.InsertRight(2)

  right := tree.InsertRight(6)
  right.InsertLeft(3)
  right.Left.InsertRight(4)

  result := isBalanced(tree)

  assertTrue(result, desc)
}

func differByTwo() {
  desc := "leaf heights differ by two"

  tree := tree.New(6)
  tree.InsertLeft(1)
  right := tree.InsertRight(0)

  rightRight := right.InsertRight(7)
  rightRight.InsertRight(8)

  result := isBalanced(tree)

  assertFalse(result, desc)
}

func main() {
 fullTree()
 differByTwo()
 bothSubTreesSuperBalanced()
}
