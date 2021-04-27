package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getter(url string, c chan string) {
	res, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	c <- string(data)
}

func createFile(path string, i int, response string) {
	filename := strconv.Itoa(i) + ".txt"
	dst, err := os.Create(filepath.Join(path, filepath.Base(filename)))
	if err != nil {
		fmt.Println("Can't create file")
		log.Fatal(err)
	}
	defer dst.Close()
	_, err = dst.Write([]byte(response))
	if err != nil {
		fmt.Println("Can't write to file")
		log.Fatal(err)
	}
}

func main() {
	var c chan string = make(chan string)

	url := "https://jsonplaceholder.typicode.com/posts/"
	path := "./storage/posts/"

	for i := 1; i < 101; i++ {
		urlput := url + strconv.Itoa(i)
		go getter(urlput, c)

	}

	for i := 1; i < 101; i++ {
		response := <-c
		createFile(path, i, response)
	}
}
