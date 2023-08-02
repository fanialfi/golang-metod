package main

import (
	"fmt"
	"strings"
)

type Student struct {
	name  string
	grade int
}

func (s Student) sayHello() {
	fmt.Println("hallo", s.name)
}
func (s Student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}
func (s Student) changeName1(name string) {
	fmt.Println("--> di changeName1, nama diubah ke", name)
	s.name = name
}
func (s *Student) changeName2(name string) {
	fmt.Println("--> di changeName2, nama diubah ke", name)
	s.name = name
}

func main() {
	// var s1 = Student{"fani alfi", 17}
	// s1.sayHello()

	// var name = s1.getNameAt(2)
	// fmt.Println("nama panggilan :", name)

	var s2 = Student{"fani alfi", 17}
	fmt.Println("s1 sebelum namanya diubah", s2.name)

	s2.changeName1("Jason Burneo")
	fmt.Println("s1 setelah diubah menggunakan changeName1", s2.name)
	// fani alfi

	s2.changeName2("fani alfirdaus")
	fmt.Println("s1 setelah diubah menggunakan changeName2", s2.name)
	// fani alfirdaus

	// var s2 = Student{"fani alfi", 17}
	// fmt.Println("s1 before", s2.name)
	// // // pengaksesan method dari variabel struct biasa
	// // s2.sayHello()

	// s2.changeName1("Jason Burneo")
	// fmt.Println("after changeName1", s2.name)

	// s2.changeName2("fani alfirdaus")
	// fmt.Println("after changeName2", s2.name)

	// fmt.Println()
	// // pengaksesan method dari variabel struct pointer
	// var s3 *Student = &Student{"Ethan Hunt", 22}
	// s3.sayHello()
	// fmt.Println(*s3)
}
