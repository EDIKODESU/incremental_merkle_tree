package tree_methods

type Block struct {
	Id   int
	Data []byte
	Hash string
}

type Node struct {
	Block  *Block
	Left   *Node
	Right  *Node
	Parent *Node
}

type IncrementalMerkleTree struct {
	Root *Node
}
