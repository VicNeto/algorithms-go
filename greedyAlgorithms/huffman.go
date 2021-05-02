package greedyAlgorithms

import (
	"container/heap"
)

type HuffmanTree struct {
	Weight float32
	Left   *HuffmanTree
	Right  *HuffmanTree
	Label  string
}

// HuffmanCode func. Creates a Tree structure
func HuffmanCode(alphabet map[string]float32) *HuffmanTree {
	// initialization
	forest := &TreesHeap{}
	heap.Init(forest)

	for lab, w := range alphabet {
		heap.Push(forest, &HuffmanTree{w, nil, nil, lab})
	}

	var t1, t2, t3 *HuffmanTree
	for forest.Len() > 1 {
		// Extract minimum frequency trees (Greedy step)
		t1 = heap.Pop(forest).(*HuffmanTree)
		t2 = heap.Pop(forest).(*HuffmanTree)
		t3 = &HuffmanTree{t1.Weight + t2.Weight, t2, t1, t2.Label + t1.Label}
		heap.Push(forest, t3)
	}

	return heap.Pop(forest).(*HuffmanTree)
}

var HuffmanPrefix map[string]string

// WriteCode func. Traverse the Huffman Tree to produce the
// Binary code
func WriteCode(huffTree *HuffmanTree, code string) {
	if huffTree.Left == nil && huffTree.Right == nil {
		HuffmanPrefix[huffTree.Label] = code
		return
	}
	if huffTree.Left != nil {
		WriteCode(huffTree.Left, code+"0")
	}
	if huffTree.Right != nil {
		WriteCode(huffTree.Right, code+"1")
	}
}

// Huffman func. Helper function to manage the algorithm func calls
func Huffman(alphabet map[string]float32) map[string]string {
	t := HuffmanCode(alphabet)
	HuffmanPrefix = make(map[string]string)
	WriteCode(t, "")

	return HuffmanPrefix
}

type TreesHeap []*HuffmanTree

func (h TreesHeap) Len() int           { return len(h) }
func (h TreesHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h TreesHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TreesHeap) Push(x interface{}) {
	*h = append(*h, x.(*HuffmanTree))
}

func (h *TreesHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
