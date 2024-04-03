package tree_methods

import "fmt"

func (imt *IncrementalMerkleTree) DeleteData(id int) error {
	if imt.Root == nil {
		fmt.Println("root is empty")
		return fmt.Errorf("data is empty")
	}

	targetNode := imt.findCurrNodeByID(id)
	if targetNode == nil {
		fmt.Printf("block with id %d not exists\n", id)
		return fmt.Errorf("block with id %d not exists\n", id)
	}

	deleteNode(targetNode)
	return nil
}

// Функція для видалення вузла та оновлення шляху до кореня
func deleteNode(node *Node) {
	parent := node.Parent
	if parent == nil {
		node.Block = nil
		return
	}

	if parent.Left == node {
		parent.Left = nil
	} else if parent.Right == node {
		parent.Right = nil
	}

	updatePath(parent)
}
