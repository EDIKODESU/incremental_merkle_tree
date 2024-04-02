package tree_methods

import "fmt"

func (imt *IncrementalMerkleTree) DeleteData(id int) {
	if imt.Root == nil {
		return
	}

	targetNode := imt.findCurrNodeByID(id)
	if targetNode == nil {
		fmt.Printf("block with id %d not exists\n", id)
		return
	}

	deleteNode(targetNode)
}

// Функція для видалення вузла та оновлення шляху до кореня
func deleteNode(node *Node) {
	node.Block = nil
	updatePath(node)
}
