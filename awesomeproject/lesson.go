package main

import (
	"fmt"
	// "os/user"
	// "time"
)

func init() {
	fmt.Println("initialized!")
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	// bazz()
	// fmt.Println("HELLO WORLD")
	// fmt.Println(time.Now())
	// fmt.Println(user.Current())

	// var i int = 1
	// fmt.Println(i)

	// var t,f bool = true, false
	// fmt.Println(t,f)

	// var (
	// 	j int = 3
	// 	isActive bool = false
	// )
	// fmt.Println(j, isActive)

	var a [2]int
	fmt.Println(a)

	var s []int
	s = append(s, 1)
	fmt.Println(s)

	ss := []int{1, 2, 3}
	fmt.Println(ss)

	m := map[string]int{"apple": 100}
	fmt.Println((m))
}
