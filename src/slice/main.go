package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func main() {

	var d *Student
	d = &Student{"ddd", 10, 100}
	a := Student{"aaa", 20, 10}
	b := a
	b.age = 10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a)

	var c *Student
	c = &a
	c.age = 10
	fmt.Println(*c)

	a.SetName("bbb")
	fmt.Println(a)

	a.SetAge(41)
	fmt.Println(a)

	d.SetAge(14)
	fmt.Println(d)

	d.SetAge(18)
	fmt.Println(d)

	PrintStudent(&a)
}

func (t *Student) SetName(newName string) {
	t.name = newName
}

func (t *Student) SetAge(newAge int) {
	t.age = newAge
}

func (t *Student) SetGrade(grade int) {
	t.grade = grade
}

func PrintStudent(u *Student) {
	fmt.Println(u)
}
