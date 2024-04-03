package tree_methods

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func (imt *IncrementalMerkleTree) UpdateData(id int, newData string) error {
	if newData == "" {
		fmt.Println("data is empty")
		return fmt.Errorf("data is empty")
	}

	if imt.Root == nil {
		fmt.Println("root is empty")
		return fmt.Errorf("root is empty")
	}

	targetNode := imt.findCurrNodeByID(id)
	if targetNode == nil {
		fmt.Printf("block with id %d not exists\n", id)
		return fmt.Errorf("block with id %d not exists\n", id)
	}
	block := targetNode.Block
	block.Data = []byte(newData)

	hash := sha256.Sum256(block.Data)
	block.Hash = hex.EncodeToString(hash[:])

	updatePath(targetNode)
	return nil

}
