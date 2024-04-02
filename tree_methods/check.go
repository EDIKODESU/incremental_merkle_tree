package tree_methods

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func (imt *IncrementalMerkleTree) CheckData(data string) {
	if data == "" {
		fmt.Println("data is empty")
		return
	}

	if imt.Root == nil {
		return
	}

	byteData := []byte(data)

	hash := sha256.Sum256(byteData)
	hashStr := hex.EncodeToString(hash[:])

	targetNode := imt.findCurrNodeByHash(hashStr)
	if targetNode == nil {
		fmt.Printf("block with hash %s not exists\n", hashStr)
		return
	}
	fmt.Printf("block with hash %s exists, id: %d\n", targetNode.Block.Hash, targetNode.Block.Id)
}
