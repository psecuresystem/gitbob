package helpers

import (
	"fmt"
	"sync"
)

func Checkout(oid string) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go (func(*sync.WaitGroup) {
		var commit *Commit = DecodeCommit(oid);
		fmt.Println(commit)
		ReadTree(commit.Tree, ".")
		defer wg.Done()
	})(wg)
	SetRef("HEAD",oid)
	wg.Wait()
}