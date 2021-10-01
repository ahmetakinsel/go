package main

import "fmt"

//Define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func main() {

	//initialize person using struct
	person1 := Person{firstName: "Ahmet", lastName: "Akinsel", city: "Nicosia", gender: "m", age: 28}

	//shorthand

	//person1 := Person{"Ahmet", "Akinsel", "Nicosia", "m", 28}

	//get an info (for ex: firstname)
	fmt.Println(person1.firstName)

	person1.age++
	fmt.Println(person1)

}
