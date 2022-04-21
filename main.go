package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

//Example 1
//func or method that will act on Student
//var name that reps Student=type then name of method, return int
func (s Student) getAge() int {
	return s.age
}

// s is equal to what student is declared
func main() {
	s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
	s1.getAge()
}

//Example 2
//method that will change age
//takes age as arg
//pointer to Student in order to change or set new value on respective Student
func (s *Student) setAge(age int) {
	s.age = age
}

//will not work when wanting to change or set unless pointer
func main() {
	s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
	fmt.Println(s1)
	s1.setAge(7)
	fmt.Println(s1)
}

//Example 3
//var name s of Student, method name returning float
//no pointer because returning a float and not changing or setting something new on Student
//useful for when multiple student or structs because can declare new student in main and use method on each
func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func main() {
	s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
	average := s1.getAverageGrade()
	fmt.Println(average)

}
