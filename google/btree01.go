/*
https://forum.golangbridge.org/t/how-to-use-github-com-google-btree/4086
*/

package main

import (
	"github.com/google/btree"
	"fmt"
	"flag"
)

var btreeDegree = flag.Int("degree", 32, "B-Tree degree")

func main() {
	tr := btree.New(*btreeDegree)
	for i := btree.Int(0); i < 10; i++ {
		tr.ReplaceOrInsert(i)
	}

	fmt.Println("len:       ", tr.Len())
	fmt.Println("len:       ", tr.Len())
	fmt.Println("get3:      ", tr.Get(btree.Int(3)))
	fmt.Println("get100:    ", tr.Get(btree.Int(100)))
	fmt.Println("del4:      ", tr.Delete(btree.Int(4)))
	fmt.Println("del100:    ", tr.Delete(btree.Int(100)))
	fmt.Println("replace5:  ", tr.ReplaceOrInsert(btree.Int(5)))
	fmt.Println("replace100:", tr.ReplaceOrInsert(btree.Int(100)))
	fmt.Println("min:       ", tr.Min())
	fmt.Println("delmin:    ", tr.DeleteMin())
	fmt.Println("max:       ", tr.Max())
	fmt.Println("delmax:    ", tr.DeleteMax())
	fmt.Println("len:       ", tr.Len())
}
