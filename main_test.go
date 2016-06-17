package main

import (
	"testing"
	"time"
)

type TestClock struct{}

func (t TestClock) Now() time.Time {
	testTime, _ := time.Parse(time.RFC822, "02 Jan 06 15:04 UTC")
	return testTime
}

func TestIAmACoolFunction(t *testing.T) {
	res := IAmACoolFunction(TestClock{})

	extime, _ := time.Parse(time.RFC822, "02 Jul 06 15:04 UTC")

	expected := DbRecord{"HelloWorld", extime}

	if res != expected {
		t.Errorf("The db record does not match %#v expected but got %#v", expected, res)
	}
}
