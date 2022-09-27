package main

import "testing"

func TestGreetingSpecificJohn(t *testing.T) {
	t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
}

func TestGreetingSpecificDemo(t *testing.T) {
	t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
}

func TestShowFailure(t *testing.T) {
	t.Errorf("Intentional failure. got: %s, want: %s.", greeting, "Hello, Demo\n")
}



func TestGreetingDefault(t *testing.T) {
	t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Guest\n")
}
 


