package main

import "fmt"

type Learner interface {
	Learn(subject string) string
}

type Teacher interface {
	Teach()
}

type Student struct {
	Name string
}

func (s Student) Learn(subject string) string {
	fmt.Printf("%s is learning %s\n", s.Name, subject)
	return s.Name
}

type Professor struct {
	Name string
}

func (t Professor) Teach() {
	fmt.Printf("%s is teaching\n", t.Name)
}

func LearnFrom(t Teacher) {
	t.Teach()
}

func TeachTo(s Learner) {
	s.Learn("Computer science")
}

func main() {
	var p Professor
	p.Name = "Andrew"
	LearnFrom(p)

	var s Student
	s.Name = "Steve"
	TeachTo(s)
}
