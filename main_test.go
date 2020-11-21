package main

import "testing"

func TestGreeter(t *testing.T) {
	if greeter("John") != "Hello John!" {
		t.Fatalf("greeter(\"John\") != \"Hello John!\"")
	}
}
