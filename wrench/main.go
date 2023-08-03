package main

import (
	"fmt"
	"log"
	"github.com/gstrand99/RESTWrench/wrench/crud"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	url2 := "https://jsonplaceholder.typicode.com/posts"
	responseInfo, err := wrench.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET request:")
	responseInfo.Print()

	responseInfo, err = wrench.Delete(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DELETE request:")
	responseInfo.Print()

  type TestData struct {
    Title string `json:"title"`
    Body string `json:"body"`
  }

  data := TestData{Title: "test", Body: "this is a test post"}
  
  responseInfo, err = wrench.Post(url2, data)
  responseInfo.Print()
 
  responseInfo, err = wrench.Put(url, data)
  responseInfo.Print()

}
