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
	// we create a channel to facilitate the communication between 
	// the main routine and the children routines of type string
    c := make(chan string)

	// looping through the loops
	for _,link := range links{
		// we introduce go routine by specifying the go keyword before the function
		// we connect the channel with the checkLink function
		go checkLink(link,c)
	}

	// we receive the message from the channel and log it immediately 
	fmt.Println(<-c)
}

// initializing a function to check through the link 
// we modify the function to accept channel of type string
func checkLink(link string,c chan string){
  _,err := http.Get(link)
  // handling the error
  if err != nil{
	fmt.Println(link,"might be down")
	// we send a message to our channel
	c <- "Might be down I think"
	return
  }
  fmt.Println(link,"is up")
  // we send a message to our channel
  c<- "Yes it's up"
}