package set

import "fmt"
import "reflect"

type Set struct {
	values  map[interface{}]bool
	sizePtr *int
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
func (s Set) Remove(value interface{}) {
	if !s.Has(value) {
		return
	}
	delete(s.values, value)
	(*s.sizePtr)--

}
func (s Set) Add(value interface{}) {
	if s.Has(value) {
		return
	}
	s.values[value] = true
	(*s.sizePtr)++

}
func (s Set) ToArray() []interface{} {
	array := make([]interface{}, 0)
	for key, _ := range s.values {
		array = append(array, key)
	}
	return array
}

func (s Set) Size() int {
	return *s.sizePtr
}
func (s Set) Show() {
	for key, val := range s.values {
		fmt.Printf("%v %t\n", key, val)
	}
}
func FromSlice(in interface{}) Set {
	s := Set{make(map[interface{}]bool), new(int)}
	if reflect.TypeOf(in).Kind() != reflect.Slice {
		fmt.Printf("Warning! %v is not slice type!\n", in)
		return New(in)
	}
	ins := reflect.ValueOf(in)
	for i := 0; i < ins.Len(); i++ {
		s.Add(ins.Index(i).Interface())
	}
	return s
}
func New(ins ...interface{}) Set {
	//fmt.Println("You have create a new Set!")
	//fmt.Println("It just is a demo and updating...")
	//fmt.Println("I'm in set.Set()")
	s := Set{make(map[interface{}]bool), new(int)}
	for _, val := range ins {
		s.Add(val)
	}
	return s
}
