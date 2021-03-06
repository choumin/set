package main

import "fmt"
import "../../set"

func test_set_basic() {
	s := set.New()
	s.Add("hello")
	s.Add(123)
	s.Add('a')
	s.Add(true)
	fmt.Println(s.Size())
	s.Show()
	fmt.Printf("%s is in s? %t\n", "hello", s.Has("hello"))
	fmt.Printf("%d is in s? %t\n", 123, s.Has(123))

	s2 := set.New()
	s2.Add("hello")
	fmt.Println(s.Union(s2))
	fmt.Println(s.Intersection(s2))
	fmt.Println(s.Subtraction(s2))
	fmt.Println(s.ToArray())

	s3 := set.New(1, "2", "hello", false)
	s3.Show()
	s3.Remove("2")
	s3.Show()
	fmt.Println(s3.Size())
	//fmt.Println(s3.values)

	s4 := set.FromSlice([]int{1, 2, 3, 4})
	s4.Show()

	s5 := set.FromSlice([]bool{false, true})
	s5.Show()

	s6 := set.FromSlice([]string{"hello", "world"})
	s6.Show()

	s7 := set.FromSlice(1)
	s7.Show()

}

func main() {
	test_set_basic()
}
