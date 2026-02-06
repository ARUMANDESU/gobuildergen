package main

import (
	"time"
)

type ExampleInterface interface {
	Example()
}

//go:generate go run ../main.go --type ExampleStruct
type ExampleStruct struct {
	Field1 string
	Field2 []string
	field3 *string
	field4 time.Time
	Field5 *time.Time
	Field6 map[string]string
	Field7 chan int
}
