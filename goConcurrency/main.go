package main

import (
	"fmt"
	"net/http"
)

func main(){
	// we'll start by defining a slice with links
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://netflix.com",
		"http://uber.com",
	}

	// looping through the loops
	for _,link := range links{
		checkLink(link)
	}
}

// initializing a function to check through the link
func checkLink(link string){
  _,err := http.Get(link)
  // handling the error
  if err != nil{
	fmt.Println(link,"might be down")
	return
  }
  fmt.Println(link,"is up")
}