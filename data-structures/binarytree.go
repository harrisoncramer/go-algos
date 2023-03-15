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

func (tree *BinaryTree) DFSReverseOrder(result []int) []int {

	if tree.right != nil {
		result = tree.right.DFSReverseOrder(result)
	}

	result = append(result, tree.value)

	if tree.left != nil {
		result = tree.left.DFSReverseOrder(result)
	}

	return result

}

func (tree *BinaryTree) DFSInOrder(result []int) []int {

	if tree.left != nil {
		result = tree.left.DFSInOrder(result)
	}

	result = append(result, tree.value)

	if tree.right != nil {
		result = tree.right.DFSInOrder(result)
	}

	return result

}

func (tree *BinaryTree) DFS(v int) bool {
	if tree.value < v {
		if tree.right == nil {
			return false
		}
		return tree.right.DFS(v)
	} else if tree.value > v {
		if tree.left == nil {
			return false
		}
		return tree.left.DFS(v)
	}

	return true
}

func (tree *BinaryTree) BFS(v int) bool {
	q := Queue[*BinaryTree]{}
	q.Init()
	q.Enqueue(tree)

	val, err := q.Dequeue()

	/* Error indicates queue is empty */
	for err == nil {
		if val.value == v {
			return true
		} else {
			if val.left != nil {
				q.Enqueue(val.left)
			}
			if val.right != nil {
				q.Enqueue(val.right)
			}
		}

		val, err = q.Dequeue()
	}

	return false

}
