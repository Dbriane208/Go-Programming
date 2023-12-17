package main

import (
	"fmt"
	"net/http"
	"time"
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
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// instead of having the multiple print out statements
	// we can simplify the calls by using a for loop
	// for i := 0; i < len(links); i++{
	// 	fmt.Println(<-c)
	// }

	// we create an infinite for loop to iterate through the repeating links
	for l:= range c{
		//go checkLink(l,c)
		// to modify our function to use function literal so that we can use time blocking
		// we declare an argument to pass a string in the function literal inorder to passa value 
		// after every iteration
		go func (link string){
            time.Sleep(5*time.Second)
			go checkLink(link,c)
		}(l)
		// we're using the brackets to invoke the function
		// we pass the value that we received to the memory to give go routine acess to it
	}
}

// initializing a function to check through the link 
// we modify the function to accept channel of type string
func checkLink(link string,c chan string){
  _,err := http.Get(link)
  // handling the error
  if err != nil{
	fmt.Println(link,"might be down")
	// we send a message to our channel
	// c <- "Might be down I think"
	// instead of sending a fixed message we can use the link itself
	c <- link
	return
  }
  fmt.Println(link,"is up")
  // we send a message to our channel
  //  c<- "Yes it's up"
  // instead of sending a fixed message we can use the link itself
  c <- link
	  
}