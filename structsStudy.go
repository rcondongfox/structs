package main

import (
	"fmt"
)

//Aggregate data type aggs data of different types
//User defined type - underlying type is struct

func main() {
	// structsStudy1()
	// ninjalevelexercises5Q1()
	// ninjalevelexercises5Q2()
	// ninjalevelexercises5Q3()
	ninjalevelexercises5Q4()
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

	//Inner type promoted to outer type for dot notation
	fmt.Println("\nAcessing via dot notation:")
	fmt.Println("Accessing 3 values of p1:", liscenseToKill1.first, liscenseToKill1.last, liscenseToKill1.ltk)
	fmt.Println("Accessing 3 values of p2:", liscenseToKill2.first, liscenseToKill2.last, liscenseToKill2.ltk,
		"\nYou can also print the embedded \"inner\" type:", liscenseToKill1.person)

}

//Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
//first name
//last name
// favorite ice cream flavors
// Create two VALUES of TYPE person.
//Print out the values, ranging over the elements in the slice which stores the favorite flavors.
func ninjalevelexercises5Q1() {
	type person struct {
		firstName                 string
		lastName                  string
		favouriteIcecreamFlavours []string
	}

	p1 := person{
		firstName:                 "Richy",
		lastName:                  "Condon",
		favouriteIcecreamFlavours: []string{"Vanilla", "Honeycomb"},
	}

	p2 := person{
		firstName:                 "Sarah",
		lastName:                  "Lafferty",
		favouriteIcecreamFlavours: []string{"Mint", "All other flavours"},
	}

	fmt.Println("Person 1: ", p1.firstName, p1.lastName)
	fmt.Println("Ranging over values in slice of fave flavours: ")
	for _, v := range p1.favouriteIcecreamFlavours {
		fmt.Println(v)
	}

	fmt.Println("\nPerson 2: ", p2.firstName, p2.lastName)
	fmt.Println("Ranging over values in slice of fave flavours: ")
	for i, v := range p2.favouriteIcecreamFlavours {
		fmt.Println(i, v)
	}
}

// Take the code from the previous exercise,
// then store the values of type person in a map with the key of last name.
// Access each value in the map.
// Print out the values, ranging over the slice.
func ninjalevelexercises5Q2() {
	type person struct {
		firstName                 string
		lastName                  string
		favouriteIcecreamFlavours []string
	}

	p1 := person{
		firstName:                 "Richy",
		lastName:                  "Condon",
		favouriteIcecreamFlavours: []string{"Vanilla", "Honeycomb"},
	}

	p2 := person{
		firstName:                 "Sarah",
		lastName:                  "Lafferty",
		favouriteIcecreamFlavours: []string{"Mint", "All other flavours"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	//Range over the map
	fmt.Println("\nPrinting out the map by ranging:")
	for k, v := range m {
		fmt.Println("\nKey:\t", k, "\nValue:\t", v)
		fmt.Println("\nFave Icecream Flavours: ")
		for _, val := range v.favouriteIcecreamFlavours {
			fmt.Println(val)
		}
	}

	fmt.Println("--------------")

	fmt.Println("\nPrinting out the map directly by accessing values & ranging over slice:")
	for _, value := range m {
		fmt.Println("\nFirst: ", value.firstName, "\nLast: ", value.lastName)
		fmt.Println(value.firstName, "'s Fave Icecream Flavours: ")
		for _, value2 := range value.favouriteIcecreamFlavours {
			fmt.Println(value2)
		}
	}

}

// Create a new type: vehicle.
// The underlying type is a struct.
// The fields:
// doors
// color
// Create two new types: truck & sedan.
// The underlying type of each of these new types is a struct.
// Embed the “vehicle” type in both truck & sedan.
// Give truck the field “fourWheel” which will be set to bool.
// Give sedan the field “luxury” which will be set to bool.
// Using the vehicle, truck, and sedan structs:
// using a composite literal, create a value of type truck and assign values to the fields;
// using a composite literal, create a value of type sedan and assign values to the fields.
// Print out each of these values.
// Print out a single field from each of these values.

func ninjalevelexercises5Q3() {
	type vehicle struct {
		doors  int
		colour string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	truck1 := truck{
		vehicle:   vehicle{4, "black"},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{5, "Silver"},
		luxury:  false,
	}

	fmt.Println("Directly printing new composite types:")
	fmt.Println("Truck:", truck1)
	fmt.Println("Sedan:", sedan1)

	fmt.Println("\nDirectly printing single field of the values:")
	fmt.Println("Vehicle info:", truck1.vehicle)
	fmt.Println("Colour", sedan1.colour)

}

//Create and use an anonymous struct.
//1 field data type map & 1 slice
//Print data
func ninjalevelexercises5Q4() {
	newPerson := struct {
		name             string
		favouriteFood    map[string]string
		favouriteHobbies []string
	}{
		name: "James",
		favouriteFood: map[string]string{
			"Top":    "Pork Chops",
			"Second": "Spuds",
		},

		favouriteHobbies: []string{"Programming", "Football"},
	}

	fmt.Println("Print whole anon type:")
	fmt.Println(newPerson)

	fmt.Println("\nPrint individual values:")
	fmt.Println("Name:", newPerson.name)
	fmt.Println("Favourite Foods: ", newPerson.favouriteFood)
	fmt.Println("Favourite Hobbies: ", newPerson.favouriteHobbies)

}
