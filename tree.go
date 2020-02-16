package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int){
	inorder(t,ch)
	close(ch)
}

func inorder(t *Tree, ch chan int){
	if t == nil{
		return
	}
	if t.Left !=nil{
		inorder(t.Left, ch)
	}
	//fmt.Println(t.Value)
	ch<-t.Value
	if t.Right !=nil{
		inorder(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool{
	t1ch :=make(chan int)
	t2ch :=make(chan int)
	go Walk(t1,t1ch)
	go Walk(t2,t2ch)
	for i:=0; i<10;i++{
		a,o1 := <-t1ch
		b,o2 := <-t2ch
		if (o1 == true && o2 == false ) || (o1== false && o2 == true){
			fmt.Printf("one tree finished \n")
			return false
		}
		fmt.Printf("numbers are %v  -- %v \n", a,b)
		if a != b{
			return false
		}
	}
	return true
}



func main() {
	t1 := New(5)
	fmt.Println("t1", t1)
	t2 := New(5)
	fmt.Println("t2", t2)

	if Same(t1,t2){
		fmt.Println("Both tree are same")
	}else{
		fmt.Println("Both tree are different")
	}

}