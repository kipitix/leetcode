package p238

import (
	"reflect"
	"testing"
)

// Graph for optimization.
// 1[0]     2[1] 3[2]     4[3]   5[4]      6[5]       7[6]
//
//	\      /     \       /        \       /         /
//	 2[0-1]       12[2-3]          30[4-5]         /
//	       \      /                       \       /
//	       24[0-3]                         210[4-6]
//	             \_______5040[0-6]________/
type node struct {
	value      int
	minIndex   int
	maxIndex   int
	leftChild  *node
	rightChild *node
}

func (n *node) contains(index int) bool {
	return n.minIndex <= index && index <= n.maxIndex
}

func productExceptSelf(nums []int) []int {
	root := initGraph(nums, 0)
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = calcExcept(root, i)
	}
	return res
}

func initGraph(nums []int, leftIndex int) *node {
	l := len(nums)

	if l == 1 {
		return &node{
			value:      nums[0],
			minIndex:   leftIndex,
			maxIndex:   leftIndex,
			leftChild:  nil,
			rightChild: nil,
		}
	}

	halfIndex := (l + 1) / 2
	leftNode := initGraph(nums[0:halfIndex], leftIndex)
	rightNode := initGraph(nums[halfIndex:], leftIndex+halfIndex)
	thisNode := &node{
		value:      leftNode.value * rightNode.value,
		minIndex:   leftIndex,
		maxIndex:   leftIndex + l - 1,
		leftChild:  leftNode,
		rightChild: rightNode,
	}

	return thisNode
}

func calcExcept(root *node, exIndex int) int {

	if root.leftChild != nil && root.leftChild.contains(exIndex) {
		return calcExcept(root.leftChild, exIndex) * root.rightChild.value
	}

	if root.rightChild != nil && root.rightChild.contains(exIndex) {
		return calcExcept(root.rightChild, exIndex) * root.leftChild.value
	}

	if root.contains(exIndex) {
		return 1
	}

	return root.value
}

func TestProductExceptSelf(t *testing.T) {
	res := productExceptSelf([]int{1})
	if !reflect.DeepEqual(res, []int{1}) {
		t.Error("failed with 1")
	}

	res = productExceptSelf([]int{1, 2})
	if !reflect.DeepEqual(res, []int{2, 1}) {
		t.Error("failed with 2")
	}

	res = productExceptSelf([]int{1, 2, 3})
	if !reflect.DeepEqual(res, []int{6, 3, 2}) {
		t.Error("failed with 3")
	}

	res = productExceptSelf([]int{1, 2, 3, 4})
	if !reflect.DeepEqual(res, []int{24, 12, 8, 6}) {
		t.Error("failed with 4")
	}

	res = productExceptSelf([]int{1, 2, 3, 4, 5, 6, 7})
	if !reflect.DeepEqual(res, []int{5040, 2520, 1680, 1260, 1008, 840, 720}) {
		t.Error("failed with 7")
	}
}
