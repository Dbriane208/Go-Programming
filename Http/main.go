package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// declaring a type which will implement write interface
type logWriter struct{}

func main(){
	// Get request from the url
	resp, err := http.Get("https://google.com")

	// handling the error
	if err != nil {
		fmt.Println("Error :",err)
		os.Exit(1)
	}

	// printing the body of the response
	// we initialize a slice with 99999 spaces
	// we convert the type of the response to a string
	// bs := make([]byte,99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// we're going to condense the above code into function that aoutomatically prints the 
	// response into the terminal
	//io.Copy(os.Stdout,resp.Body)

	//we initialize the log writer function and print the response of our body
	lw := logWriter{}
	io.Copy(lw,resp.Body)

	// hard mode interfaces
	openFileInTerminal()
	
}

func (logWriter) Write(bs []byte) (int,error){
	// turning a byte slice into a string
	fmt.Println(string(bs))
	// printing the number of bytes converted to string
	fmt.Println("Just Wrote this many bytes :",len((bs)))

	//return the length of the bs and a nil error
	return len(bs),nil
}

