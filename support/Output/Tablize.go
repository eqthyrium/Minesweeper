package Output

import (
	"Adlet/container"
	"fmt"
)

func full() {
}

func element() {
}

func Tablize() {
	var ll string = ""
	fmt.Printf("%5s ", ll)
	for i := 0; i < container.Width; i++ {
		fmt.Print("    ")
		fmt.Print(i + 1)
		fmt.Print("   ")
	}
	fmt.Println()
	fmt.Printf("%5s  ", ll)
	for i := 0; i < container.Width; i++ {
		for j := 0; j < 8; j++ {
			if i+1 == container.Width {
				if j != 7 {
					fmt.Print("_")
				}
			} else {
				fmt.Print("_")
			}
		}
	}
	fmt.Println()
	for i := 0; i < len(container.State_Array); i++ {
		for z := 0; z < 3; z++ {
			if z == 1 {
				fmt.Printf("%5v |", i+1)
			} else {
				fmt.Printf("%5s |", ll)
			}

			for w := 0; w < len(container.State_Array[i]); w++ {
				for q := 0; q < 7; q++ {
					if !container.State_Array[i][w] {
						fmt.Print("X")
					} else {
						if z == 1 && q == 3 {
							if container.Data_Array[i][w] != 9 {
								if container.Data_Array[i][w] == 0 {
									fmt.Print(" ")
								} else {
									fmt.Print(container.Data_Array[i][w])
								}
							} else {
								fmt.Print("*")
							}
						} else if z == 2 {
							fmt.Print("_")
						} else {
							fmt.Print(" ")
						}
					}
				}
				fmt.Print("|")
			}
			fmt.Println()

		}
	}
}
