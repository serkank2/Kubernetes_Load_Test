package main

import (
	"fmt"
	"net/http"
)

func load() {
	for {
		requestURL := "http://20.76.246.240"
		// requestURL := " http://
		res, err := http.Get(requestURL)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(res)
	}
}

func main() {

	for i := 0; i < 500; i++ {
		go load()
	}
	load()

}
