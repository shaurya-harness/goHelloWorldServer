package main

import "testing"

func TestGreetingSpecificJohn(t *testing.T) {
	t.Errorf("Greeting was incorrect")
}

func TestGreetingSpecificDemo(t *testing.T) {
	t.Errorf("Greeting was incorrect")
}

func TestShowFailure(t *testing.T) {
	t.Errorf("Intentional failure")
}



func TestGreetingDefault(t *testing.T) {
	t.Errorf("Greeting was incorrect")
}
 


