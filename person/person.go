package person

import (
	"errors"
	"fmt"
	"strings"
)

//interfaces only contain functions
type Identifiable interface {
	ID() string
}

type Citizen interface {
	Country() string
	Identifiable
}

//type declaration
type socialSecurityNumber string

//constructor for the struct
func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

//implementing the identifiable interface
func (ssn socialSecurityNumber) ID() string {
	return string(ssn) //create a string from the value passed in from the constructor
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

//type declaration
type europeanUnionIdentifier struct {
	id      string
	country string
}

//constructor for the struct
func NewEuropeanUnionIdentifier(value string, country string) Citizen {
	return &europeanUnionIdentifier{id: value, country: country}
}

//implementing the identifiable interface
func (eui europeanUnionIdentifier) ID() string {
	return fmt.Sprintf("ID: %v, from: %v", eui.id, eui.country) //create a string from the value passed in from the constructor
}

func (eui europeanUnionIdentifier) Country() string {
	return eui.country
}

type Name struct {
	first string
	last  string
}

type ContactDetail struct {
	Email string
	Phone string
}

//type declaration
type TwitterHander string

//type declarations enables you to add code to it
func (th TwitterHander) RedirectUrl() string {
	cleanedHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanedHandler)
}

//lower cased properties are non-exported fields
type Person struct {
	Name
	twitterHandle TwitterHander
	Citizen
}

/* func (p *Person) ID() string {
	return "12345"
} */

//a constructor is just a function that returns an instance
func NewPerson(first string, last string, identifiable Citizen) *Person {
	return &Person{
		Name:    Name{first: first, last: last},
		Citizen: identifiable,
	}
}

func (p *Person) FirstAndLast() (string, string) {
	return p.first, p.last
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%v %v", p.first, p.last)
}

func (p *Person) GetTwitterHandle() TwitterHander {
	return p.twitterHandle
}

func (p *Person) SetTwitterHandle(handle TwitterHander) error {
	if !strings.HasPrefix(string(handle), "@") {
		return errors.New("twitter handles must begin with '@'")
	}

	p.twitterHandle = handle
	return nil //return a nil error
}
