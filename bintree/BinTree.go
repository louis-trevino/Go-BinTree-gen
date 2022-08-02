/*
 * BinTree.go
 * Custom Binary Tree
 *
 * Author: Louis Trevino
 * Copyright(C) Torino Consulting, 2022.
 *
 * Compiled and tested using Go version go1.18
 */
package bintree

import (
	"fmt"
)

type Scalar interface{ string | float64 | int64 }

type BinNode[T Scalar] struct {
	Data  T
	Left  *BinNode[T]
	Right *BinNode[T]
}

type BinTree[T Scalar] struct {
	Root *BinNode[T]
}

func NewBinTree[T Scalar](initialValues []T) *BinTree[T] {
	binTree := &BinTree[T]{}
	if initialValues == nil {
		return binTree
	}
	for _, iData := range initialValues {
		//fmt.Printf("Adding %v \n", iData)
		binTree.Insert(iData)
	}
	return binTree
}

func (bt *BinTree[T]) InsertNode(iNode *BinNode[T], key T) *BinNode[T] {
	if iNode == nil {
		iNode = &BinNode[T]{Data: key, Left: nil, Right: nil}
		if bt.Root == nil {
			bt.Root = iNode
		}
		return iNode
	}
	if key < iNode.Data {
		iNode.Left = bt.InsertNode(iNode.Left, key)
	} else if key > iNode.Data {
		iNode.Right = bt.InsertNode(iNode.Right, key)
	}
	return iNode
}

func (bt *BinTree[T]) Insert(data T) *BinNode[T] {
	var node *BinNode[T] = bt.InsertNode(bt.Root, data)
	return node
}

func (bt *BinTree[T]) Inorder() string {
	var sb string = ""
	bt.InorderNode(bt.Root, &sb)
	return sb
}

func (bt *BinTree[T]) ToString() string {
	return bt.Inorder()
}

func (bt *BinTree[T]) InorderNode(iNode *BinNode[T], sb *string) {
	if iNode != nil {
		bt.InorderNode(iNode.Left, sb)
		var sep string = ", "
		if len(*sb) == 0 {
			sep = ""
		}
		var fmtData string = fmt.Sprintf("%v", iNode.Data)
		*sb = fmt.Sprintf("%s%s%s", *sb, sep, fmtData)
		bt.InorderNode(iNode.Right, sb)
	}
}

func (bt *BinTree[T]) SearchNode(data T, iNode *BinNode[T]) *BinNode[T] {
	if iNode == nil {
		fmt.Println("iNode should be populated.")
		return nil
	}
	if data == iNode.Data {
		return iNode
	}
	if data < iNode.Data {
		return bt.SearchNode(data, iNode.Left)
	}
	if data > iNode.Data {
		return bt.SearchNode(data, iNode.Right)
	}
	return nil
}

func (bt *BinTree[T]) SearchNodeData(data T) *BinNode[T] {
	var binNode *BinNode[T] = bt.SearchNode(data, bt.Root)
	return binNode
}
