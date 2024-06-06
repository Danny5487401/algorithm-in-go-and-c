package main

import (
	"fmt"
	"strconv"

	"github.com/google/btree"
)

type Book struct {
	Code int
	Name string
}

func (b *Book) Less(item btree.Item) bool {
	if v, ok := item.(*Book); ok {
		return b.Code < v.Code
	}

	fmt.Println("Error: not valid item")
	return false
}

func testAddSearch() {
	tree := btree.New(3) //初始3阶Btree

	for i := 0; i < 100; i++ {
		//插入数据
		tree.ReplaceOrInsert(&Book{Code: i, Name: "freedom" + strconv.Itoa(i)})
	}

	//范围查找
	tree.DescendRange(&Book{Code: 50}, &Book{Code: 48}, func(a btree.Item) bool {
		item := a.(*Book)
		fmt.Println(item)
		return true
	})

	//降序遍历
	tree.Ascend(btree.ItemIterator(func(a btree.Item) bool {
		item := a.(*Book)
		fmt.Println(item)
		return true
	}))

	//升序遍历
	tree.Descend(btree.ItemIterator(func(a btree.Item) bool {
		item := a.(*Book)
		fmt.Println(item)
		return true
	}))

	// Strict Weak Ordering
	dannyBook := &Book{Code: 123, Name: "danny1"}
	JoyBook := &Book{Code: 123, Name: "Joy"}
	if !dannyBook.Less(JoyBook) && !JoyBook.Less(dannyBook) {
		fmt.Println("dannyBook == JoyBook")
	}

}

func main() {
	testAddSearch()
}
