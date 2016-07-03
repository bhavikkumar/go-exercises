package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c := make(chan int)
	go func() {
		Walk(t1, c)
		Walk(t2, c)
		close(c)
	}()

	countMap := make(map[int]int)
	for i := range c {
		countMap[i]++
	}

	for _, value := range countMap {
		if value != 2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()
	for x := range ch {
		fmt.Println("Walk found value:",x)
	}

	fmt.Println("Are the binary trees the same?:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Are the binary trees the same?:", Same(tree.New(2), tree.New(1)))
}
