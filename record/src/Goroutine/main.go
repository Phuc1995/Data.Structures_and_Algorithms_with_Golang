package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("This is main")
	ch := make(chan string,5)
	go func() {
		crawl(ch)
		close(ch)
	}()
}

func crawl(ch chan string)  {
	for i := 0; i < 10; i++{
		img := fmt.Sprint("image_%v.png",i)
		ch <- img
		time.Sleep(time.Second)
		fmt.Printf("%s: crawled %d. sent to chanel ...\n", timestamp(), img)

	}
}

func analyze(ch chan string)  {
	for img := range ch {
		fmt.Printf("%s: analyze %s\n", timestamp(), img)
	}
}

func timestamp() string {
	return  time.Now().Format("05")
}