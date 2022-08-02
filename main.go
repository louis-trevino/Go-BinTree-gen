package main

import (
	"fmt"

	"github.com/louis-trevino/go-bintree-gen/bintree"
)

func BinTreeDemo() {
	fmt.Printf("\n* Demo of uno.BinTree (custom binary tree) \n")
	var initFloatValues = []float64{-5, 54, 36, 27, 20, 11, 7, 5}
	fmt.Printf("Initial non-sorted numeric values: %v \n", initFloatValues)
	//var initValues []float64 = [float64](initFloatValues)
	binTree := bintree.NewBinTree(initFloatValues)
	fmt.Printf("BinTree: %v \n", binTree.ToString())
	var searchData float64 = 27
	var found *bintree.BinNode[float64] = binTree.SearchNodeData(searchData)
	fmt.Printf("Searching: %v, found: %v \n", searchData, found.Data)
}

func main() {
	BinTreeDemo()
}
