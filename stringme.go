package main

import "fmt"

type student struct {
	Name string
}

func (s student) String() string {
	return fmt.Sprintf("%s (who's a Student)", s.Name)
}

func main() {
	var s student
	s.Name = "Andrew"
	fmt.Printf("%s\n", s)
}
