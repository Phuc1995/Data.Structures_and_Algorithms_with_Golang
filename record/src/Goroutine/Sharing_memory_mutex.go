package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := 0
	mutex := sync.Mutex{}
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			c++
			mutex.Unlock()
		}()
	}
	time.Sleep(time.Second)
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println(c)
}
