package datastructures

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func (n *BinaryTree) Add(v int) *BinaryTree {
	if n == nil {
		return &BinaryTree{v, nil, nil}
	} else if v <= n.value {
		n.left = n.left.Add(v)
	} else {
		n.right = n.right.Add(v)
	}
	return n
}

func (tree *BinaryTree) DepthFirstSearchPrint(result []int) []int {

	if tree.left != nil {
		result = tree.left.DepthFirstSearchPrint(result)
	}

	result = append(result, tree.value)

	if tree.right != nil {
		result = tree.right.DepthFirstSearchPrint(result)
	}

	return result

}
