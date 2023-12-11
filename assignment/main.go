package main

import "fmt"

type randomNumbers []int

func main (){
// creating a slice of integers
numbers := randomNumbers{47,12,76,9,23,46,68,23,14,30,45}

// Looping through the numbers
 for _,num := range numbers{
	if num % 2 != 0 {
		fmt.Println(num,"is odd")
	}else{
		fmt.Println(num,"is even")
	}  
 }
}