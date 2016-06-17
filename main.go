package main

import (
	"fmt"
	"time"
)

// DbRecord is a db record wrapper
type DbRecord struct {
	name      string
	createdAt time.Time
}

func main() {
	fmt.Println(IAmACoolFunction(RealClock{}).createdAt)
}

// RealClock actually is a real clock
type RealClock struct{}

// Now really is now!
func (r RealClock) Now() time.Time {
	return time.Now()
}

// Nowable responds to now
type Nowable interface {
	Now() time.Time
}

// IAmACoolFunction does something
func IAmACoolFunction(clock Nowable) DbRecord {
	now := clock.Now()
	after := now.AddDate(0, 6, 0)

	return DbRecord{"HelloWorld", after}
}
