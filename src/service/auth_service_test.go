package service

import "testing"

func TestHello(t *testing.T) {
	emptyResult := hello("");

	if emptyResult == "Hello Done" {
		t.Logf("Hello with sucess. Value expected %v", emptyResult)
	} else {
		t.Errorf("Hello with fail. Value expected Hello Done passed %v", emptyResult)
	}

	emptyResult = hello("Edy")

	if emptyResult == "Hello Edy" {
		t.Logf("Hello with sucess. Value expected %v", emptyResult)
	} else {
		t.Error("Hello With Errors")
	}
}