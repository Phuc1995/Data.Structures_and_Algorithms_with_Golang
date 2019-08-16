package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)
//search?q=golang&timemout=1s
func main()  {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx context.Context
			cancel context.CancelFunc
		)
		query := r.FormValue("q")
		timeout, err := time.ParseDuration(r.FormValue("timeout"))
		//ctx = context.WithValue(context.)
		if err != nil{
			ctx, cancel = context.WithCancel(context.Background())
		}else {
			ctx, cancel = context.WithTimeout(context.Background(),timeout)
		}
		defer cancel()
		v, err := searchWithTimeOut(ctx, query)
		if err != nil{
			http.Error(w, err.Error(), http.StatusNonAuthoritativeInfo)
			return
		}
		w.Write([]byte(v))

	})

	log.Fatal(http.ListenAndServe(":8088", nil))
}

func searchWithTimeOut(ctx context.Context, kw string)(string, error)  {
	select {
	case v:= <- searchFromGoogle(ctx, kw):
		return "from google" + v, nil
	case v:= <- searchFromMongo(ctx, kw):
		return "from google" + v, nil
	case v:= <- searchFromELK(ctx, kw):
		return "from google" + v, nil
	case <-ctx.Done():
		
	}
	return "", nil
}

func searchFromGoogle(ctx context.Context, kw string) <-chan string {
	return search(ctx, "google", kw, func(kw string) string {
		time.Sleep(time.Duration(rand.Intn(3))*time.Second)
		return kw
	})
}

func searchFromMongo(ctx context.Context, kw string) <-chan string {
	return search(ctx, "Mongo", kw, func(kw string) string {
		time.Sleep(time.Duration(rand.Intn(3))*time.Second)
		return kw
	})
}

func searchFromELK(ctx context.Context, kw string) <-chan string {
	return search(ctx, "ELK", kw, func(kw string) string {
		time.Sleep(time.Duration(rand.Intn(3))*time.Second)
		return kw
	})
}

func search(ctx context.Context, typ string, kw string, f func(kw string) string) <- chan string  {
	fmt.Println("searching on %s\n", typ)
	rs := make(chan string)
	ch := make(chan string)
	go func() {
		select {
		case <- ctx.Done():
			fmt.Printf("searching on %s is cancale \n", typ)
		case v := <- ch:
			rs <- v
			return
		}
	}()

	go func() {
		ch <- f(kw)
	}()
	return rs
}

