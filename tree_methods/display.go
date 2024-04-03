package tree_methods

import "fmt"

func (imt *IncrementalMerkleTree) Display() {
	if imt.Root == nil {
		fmt.Println("Tree is empty")
		return
	}

	queue := []*Node{imt.Root}
	level := 1
	nodesInLevel := 1
	nextLevelNodes := 0

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		fmt.Printf("id block: %d, %s ", currNode.Block.Id, currNode.Block.Hash)

		nodesInLevel--

		if currNode.Left != nil {
			queue = append(queue, currNode.Left)
			nextLevelNodes++
		}
		if currNode.Right != nil {
			queue = append(queue, currNode.Right)
			nextLevelNodes++
		}

		if nodesInLevel == 0 {
			fmt.Printf(":%d level", level)
			fmt.Println() // Друкуємо новий рядок після кожного рівня дерева
			level++
			nodesInLevel = nextLevelNodes
			nextLevelNodes = 0
		}
	}
}
