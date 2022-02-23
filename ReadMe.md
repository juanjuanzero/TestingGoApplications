# Learning Interfaces and Structs

Working with interface and struct

- in go, an interface is just something that holds a set of functions, structs hold private and public fields
- when you have Pointer based receivers, the method that you work in, works on the object stored in memory (dereference) a pointer to that actuall object as opposed to value based functions.

You see that the functions in the Person struct type have an \* next to the struct name Person. This is signifies that this function is pointing to the object and it works on that object instance.

```Go
//interfaces only contain functions
type Identifiable interface {
	ID() string
}

//lower cased properties are non-exported fields
type Person struct {
	firstName     string
	lastName      string
	twitterHandle string
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

func (p *Person) GetTwitterHandle() string {
	return p.twitterHandle
}

func (p *Person) SetTwitterHandle(handle string) error {
	if !strings.HasPrefix(handle, "@") {
		return errors.New("twitter handles must begin with '@'")
	}

	p.twitterHandle = handle
	return nil //return a nil error
}
```

# Working with Type Aliases & Declarations

Type Aliases are references to another type. This copies the fields and the method sets. You alias a type by writing `type NewTypeAlias = OldType`

If you want to use type declarations, it copies the fields of an object over to another new type. You declre type by writing `type NewTypeDeclare OldType`

Here we will add a new type and call it TwitterHandler in person.go

```Go
//type declaration
type TwitterHander string

//type declarations enables you to add code to it
func (th TwitterHander) RedirectUrl() string {
	cleanedHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanedHandler)
}
```

# Embeding Types

You can embed structs into other structs. That effectively takes the memory layout of the struct and embed it into the other struct as well. To see this in action, we'll go ahead update the code to create a new struct called Name and move the FirstAndLast() call as a part of that name.

```Go
type Name struct {
	first string
	last string
}

//type declaration
type TwitterHander string

//type declarations enables you to add code to it
func (th TwitterHander) RedirectUrl() string {
	cleanedHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanedHandler)
}

//Update Person to have a name type
type Person struct {
	Name
	twitterHandle TwitterHander
}

//a constructor is just a function that returns an instance
func NewPerson(first string, last string) *Person {
	return &Person{
		Name: Name{ first : first, last: last},
	}
}

//here you can call first and last properties of the Name field without calling .Name first
func (p *Person) FirstAndLast() (string, string) {
	return p.first, p.last //could also be p.Name.first , p.Name.last
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%v %v", p.first, p.last) //could also be p.Name.first , p.Name.last
}
```

See how you didn't have to work with layers like abstractions like abstract classes and things like that.

## Embeding Interfaces

You can embed interfaces as well. The _method set_ on the class is what gets embeded into the class. We'll go ahead and implement another identifiable class that we can use.

```Go
//type declaration
type socialSecurityNumber string
//constructor for the struct
func NewSocialSecurityNumber(value string) Identifiable {
	return socialSecurityNumber(value)
}
//implementing the identifiable interface
func (ssn socialSecurityNumber) ID() string {
	return string(ssn) //create a string from the value passed in from the constructor
}

//a constructor is just a function that returns an instance, add it to the constructor
func NewPerson(first string, last string, identifiable Identifiable) *Person {
	return &Person{
		Name:         Name{first: first, last: last},
		Identifiable: identifiable,
	}
}

//adding identifiable to Person
type Person struct {
	Name
	twitterHandle TwitterHander
	Identifiable
}

//comment this out
//func (p *Person) ID() string {
//	return "12345"
//}
```

So here we embeded an interface onto a struct, but on instantiation we constructed an implementation of that interface `socialSecurityNumber` which is the type declaration that implements Identifiable

We can extend this even more by creating another struct, like a country method. We'll create another struct here is the new interface and the new struc
.

```GO
type Citizen interface {
	Country() string
	Identifiable
}

//new struct
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
```

As you can see the new struct retuns an object that implements the Citizen interface which has the ID() and the Citizen(). The instances now implement the Citizen interface independently, and meanwhile the Citizen interface embeds the identifieable interface, which just implements ID.

Conflicts, you can still have the same field name types, the compiler will stop you from doing that. If you have a conflict the compiler will tell you that you have an ambiguous reference.

# Comparability and Equality

As long as your go type has a predictable memory layout you should be good to compare.

- The compiler cant protect you from comparing two types that have a different memory layout. What would not have an different memory layout? These could be functions, slices or maps that grows in its memory layout.

```Go
// person.go
type ContactDetail struct {
	Email string
	Phone string
}
```

```Go
// main.go
fmt.Printf("\n--- Contact Detail Matching? ---\n")
cd1 := person.ContactDetail{Email: "hello@world.com", Phone: "867-5309"}
cd2 := person.ContactDetail{Email: "hello@world.com", Phone: "867-5309"}

if cd1 == cd2 {
    fmt.Println("We match! ya!")
}
```

Here we have a simple struct called ContactDetail, it has an Email and Phone fields and in main.go we instantiate two instances cd1 and cd2. We check if they are equal and when the code is executed we get the message confirming that they do match. Go takes the struct with a consistent memory layout of Contact Detail and compares the two, field for field. You could get an in consistent memory layout when you have data structures that are mutable, like maps and slices. Also functions too!

- If you have a consistent memory layout in Go you they are hashable, which means that you can use a key
- Equals method, would be like implemting how your equality is implemented.

# Go Functions

You can attach methods to structs in Go, these behave like methods. Just like the ID() is a method on the Person struct. By default, the behavior is value based, Go creates a copy of the instance and returns that copy. You can change this by declaring the dereference pointer on the method with a `*`. This makes it a pointer based method receiver.

## Value Receivers

This is the the behavoir of most of the functions that you write in go. When a function returns something it returns the value of that return statement, not necessarily the object stored in a particular memory location.

## Pointer Based Receivers

This is like attaching methods in golang. When you have a pointer based receiver function, this function takes a pointer to the object that is stored in memory handles works with that object to either mutate it.

### When to use pointers?

I think its best to use pointers when you are going to work with an object instance and expect to return it in your function. Value based receivers create a copy of the object passed in and return it (maybe), so it reduces that work.

## Variadric Functions

If you want to spread parameters into a variadic function all you have to do is spread it.

# Functions as values, and anonymouse functions

Functions are first class citizens in Golang, just like in javascript. Functions can be stored as a variable and then invoked using `(<parameter>)`. The code below shows how an anonymous function that gets stored as a variable and then invoked.

```Go
a := func(name string) {
		fmt.Printf("hello there, %v how are you?\n", name)
	}

a("Juan")
```

# Learning Testing

Useful packages

- testing
- iotest/testing: when working with io readers and writers.
- testing/quick: black box testing.
- net/http/httptest: contains valuable apis, response recorders, and can have test servers with end-to-end test

Community Projects

- testify: an assertion framework
- ginko: behavior driven
- GoConvey: changes the results in a browser based format
- httpexpect: re-envisioned for end-to-end tests
- gomock: simplifies the creation of mock
- go-sqlmock: from data dog, works as an in-memory mockable database.

# Go Naming Conventions

- test files are sufixed by `_test`.
- functions are prefixed by `Test`.
- you will also need a pointer to the testing.T object in your test.
- you include packages that you are testing in the package declaration. This is called whitebox testing, where the test has access to the thing. If you change it to the name of the file then it wouldnt have access to the variables, this is called blackbox testing. Blackbox is better and more representative.

# Reporting Failures

- immediate failures: for catastrophic failures.
  - t.FailNow()
  - t.Fatal()
  - t.Fatalf()
- non-immediate failures: where things are ok to fail, when testing a lot of things.
  - t.Fail()
  - t.Error()
  - t.Errorf()

# Running Tests

- you go to the directory that you are trying to test
- you can run specific test using, `go test -run` as long as you are in the directory that your test is in.

# Covers

You can get an output of the test coverage using `go test -coverprofile`

# Table Driven Tests

Create a slice of structs that you pass into the test using a for-loop

# Useful Functions

- Log and Logf
- Helper
- Skip, Skipf and SkipNow: thing for skipping
- Run: good for creating sub-tests, top level runs the test suite.
- Parallel: if you want to run tests, to run them in parallel.

# Benchmarking and Profiling

- testing.B:
  - b.N: the number of times our benchmarking to run.
  - Timer methods
    - b.StartTimer: start the timer for the benchmark
    - b.StopTimer:
    - b.ResetTimer:
  - b.RunParallel: runs routines in parallel
- Benchmarks are always prefixed with Benchmark
- run the test using the `go test -bench`
- provide a duration using the command `go test -bench -bencthtime 10s`, which will run the benchmark for 10s. Default is 1s. You could use this on end-to-end tests.
- `go test -benchmem` will show you the memory profile in go.
- `go test -trace {trace.out}` will help you trace your tests
- graphviz for go tool pprof to profile and output a file.
