package avl_tree

type IAVLTree interface {
	Init()
	Insert(int)
	Delete(int) (int, bool)
	Print()
}

type IAVLTreeOperations interface {
	Init(IAVLTree)
	Insert(int)
	Delete(int) (int, bool)
	Close()
}

type AVLTreeOperation func(tree IAVLTree)
