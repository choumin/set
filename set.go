package set

import "fmt"

type Set struct {
	values map[interface{}]bool
}

func (s Set) Add(value interface{}) {
	s.values[value] = true
}

func (s Set) Length() int {
	return len(s.values)
}
func (s Set) Show() {
	for key, val := range s.values {
		fmt.Printf("%v %t\n", key, val)
	}
}
func New() Set {
	//fmt.Println("You have create a new Set!")
	//fmt.Println("It just is a demo and updating...")
	//fmt.Println("I'm in set.Set()")
	return Set{make(map[interface{}]bool)}
}
