package set

import "fmt"

type Set struct {
	values map[interface{}]bool
}

func (s Set) Union(s2 Set) Set {
	s3 := New()
	for _, val := range s.ToArray() {
		s3.Add(val)
	}
	for _, val := range s2.ToArray() {
		s3.Add(val)
	}

	return s3

}

func (s Set) Intersection(s2 Set) Set {
	s3 := New()
	for _, val := range s.ToArray() {
		if s2.Has(val) {
			s3.Add(val)
		}
	}

	return s3
}
func (s Set) Subtraction(s2 Set) Set {
	s3 := New()
	for _, val := range s.ToArray() {
		if !s2.Has(val) {
			s3.Add(val)
		}
	}
	return s3
}
func (s Set) Has(value interface{}) bool {
	return s.values[value]
}
func (s Set) Add(value interface{}) {
	s.values[value] = true
}
func (s Set) ToArray() []interface{} {
	array := make([]interface{}, 0)
	for key, _ := range s.values {
		array = append(array, key)
	}
	return array
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
