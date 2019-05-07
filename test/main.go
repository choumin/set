package main

import "fmt"
import "../../set"

func test_set_basic() {
	s := set.New()
	s.Add("hello")
	s.Add(123)
	s.Add('a')
	s.Add(true)
	fmt.Println(s.Length())
	s.Show()
}

func main() {
	test_set_basic()
}
