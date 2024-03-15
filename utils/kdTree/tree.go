package kdTree

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"
	"math"
)

type Instance struct {
	Coords []float64
}

type Node struct {
	inst      *Instance
	left      *Node
	right     *Node
	rangeVals [][]float64
}

type KdTree struct {
	root *Node
}

func NewKDTree() *KdTree {
	return &KdTree{}
}

func (tree *KdTree) Insert(inst *Instance) {
	tree.root = tree.insertNode(tree.root, inst, 0, nil)
}

func (tree *KdTree) insertNode(curr *Node, inst *Instance, depth int, parentRange [][]float64) *Node {
	if curr == nil {
		return &Node{inst: inst, rangeVals: tree.getRange(inst, parentRange)}
	}

	cd := depth % len(inst.Coords)
	if inst.Coords[cd] < curr.inst.Coords[cd] {
		curr.left = tree.insertNode(curr.left, inst, depth+1, curr.rangeVals)
	} else {
		curr.right = tree.insertNode(curr.right, inst, depth+1, curr.rangeVals)
	}

	return curr
}

func (tree *KdTree) PrintInorder() {
	tree.printNode(tree.root)
}

func (tree *KdTree) printNode(n *Node) {
	if n == nil {
		return
	}

	tree.printNode(n.left)
	fmt.Printf("Instance: %v, Range: %v\n", n.inst.Coords, n.rangeVals)
	tree.printNode(n.right)
}

func (tree *KdTree) getRange(inst *Instance, parentRange [][]float64) [][]float64 {
	var rangeVals [][]float64
	for i := 0; i < len(inst.Coords); i++ {
		var min, max float64
		if parentRange == nil {
			min, max = inst.Coords[i], inst.Coords[i]
		} else {
			min, max = parentRange[i][0], parentRange[i][1]
			if inst.Coords[i] < min {
				min = inst.Coords[i]
			}
			if inst.Coords[i] > max {
				max = inst.Coords[i]
			}
		}
		rangeVals = append(rangeVals, []float64{min, max})
	}
	return rangeVals
}

func (kdtree *KdTree) RangeSearch(min, max []float64) []*Instance {
	if kdtree.root == nil {
		return []*Instance{}
	}
	var results []*Instance
	rangeVals := make([][]float64, len(min))
	for i := range rangeVals {
		rangeVals[i] = []float64{min[i], max[i]}
	}
	kdtree.rangeSearchHelper(kdtree.root, rangeVals, &results)
	return results
}

func (kdtree *KdTree) rangeSearchHelper(node *Node, rangeVals [][]float64, results *[]*Instance) {
	if node == nil {
		return
	}
	if kdtree.isInRange(node.inst.Coords, rangeVals) {
		*results = append(*results, node.inst)
	}
	if node.left != nil && kdtree.isOverlapping(node.left.rangeVals, rangeVals) {
		kdtree.rangeSearchHelper(node.left, rangeVals, results)
	}
	if node.right != nil && kdtree.isOverlapping(node.right.rangeVals, rangeVals) {
		kdtree.rangeSearchHelper(node.right, rangeVals, results)
	}
}

func (kdtree *KdTree) isInRange(coords []float64, rangeVals [][]float64) bool {
	for i, rv := range rangeVals {
		if coords[i] < rv[0] || coords[i] > rv[1] {
			return false
		}
	}
	return true
}

func (kdtree *KdTree) isOverlapping(nodeRangeVals, searchRangeVals [][]float64) bool {
	for i, rv := range nodeRangeVals {
		if rv[1] < searchRangeVals[i][0] || rv[0] > searchRangeVals[i][1] {
			return false
		}
	}
	return true
}

// EncryptKdTree encrypts each node in the given KdTree using AES algorithm
func EncryptKdTree(kdtree *KdTree, key []byte) error {
	// Create the AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Encrypt each node in the KdTree
	encryptNode(kdtree.root, block)

	return nil
}

// Encrypts a node using AES algorithm
func encryptNode(node *Node, block cipher.Block) {
	// Convert the instance coordinates to []byte slice
	coordsBytes := make([]byte, len(node.inst.Coords)*8)
	for i, v := range node.inst.Coords {
		bits := math.Float64bits(v)
		binary.LittleEndian.PutUint64(coordsBytes[i*8:], bits)
	}

	// Encrypt the instance coordinates
	encryptedCoords := make([]byte, len(coordsBytes))
	block.Encrypt(encryptedCoords, coordsBytes)

	// Convert the encrypted coordinates back to []float64 slice
	encryptedFloats := make([]float64, len(node.inst.Coords))
	for i := 0; i < len(node.inst.Coords); i++ {
		bits := binary.LittleEndian.Uint64(encryptedCoords[i*8:])
		encryptedFloats[i] = math.Float64frombits(bits)
	}

	// Set the encrypted instance coordinates
	node.inst.Coords = encryptedFloats

	// Recursively encrypt the left and right child nodes, if they exist
	if node.left != nil {
		encryptNode(node.left, block)
	}
	if node.right != nil {
		encryptNode(node.right, block)
	}
}

// DecryptKdTree decrypts each node in the given KdTree using AES algorithm
func DecryptKdTree(kdtree *KdTree, key []byte) error {
	// Create the AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Decrypt each node in the KdTree
	decryptNode(kdtree.root, block)

	return nil
}

// Decrypts a node using AES algorithm
func decryptNode(node *Node, block cipher.Block) {
	// Convert the encrypted instance coordinates to []byte slice
	encryptedCoords := make([]byte, len(node.inst.Coords)*8)
	for i, v := range node.inst.Coords {
		bits := math.Float64bits(v)
		binary.LittleEndian.PutUint64(encryptedCoords[i*8:], bits)
	}

	// Decrypt the instance coordinates
	decryptedCoords := make([]byte, len(encryptedCoords))
	block.Decrypt(decryptedCoords, encryptedCoords)

	// Convert the decrypted coordinates back to []float64 slice
	decryptedFloats := make([]float64, len(node.inst.Coords))
	for i := 0; i < len(node.inst.Coords); i++ {
		bits := binary.LittleEndian.Uint64(decryptedCoords[i*8:])
		decryptedFloats[i] = math.Float64frombits(bits)
	}

	// Set the decrypted instance coordinates
	node.inst.Coords = decryptedFloats

	// Recursively decrypt the left and right child nodes, if they exist
	if node.left != nil {
		decryptNode(node.left, block)
	}
	if node.right != nil {
		decryptNode(node.right, block)
	}
}
