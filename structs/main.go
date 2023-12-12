package main

import "fmt"

// declaring a struct
type person struct{
	firstName string
	lastName string
	contact contactInfo

}

// declaring another struct
type contactInfo struct{
	email string
	zipCode int
}

func main(){
	// declaring a variable using the first approach
	// db := person{"Daniel","Brian"}

	// declaring a variable using the second approach 
	// by specifying which name is first and last
	alex := person{firstName : "Alex",lastName: "Anderson"}

	// updating the values of a struct
	// db.firstName = "Albert"
	// db.lastName = "Einsten"

	alex.firstName = "John"
	alex.lastName = "Doe"

	// declaring another person and using an embedded struct
	elon := person{
		firstName: "Elon",
		lastName: "Musk",
		// using an embedded struct
		contact: contactInfo{
			email: "elon@x.com",
			zipCode: 01504,
		},

	}

	
	// the code works but doesn't update the name
	// elon.updateName("Kimbal")

	// correcting the code using pointers
	// we access the address of the value we want to update using ampersand
	// kimbalPointer := &elon
	// kimbalPointer.updateName("Kimbal")

	// we can remove the above two lines to use
	elon.updateName("Kimbal")
	elon.print()

}

// creating a function to update - the code is correct but 
// it does not perform its function
// func (p person) updateName(newFirstName string){
//     p.firstName = newFirstName
// }

// updating the updateName function
func (pointerToPerson *person) updateName(newFirstName string){
   // getting the value from the pointer address
   (*pointerToPerson).firstName = newFirstName
}

// creating a function to print
func (p person) print(){
	fmt.Printf("%+v",p)
}