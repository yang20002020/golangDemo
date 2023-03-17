package main

import (
	"fmt"
	"sync"
	"time"
)

/*
控制并发的两种方式
(1)waitgroup
(2)使用context
*/
//waitgroup
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("job 1 done")
		wg.Done()
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("job 2 Done.")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("All Done")
}
