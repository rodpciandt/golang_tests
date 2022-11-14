package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	// contact contactInfo
	contactInfo
}

func main() {
	// alex := person {firstName: "Alex", lastName: "Anderson"}

	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}
	
	// jimPointer := &jim // get memory address / pointer
	// jimPointer.updateName("Jimmy") 
	
	jim.updateName("Jimmy")
	jim.print()
	
	//fmt.Printf("%+v", alex) // %+v will print like json format
	
	fmt.Println("slice")
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
	
}

// receiver of pointer, can be called by the type of the pointer
// an person could call this function,you can ommit creating a new var of &person
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}


func updateSlice(s []string) {
	s[0] = "Bye"
}