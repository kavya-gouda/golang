package main

import "fmt"
func main() {
	var name string = "golang var"
    // golang will infer the type
	var name2 = "golang var2"
	fmt.Println(name, name2)

	// shorthand syntax
	name3 := "hello"
	fmt.Println(name3)
	// there are situation where variables are defined first and then value will be assigned in that case we need to use var name string 

	var name4 string
	name4 = "hello world"
	fmt.Println(name4)
}