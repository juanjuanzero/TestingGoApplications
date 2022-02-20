package person

import (
	"errors"
	"fmt"
	"strings"
)

//interfaces only contain functions
type Indetifiable interface {
	ID() string
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
	firstName     string
	lastName      string
	twitterHandle TwitterHander
}

func (p *Person) ID() string {
	return "12345"
}

//a constructor is just a function that returns an instance
func NewPerson(first string, last string) *Person {
	return &Person{
		firstName: first,
		lastName:  last,
	}
}

func (p *Person) FirstAndLast() (string, string) {
	return p.firstName, p.lastName
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%v %v", p.firstName, p.lastName)
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
