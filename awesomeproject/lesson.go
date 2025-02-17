package main

import (
	"fmt"
	"os/user"
	"time"
)

func init(){
	fmt.Println("initialized!")
}

func bazz(){
	fmt.Println("Bazz")
}

func main(){
	bazz()
	fmt.Println("HELLO WORLD")
	fmt.Println(time.Now())
	fmt.Println(user.Current())

	var i int = 1
	fmt.Println(i)

	var t,f bool = true, false
	fmt.Println(t,f)


	var (
		j int = 3
		isActive bool = false
	)
	fmt.Println(j, isActive)
}