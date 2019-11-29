package avl_tree

type avlTreeNode struct {
	value         int
	count         int
	parent        *avlTreeNode
	leftNode      *avlTreeNode
	rightNode     *avlTreeNode
	balanceFactor int
}

func initTreeNode(element int) *avlTreeNode {
	return &avlTreeNode{element, 0, nil, nil, nil, 0}
}

type AVLTree struct {
	root *avlTreeNode
}

func (avlTree *AVLTree) Init() {}

func (avlTree *AVLTree) Insert(element int) {
	if avlTree.root == nil {
		avlTree.root = initTreeNode(element)
	} else {
		avlTree.insert(avlTree.root, element)
	}
}

func (avlTree *AVLTree) insert(currentNode *avlTreeNode, element int) {
	if element == currentNode.value {
		currentNode.count++
	} else if element < currentNode.value {
		if currentNode.leftNode == nil {
			node := initTreeNode(element)
			currentNode.leftNode = node
			node.parent = currentNode
		} else {
			avlTree.insert(currentNode.leftNode, element)
		}
	} else if element > currentNode.value {
		if currentNode.rightNode == nil {
			node := initTreeNode(element)
			currentNode.rightNode = node
			node.parent = currentNode
		} else {
			avlTree.insert(currentNode.rightNode, element)
		}
	}
}

func (avlTree *AVLTree) Remove(element int) (int, bool) {
	if avlTree.root == nil {
		return 0, false
	} else {
		node := avlTree.findNode(avlTree.root, element)
		if node == nil {
			return 0, false
		} else {

		}
	}

	return 0, false
}

func (avlTree *AVLTree) Print() {
}

func (avlTree *AVLTree) balance() {
	avlTree.calculateBalanceFactor()
}

func (avlTree *AVLTree) calculateBalanceFactor() {
}

func (avlTree *AVLTree) findNode(currentNode *avlTreeNode, value int) *avlTreeNode {
	if currentNode == nil {
		return nil
	}

	switch true {
	case value == currentNode.value:
		return currentNode
	case value < currentNode.value:
		return avlTree.findNode(currentNode.leftNode, value)
	case value > currentNode.value:
		return avlTree.findNode(currentNode.rightNode, value)
	default:
		return nil
	}
}
