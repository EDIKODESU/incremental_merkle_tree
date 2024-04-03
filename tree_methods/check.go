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
		fmt.Println("root is empty")
		return
	}

	// Перетворення даних у байтовий масив та обчислення хешу
	byteData := []byte(data)
	hash := sha256.Sum256(byteData)
	hashStr := hex.EncodeToString(hash[:])

	// Рекурсивний пошук від кореня
	findNode := recursiveSearch(imt.Root, hashStr)
	if findNode == nil {
		fmt.Printf("block with hash %s not exists\n", hashStr)
		return
	}
	fmt.Printf("block with hash %s exists, id: %d\n", findNode.Block.Hash, findNode.Block.Id)
}

func recursiveSearch(node *Node, targetHash string) *Node {
	// Перевірка, чи поточний вузол має шукане хеш-значення
	if node.Block != nil && node.Block.Hash == targetHash {
		return node
	}

	// Перевірка, чи це внутрішній вузол
	if node.Left != nil && node.Right != nil {
		leftResult := recursiveSearch(node.Left, targetHash)
		if leftResult != nil {
			return leftResult
		}
		rightResult := recursiveSearch(node.Right, targetHash)
		if rightResult != nil {
			return rightResult
		}
	}

	return nil
}
