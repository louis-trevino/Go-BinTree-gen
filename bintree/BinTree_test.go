package bintree

import (
	"testing"
)

var (
	initValues []float64 = []float64{-5, 54, 36, 27, 20, 11, 7, 5}
)

// func init() {
// 	initValues = []float64{-5, 54, 36, 27, 20, 11, 7, 5}
// }

func TestNewBinTree(t *testing.T) {
	bt := NewBinTree[float64](initValues)
	if bt.Root == nil {
		t.Error("Bin Tree root should be set.")
	}
}

func TestSearchNode(t *testing.T) {
	bt := NewBinTree[float64](initValues)
	var scalar float64 = 27
	bn := bt.SearchNode(scalar, bt.Root)
	if bn == nil {
		t.Errorf("Scalar %v should have been found.", scalar)
	}
}

func TestSearchNodeData(t *testing.T) {
	bt := NewBinTree[float64](initValues)
	var scalar float64 = 27
	bn := bt.SearchNodeData(scalar)
	if bn == nil {
		t.Errorf("Scalar %v should have been found.", scalar)
	}
}

func TestSearchNodeDataNotInList(t *testing.T) {
	bt := NewBinTree[float64](initValues)
	var scalar float64 = 2799
	bn := bt.SearchNodeData(scalar)
	if bn != nil {
		t.Errorf("Scalar %v should have NOT been found.", scalar)
	}
}
