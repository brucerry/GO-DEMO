// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package binarysearchtree should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// NewBst creates and returns a new SearchTreeData.
func NewBst(i int) SearchTreeData {
	return SearchTreeData{data:i}
}

// Insert inserts an int into the SearchTreeData.
// Inserts happen based on the rules of a BinarySearchTree
func (std *SearchTreeData) Insert(i int) {
    nbst := NewBst(i)
	if i <= std.data {
        if std.left == nil {
            std.left = &nbst
        } else {
        	std.left.Insert(i)
        }
    } else {
        if std.right == nil {
            std.right = &nbst
        } else {
        	std.right.Insert(i)
        }
    }
}

// MapString returns the ordered contents of SearchTreeData as a []string.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []string ["1", "3", "5", "7"].
func (std *SearchTreeData) MapString(f func(int) string) (result []string) {
	if std.left != nil {
        result = append(result, std.left.MapString(f)...)
    }
	result = append(result, f(std.data))
    if std.right != nil {
        result = append(result, std.right.MapString(f)...)
    }
	return
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (std *SearchTreeData) MapInt(f func(int) int) (result []int) {
	if std.left != nil {
        result = append(result, std.left.MapInt(f)...)
    }
	result = append(result, f(std.data))
    if std.right != nil {
        result = append(result, std.right.MapInt(f)...)
    }
	return
}
