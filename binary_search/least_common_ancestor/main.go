package main

type Node struct {
	val   int32
	left  *Node
	right *Node
}

func leastCommonAncestorInBST(root *Node, n1 int32, n2 int32) int32 {
	if root == nil {
		return -1
	}

	if n1 > root.val && n2 > root.val {
		return leastCommonAncestorInBST(root.right, n1, n2)
	}

	if n1 < root.val && n2 < root.val {
		return leastCommonAncestorInBST(root.left, n1, n2)
	}

	return root.val
}
