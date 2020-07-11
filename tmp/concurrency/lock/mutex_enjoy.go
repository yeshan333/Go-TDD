/* package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lock sync.Mutex

	go func() {
		lock.Lock()

		defer lock.Unlock()

		fmt.Println("func1 get lock")

		time.Sleep(time.Second)

		fmt.Println("func release lock")
	}()

	go func() {
		lock.Lock()

		defer lock.Unlock()

		fmt.Println("func2 get lock at " + time.Now().String())

		time.Sleep(time.Second)

		fmt.Println("func2 release lock at " + time.Now().String())
	}()

	time.Sleep(time.Second * 4) // 给其它 goroutine 执行完成的机会
}
 */