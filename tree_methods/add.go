package tree_methods

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func (imt *IncrementalMerkleTree) AddData(id int, data string) error {

	if data == "" {
		fmt.Println("data is empty")
		return fmt.Errorf("data is empty")
	}

	// Перевірка, чи існує вузел з таким ідентифікатором
	if imt.findCurrNodeByID(id) != nil {
		fmt.Printf("block with id %d already exists\n", id)
		return fmt.Errorf("block with id %d already exists", id)
	}

	block := Block{Id: id, Data: []byte(data)}
	//block.Hash = fmt.Sprintf("%x", sha256.Sum256(block.Data))
	hash := sha256.Sum256(block.Data)
	block.Hash = hex.EncodeToString(hash[:])

	newNode := &Node{Block: &block}

	if imt.Root == nil {
		imt.Root = newNode
		return nil
	}

	queue := []*Node{imt.Root}
	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		if currNode.Left == nil {
			currNode.Left = newNode
			updatePath(currNode)
			return nil
		} else if currNode.Right == nil {
			currNode.Right = newNode
			updatePath(currNode)
			return nil
		}

		queue = append(queue, currNode.Left, currNode.Right)
	}
	return nil
}
