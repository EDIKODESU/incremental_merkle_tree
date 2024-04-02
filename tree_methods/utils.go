package tree_methods

import (
	"crypto/sha256"
	"encoding/hex"
)

// Оновлення шляху від доданого вузла до кореня
func updatePath(node *Node) {
	for node.Parent != nil {
		parent := node.Parent
		parent.Block.Hash = hashFunction(parent.Left.Block.Hash, parent.Right.Block.Hash)
		node = parent
	}
}

// Хеш-функція для обчислення нового хеша з двох вхідних хешів
func hashFunction(leftHash, rightHash string) string {
	data := []byte(leftHash + rightHash)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func (imt *IncrementalMerkleTree) findCurrNode(predicate func(*Node) bool) *Node {
	if imt.Root == nil {
		return nil
	}

	queue := []*Node{imt.Root}
	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		if currNode.Left != nil {
			queue = append(queue, currNode.Left)
		}
		if currNode.Right != nil {
			queue = append(queue, currNode.Right)
		}

		if predicate(currNode) {
			return currNode
		}
	}

	return nil
}

func (imt *IncrementalMerkleTree) findCurrNodeByID(id int) *Node {
	return imt.findCurrNode(func(node *Node) bool {
		return node.Block != nil && node.Block.Id == id
	})
}

func (imt *IncrementalMerkleTree) findCurrNodeByHash(hashData string) *Node {
	return imt.findCurrNode(func(node *Node) bool {
		return node.Block != nil && node.Block.Hash == hashData
	})
}
