package exercise

import (
	"fmt"
	"testing"

	"golang.org/x/tour/tree"
)

func walkImpl(t *tree.Tree, c, quit chan int) {
	if t == nil {
		return
	}

	walkImpl(t.Left, c, quit)
	select {
	case c <- t.Value:
	case <-quit:
		return
	}

	walkImpl(t.Right, c, quit)

}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch, quit chan int) {
	walkImpl(t, ch, quit)
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	quit := make(chan int)
	defer close(quit)

	go Walk(t1, ch1, quit)
	go Walk(t2, ch2, quit)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func TestBinaryTree(t *testing.T) {
	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}
