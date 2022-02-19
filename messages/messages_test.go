package messages

import (
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("gopher")
	expect := "hello, gopher\n"

	if got != expect {
		t.Errorf("Did not get expected. Wanted %q, got %q\n", got, expect)
	}
}

func TestDepart(t *testing.T) {
	got := depart("gopher")
	expect := "bye now, gopher"

	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got %q\n", got, expect)
	}
}

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "hello, Gopher\n"},
		{input: "", expect: "hello, \n"},
	}

	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not get expected. Wanted %q, got %q\n", got, s.expect)
		}
	}
}

func TestFailures(t *testing.T) {
	t.Error("This should run")
	t.Fatal("This should run an stop execution")
	t.Error("This wont be reported...")
}
