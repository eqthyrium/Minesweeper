package main

import (
	"Adlet/function"
	"Adlet/support"
	"Adlet/support/Data"
	"Adlet/support/Output"
	"fmt"
)

func main() {
	if err := support.Initialize(); err {
		fmt.Println("Error: invalid input.")
		return
	}

	Data.InitializeData()
	Output.Tablize()
	function.Solve()
}
