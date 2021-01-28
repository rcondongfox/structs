package main

import "fmt"

//Aggregate data type aggs data of different types
//User defined type - underlying type is struct

func main() {
	structsStudy1()
}

func structsStudy1() {
	//Declaring a user defined type, underlying type is a struct
	type person struct {
		first    string
		last     string
		age      int
		location string
	}

	//Embedded struct
	type liscenseToKill struct {
		person
		ltk bool
	}

	//You don't seem to need to include all the values to create an instance
	//Zero values get inserted instead
	println()
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p2 := person{
		first:    "John",
		last:     "Jaws",
		age:      32,
		location: "Silver Screen",
	}

	p3 := person{
		first: "Donny",
		last:  "Malone",
	}

	liscenseToKill1 := liscenseToKill{
		person: p1,
		ltk:    true,
	}

	liscenseToKill2 := liscenseToKill{
		person: p2,
		ltk:    true,
	}

	fmt.Println("Printing whole struct:")
	//Printing whole struct
	fmt.Println("\tSome values missing as weren't created so zero values are shown:\t", p1, p3)
	fmt.Println("\tFully initialised instance:\t", p2)

	//Accessing the values of the struct with dot notation
	fmt.Println("\nAcessing via dot notation:")
	fmt.Println("\t", p1.last, p1.first, p1.first, p1.last)
	fmt.Println("\t", p2.last, p2.first, p2.first, p2.last)

	fmt.Println("\nEmbedded Structs:")
	fmt.Println("Printing whole struct:")
	fmt.Println("\t", liscenseToKill1, liscenseToKill2)

	//Iner type promoted to outer type for dot notation
	fmt.Println("\nAcessing via dot notation:")
	fmt.Println("Accessing 3 values of p1:", liscenseToKill1.first, liscenseToKill1.last, liscenseToKill1.ltk)
	fmt.Println("Accessing 3 values of p2:", liscenseToKill2.first, liscenseToKill2.last, liscenseToKill2.ltk,
		"\nYou can also print the embedded \"inner\" type:", liscenseToKill1.person)

}
