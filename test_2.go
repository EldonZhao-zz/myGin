package main

import (
	"fmt"
)

// import "strings"

type Man interface {
	study() string
	eat() string
}

type Student struct {
	name string
	food string
	book string
}

type Worker struct {
	name string
	food string
	book string
}

func (s Student) study() string {
	return "Student " + s.name + " study " + s.book
}

func (s Student) eat() string {
	return "Student " + s.name + " study " + s.food
}

func (w Worker) study() string {
	return "Worker " + w.name + " study " + w.book
}

func (w Worker) eat() string {
	return "Worker " + w.name + " study " + w.food
}

func testItf() {
	s := Student{name: "123", book: "xasx", food: "xasax"}
	fmt.Print(s.study())
}
