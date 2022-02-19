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

-
