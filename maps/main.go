package main

import "fmt"

func main(){
	// option 1
	// declaring a map
	colors := map[string]string{
		"red" : "ff0000",
		"green" : "00ff00",
		"blue" : "0000ff",
	}

	// option 2
	// Declaring the map
	// var colors map[string]string
	
	// Initializing the map
	// colors = map[string]string{
	// 	"red" : "ff0000",
	// }

    // option 3
	// declaring the map
	// colors := make(map[int]string)
	
    // manipulating a map
    // colors[2] = "#ffffff"
	// colors[1] = "#ff0000"
	// colors[3] = "#0000ff"

	// to delete a color we use delete function and pass the map and key to delete
	// delete(colors,3)

    // printing the colors in the map
	printMap(colors)

}
// Iterating over the map colors
func printMap(c map[string]string){
	// using a for loop
	for color,hex := range c{
	  fmt.Println("Hex code for",color,"is",hex)
	}
  }