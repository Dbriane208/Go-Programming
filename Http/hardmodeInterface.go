package main

import (
	"fmt"
	"io"
	"os"
)

// declaring a function to open a file in the terminal as an argument
func openFileInTerminal(){
	f, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	// using thr io.copy function to output the file contents
	io.Copy(os.Stdout,f)
}