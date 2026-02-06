package example

import (
	"time"

	examplepackage "githab.com/ARUMANDESU/gobuildergen/example/example-package"
)

type ExampleInterface interface {
	Example()
}

//go:generate go run ../main.go --type ExampleStruct
type ExampleStruct struct {
	Field1  string
	Field2  []string
	field3  *string
	field4  time.Time
	Field5  *time.Time
	Field6  map[string]string
	Field7  chan int
	Field8  ExampleFieldStruct
	Field9  *ExampleFieldStruct
	Field10 examplepackage.ExampleTest
	Field11 *examplepackage.ExampleTest
	Field12 []examplepackage.ExampleTest
	Field13 []*examplepackage.ExampleTest
	Field14 map[ExampleFieldStruct]string
	Field15 map[examplepackage.ExampleTest]examplepackage.ExampleTest
	Field16 map[examplepackage.ExampleTest]string
	Field17 [8]string
	Field18 func(string) (string, error)
}

type ExampleFieldStruct struct {
	Test string
}
