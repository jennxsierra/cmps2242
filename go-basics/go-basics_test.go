// Filename: hello_test.go
// go test -v

package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, world!"
	result := Hello()

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}