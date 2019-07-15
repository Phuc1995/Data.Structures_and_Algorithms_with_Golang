package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"regexp"
	"unicode"
	"unicode/utf8"
)

type (
	filterFunc func(w string) bool
)

func main() {
	keyword := flag.String("keyword", "", "search by keyword")
	upperCase := flag.Bool("uppercase", false, "search upper case word")
	limit := flag.Int("limit", 2000, "limit number of words")
	url := "https://raw.githubusercontent.com/dwyl/english-words/master/words.txt"
	flag.Parse()
	switch {
	case *keyword != "":
		//dowload
		//read -> []string
		//filter
		words, err := download(url, *limit)
		filtered := filter(words, byKeyword(*keyword))
		if err != nil {
			panic(err)
		}

		printWords(filtered...)
	case *upperCase:
		words, err := download(url, 100)
		filtered := filter(words, byUpperCase)
		if err != nil {
			panic(err)
		}

		printWords(filtered...)
	default:
		flag.PrintDefaults()
		return
	}

}

func search(url string, limit int, filterFunc filterFunc) error {
	words, err := download(url, limit)
	if err != nil {
		return err
	}
	filtered := filter(words, filterFunc)
	printWords(filtered...)
	return nil
}

func download(url string, limit int) ([]string, error) {
	rs, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	scanner := bufio.NewScanner(rs.Body)
	result := make([]string, 0)
	for i := 0; scanner.Scan() && i < limit; i++ {
		result = append(result, scanner.Text())
	}
	return result, nil
}

func printWords(words ...string) {
	for _, s := range words {
		fmt.Println(s)
	}
}

func filter(words []string, filter filterFunc) []string {
	result := make([]string, 0)
	for _, w := range words {
		if filter(w) {
			result = append(result, w)
		}
	}
	return result
}

func byKeyword(kw string) filterFunc {
	return func(w string) bool {
		matched, _ := regexp.MatchString(kw, w)
		return matched
	}
}

func byUpperCase(w string) bool {
	firstRune, _ := utf8.DecodeRuneInString(w)
	return unicode.IsUpper(firstRune)
}
