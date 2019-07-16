package main

import (
	"fmt"
	"io"
	"os"
)

type (
	Logger interface {
		Println(msg string)
		Printf(format string, args ...interface{})
	}

	Service struct {
		logger Logger
		repository interface{
			Save(string) bool
		}
		service int

	}

	simpleLogger struct {
		w io.Writer
	}
)

func main()  {
	f, _ := os.Create("logger.log")
	loger := simpleLogger{
		w: f,
	}
	loger.Println("test")
}

func (l simpleLogger) Println(msg string){
	fmt.Fprint(l.w, msg)
}

func (l simpleLogger) Printf(format string, args ...interface{}){
	fmt.Fprint(l.w, format)
}