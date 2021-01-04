package main

import "fmt"

func main()  {
	var firstName string ="fadhlan"
	var lastName string
	var first, second, third string
	name := new(string)

	first,second,third = "mentari", "enggar", "rizki"
	lastName= "hawali"

	fmt.Printf("%s %s %s adalah pacarnya %s %s\n", first,second,third, firstName,lastName)
	fmt.Println(name)
}
