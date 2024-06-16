package main

import (
	"Adlet/function"
	"Adlet/support"
	"Adlet/support/Data"
	"Adlet/support/Output"
	"fmt"
)

func main() {
	var option int
	fmt.Println("Creating table by manual - 1")
	fmt.Println("Creating table by randomizer - 2")
	fmt.Println("Enter the option:")
	fmt.Scanf("%v", &option)

	if option == 1 {
		if err := support.Initialize(); err {
			fmt.Println("Error: invalid input.")
			return
		}
	} else {
		support.RandomInitialize()
	}

	Data.InitializeData()
	Output.Tablize()
	function.Solve()
}
