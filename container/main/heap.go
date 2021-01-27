package main

import (
	"container/heap"
	"fmt"
	"os"
	"path"
	"strings"
)

type heapFloat32 []float32

func (n *heapFloat32) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	newH := old[0:len(old)-1]

	*n = newH

	return x
}

func (n *heapFloat32) Push(x interface{}) {
	*n = append(*n, x.(float32))
}

func (n heapFloat32) Len() int {
	return len(n)
}

func (n heapFloat32) Less(a, b int) bool {
	return n[a] < n[b]
}

func (n heapFloat32) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

func main() {
	myHeap := &heapFloat32{1.2, 2.1, 3.1, -100.1}

	heap.Init(myHeap)

	size := len(*myHeap)

	fmt.Println("heap size", size)
	fmt.Println(myHeap)

	myHeap.Push(float32(-100.2))
	myHeap.Push(float32(0.2))

	fmt.Printf("Heap size: %d\n", len(*myHeap))
	fmt.Printf("%v\n", myHeap)

	heap.Init(myHeap)

	fmt.Printf("%v\n", myHeap)

	a := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(a)
	fmt.Println(a[:len(a)-1])

	cwd, _ := os.Getwd()

	fmt.Println(strings.Split(cwd, string(os.PathSeparator)))

	fmt.Println(path.Join(cwd, "../../"))
}
