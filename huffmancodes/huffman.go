package main

import (
	"bufio"
	"io"
	"fmt"
	"os"
	"strconv"
	"container/heap"
)

type code struct {
  nodeName string
  left, right *code
  wieght int
  index int
 }

type wieghtHeap []*code

func (h wieghtHeap) Len() int           { return len(h) }
func (h wieghtHeap) Less(i, j int) bool { return h[i].wieght < h[j].wieght }
func (h wieghtHeap) Swap(i, j int)      { 
	h[i], h[j] = h[j], h[i] 
	h[i].index = i
	h[j].index = j
}

func (h *wieghtHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*code)
	item.index = n
	*h = append(*h, item)
}

func (h *wieghtHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item

}

func readData(reader *bufio.Scanner)(*wieghtHeap) {
	
	reader.Scan()
	text := reader.Text()
	charsetSize, e := strconv.Atoi(text);
		if e != nil {
		fmt.Println(e)
		panic(e)
	}

	charHeap := make(wieghtHeap, charsetSize)
	
	line := 0
	for ; line < charsetSize; {
		reader.Scan()
		textData := reader.Text()
		
		weight, e := strconv.Atoi(textData)
		if e != nil {
			fmt.Println(e)
			panic(e)
		}
		charHeap[line] = &code{
			nodeName:  "-" + strconv.Itoa(line),
			wieght : weight,
			index:  line,
		}
		line++
	}
	heap.Init(&charHeap)
	return &charHeap;
}

func createTree(chHeap *wieghtHeap)(code){

	fmt.Println(chHeap.Len());
	for chHeap.Len() > 2 {
		var ch code
		ch.left = heap.Pop(chHeap).(*code)
		ch.right = heap.Pop(chHeap).(*code)
		ch.wieght = ch.left.wieght + ch.right.wieght
		ch.nodeName = ch.left.nodeName + ch.right.nodeName
		heap.Push(chHeap,&ch)
	}
	var ch code
	ch.left = heap.Pop(chHeap).(*code)
	ch.right = heap.Pop(chHeap).(*code)
	ch.wieght = ch.left.wieght + ch.right.wieght
	ch.nodeName = ch.left.nodeName + ch.right.nodeName
	return ch
}

func findRDepth(tree code, min *int, max *int,depth int){
	depth++
	if (*max < depth){
			*max = depth
	}
	if tree.left != nil {
		findRDepth(*tree.left,min,max,depth)
	} else {
		if (*min > depth){
			*min = depth
		}
	}
	if tree.right != nil {
		findRDepth(*tree.right,min,max,depth)
	} else {
		if (*min > depth){
			*min = depth
		}
	}
}

func printCode(wr io.Writer, hfcode string,tree code){
	if tree.left == nil && tree.right == nil {
		fmt.Fprintln(wr,tree.nodeName,"=", hfcode)
	}
	if tree.left != nil {
		printCode(wr,hfcode+"0",*tree.left)
	}
	if tree.right != nil{
		printCode(wr,hfcode+"1",*tree.right)
	}
	
}

func findDepth(tree code)(int, int){
  min := 999999
  max := 0
  depth := -1
  findRDepth(tree, &min, &max,depth)
  return min, max
}


func main() {

	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	
	file := os.Args[1]
		
	f, err := os.Open(pwd + "\\"+ file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewScanner(f)

	charHeap := readData(reader)
	tree := createTree(charHeap)
	min, max := findDepth(tree)
	fmt.Println("min depth = ", min)
	fmt.Println("max depth = ", max)
}