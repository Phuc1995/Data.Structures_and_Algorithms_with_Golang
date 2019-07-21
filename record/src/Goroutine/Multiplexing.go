package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	v, error := searchTimeout("test", time.Second*8)
	if error != nil {
		fmt.Println("time out")
	}
	fmt.Println(v)
}

func searchTimeout(kw string, t time.Duration) (string, error) {
	select {
	case v := <-mongo(kw):
		return v, nil
	case v := <-elec(kw):
		return v, nil
	case v := <-google(kw):
		return v, nil
	case <-time.After(t):
		return "", fmt.Errorf("search time out")

	}
	return "", nil
}

func search(kw string) string {
	select {
	case v := <-mongo(kw):
		return v
	case v := <-elec(kw):
		return v
	case v := <-google(kw):
		return v
	default:
		fmt.Println("error")
	}
	return ""
}

func mongo(kw string) chan string {
	fmt.Println("mongo search...")
	ch := make(chan string)
	go func() {
		time.Sleep(9 * time.Second)
		ch <- "mongo"
	}()
	return ch
}

func elec(kw string) chan string {
	fmt.Println("elec search...")
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		ch <- "elec"
	}()
	return ch
}

func google(kw string) chan string {
	fmt.Println("google search...")
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		ch <- "google"
	}()
	return ch
}
