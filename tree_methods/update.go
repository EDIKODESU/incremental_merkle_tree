package tree_methods

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func (imt *IncrementalMerkleTree) UpdateData(id int, newData string) {
	if newData == "" {
		fmt.Println("data is empty")
		return
	}

	if imt.Root == nil {
		return
	}

	targetNode := imt.findCurrNodeByID(id)
	if targetNode == nil {
		fmt.Printf("block with id %d not exists\n", id)
		return
	}
	block := targetNode.Block

	block.Data = []byte(newData)

	hash := sha256.Sum256(block.Data)
	block.Hash = hex.EncodeToString(hash[:])

	updatePath(targetNode)

}
