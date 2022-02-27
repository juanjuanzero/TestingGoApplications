package main

import (
	"fmt"
	"gotesting/person"
	"math"
)

func main() {
	/* var p person.Indetifiable = &person.Person{}
	fmt.Println(p.ID()) //prints 12345 */

	var p2 person.Identifiable = person.NewPerson("Juan", "Afable", person.NewSocialSecurityNumber("123-45-6789"))
	fmt.Println(p2.ID()) //prints 12345
	// mt.Println(p2.Country()) //this is undefined
	//fmt.Println(p2.FirstAndLast()) //prints fails because not in interface

	var p3 *person.Person = person.NewPerson("Juan", "Afable", person.NewEuropeanUnionIdentifier("123-45-6789", "Philippines"))
	fmt.Println(p3.ID())           //prints ID: 123-45-6789, from Philippines
	fmt.Println(p3.FirstAndLast()) //prints first and last
	fmt.Println(p3.Country())      //prints Philippines

	fmt.Printf("\n--- USA New Person ---\n")
	p4 := person.NewPerson("Juan", "Afable", person.NewSocialSecurityNumber("123-45-6789"))
	fmt.Println(p4.ID())           //prints 123-45-6789
	fmt.Println(p4.FirstAndLast()) //prints first and last
	fmt.Println(p4.Country())      //prints USA

	fmt.Printf("\n--- Proper Twitter Handle ---\n")
	err := p3.SetTwitterHandle("juanjuanzero")
	if err != nil {
		fmt.Printf("an error occured : %v \n", err.Error())
	}
	p3.SetTwitterHandle("@juanjuanzero")
	fmt.Println(p3.GetTwitterHandle())
	//GetTwitterHandle returns the type definition which also has the redirect Url
	fmt.Println(p3.GetTwitterHandle().RedirectUrl())

	fmt.Printf("\n--- Contact Detail Matching? ---\n")
	cd1 := person.ContactDetail{Email: "hello@world.com", Phone: "867-5309"}
	cd2 := person.ContactDetail{Email: "hello@world.com", Phone: "867-5309"}

	if cd1 == cd2 {
		fmt.Println("We match! ya!")
	}

	fmt.Printf("\n--- Anonymous Functions ---\n")
	a := func(name string) {
		fmt.Printf("hello there, %v how are you?\n", name)
	}

	a("Juan")

	fmt.Printf("\n--- Stateful Functions ---\n")
	b1 := powerOfThree()
	fmt.Printf("b1: %v\n", b1()) //first call

	b2 := powerOfThree()
	fmt.Printf("b2: %v\n", b2())       //8, different scope
	fmt.Printf("b1 again: %v\n", b1()) //27
	fmt.Printf("b1 again: %v\n", b1()) //64
}

func powerOfThree() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 3))
	}
}
