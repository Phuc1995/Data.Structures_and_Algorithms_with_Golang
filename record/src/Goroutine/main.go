package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("This is main")
	ch := make(chan string, 5)
	go func() {
		crawl(ch)
		close(ch)
	}()
	analyze(ch)
}

func crawl(ch chan string) {
	workers := 10
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(i int) {
			img := fmt.Sprintf("image_%v.png", i)
			ch <- img
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			fmt.Printf("%s: crawled %s. sent to chanel ...\n", timestamp(), img)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func analyze(ch chan string) {
	worker := 10
	wg := sync.WaitGroup{}
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			for img := range ch {
				time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
				fmt.Printf("%s: analyze %s\n", timestamp(), img)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func timestamp() string {
	return time.Now().Format("05")
}
