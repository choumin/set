package set

import "fmt"

type Set struct {
	values map[string]int
}

func (s *Set) Add(val string) {
	s.values[val] = 1
}

func New() *Set {
	//fmt.Println("You have create a new Set!")
	//fmt.Println("It just is a demo and updating...")
	fmt.Println("I'm in set.Set()")
	return &Set{make(map[string]int)}

}
