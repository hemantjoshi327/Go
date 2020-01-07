package main

import(
	"testing"
)

func TestCal(t *testing.T){
	if Cal(1) != 4{
		t.Error("Expected 2 + 2 is equal 4")
	}
}