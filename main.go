package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)

	}
	for range os.Args[1:] {
		fmt.Println("#########################################################################################")
		//fmt.Println(<-ch)
		fmt.Println(<-ch)
	}

}

func fetch(url string, ch chan<- string) {

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	ch <- fmt.Sprintf("%s", b)

}
