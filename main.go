package main

import (
	"fmt"
	"gotesting/person"
)

func main() {
	var p person.Indetifiable = &person.Person{}
	fmt.Println(p.ID()) //prints 12345

	var p2 person.Indetifiable = person.NewPerson("Juan", "Afable")
	fmt.Println(p2.ID()) //prints 12345
	//fmt.Println(p2.FirstAndLast()) //prints fails because not in interface

	var p3 *person.Person = person.NewPerson("Juan", "Afable")
	fmt.Println(p3.ID())           //prints 12345
	fmt.Println(p3.FirstAndLast()) //prints fails because not in interface

	err := p3.SetTwitterHandle("juanjuanzero")
	if err != nil {
		fmt.Printf("an error occured : %v \n", err.Error())
	}
	fmt.Printf("\n--- Proper Twitter Handle ---\n")
	p3.SetTwitterHandle("@juanjuanzero")
	fmt.Println(p3.GetTwitterHandle())
	//GetTwitterHandle returns the type definition which also has the redirect Url
	fmt.Println(p3.GetTwitterHandle().RedirectUrl())

}
