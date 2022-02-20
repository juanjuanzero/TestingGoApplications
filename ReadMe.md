# Learning Interfaces and Structs

Working with interface and struct

- in go, an interface is just something that holds a set of functions, structs hold private and public fields
- when you have Pointer based receivers, the method that you work in, works on the object stored in memory (dereference) a pointer to that actuall object as opposed to value based functions.

You see that the functions in the Person struct type have an \* next to the struct name Person. This is signifies that this function is pointing to the object and it works on that object instance.

```Go
//interfaces only contain functions
type Indetifiable interface {
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

Conflicts, you can still have the same field name types, the compiler will stop you from doing that. If you have a conflict.

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
