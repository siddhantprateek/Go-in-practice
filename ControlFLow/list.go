package main

// The container/list package implements a doubly-linked list. A linked list is a type of data structure that looks like this:
import (
	"container/list"
	"fmt"
)

func main() {
  var x list.List
  x.PushBack(1)
  x.PushBack(2)
  x.PushBack(3)

  for e := x.Front(); e != nil; e=e.Next() {
    fmt.Println(e.Value.(int))
  }
}