package main

import "fmt"

type User struct {
	name string
	age  int
}

func IncrementAge(user *User) {
	user.age++
}

func main() {
	a := User{"Bob", 20}
	IncrementAge(&a)
	fmt.Println(a)
}
