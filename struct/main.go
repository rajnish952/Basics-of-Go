package main

import "fmt"

// * in front of type -> means pointer to that type
// * in front of pointer -> means value at that address location

type person struct { // Create a stucture, of type person
	firstName string
	lastName  string
	contact   contactInfo // Embeded sturct
}

type contactInfo struct {
	email   string
	zipcode int
}

func main() {
	//r := person{"Rajnish", "Kumar"} // raj variable has first name = "Rajnish", lastname = "Kumar"
	ra := person{
		firstName: "Raj",
		lastName:  "Thakur",
		contact: contactInfo{
			email:   "test@gmail.com",
			zipcode: 1234,
		}}

	ra.print()
	raPointer := &ra // Reference to ra varaiable or Pointer to ra variable -> give access to mem add of the variable
	raPointer.updateFistName("Rajnish")
	ra.print()

	ra.updateFistName("R@jnish") // We can directly call a function with reciever as pointer
	ra.print()
}

func (personPointer *person) updateFistName(firstname string) {
	(*personPointer).firstName = firstname // Derefernce personPointer -> access value at the memory address

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
