/*
https://forum.golangbridge.org/t/how-to-use-github-com-google-btree/4086
*/

package main

import (
	"github.com/google/btree"
	"fmt"
	"flag"
)

var btreeDegree = flag.Int("degree", 3, "B-Tree degree")

func main() {
	tr := btree.New(*btreeDegree)
	for i := btree.Int(0); i < 10; i++ {
		tr.ReplaceOrInsert(i)
	}

	fmt.Println("len:       ", tr.Len())
	fmt.Println("get2:      ", tr.Get(btree.Int(2)))
	fmt.Println("get3:      ", tr.Get(btree.Int(3)))
	fmt.Println("del7:      ", tr.Delete(btree.Int(7)))
	fmt.Println("del8:      ", tr.Delete(btree.Int(8)))
	fmt.Println("replace4:  ", tr.ReplaceOrInsert(btree.Int(4)))
	fmt.Println("replace5:  ", tr.ReplaceOrInsert(btree.Int(5)))
	fmt.Println("min:       ", tr.Min())
	fmt.Println("delmin:    ", tr.DeleteMin())
	fmt.Println("max:       ", tr.Max())
	fmt.Println("delmax:    ", tr.DeleteMax())
	fmt.Println("len:       ", tr.Len())
}
