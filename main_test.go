package main

import "testing"
 
func TestSayHello(t *testing.T) {
	msg := sayHello("Go")

	want := "Hello Go Gopher"

	if want != msg {
		t.Error()
	}
}
