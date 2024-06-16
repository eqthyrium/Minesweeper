package support

import (
	"Adlet/container"
	"fmt"
)

func Initialize() bool {
	var height, width int = 0, 0
	var promt_error bool = false

	fmt.Println("Enter your coordiantes: ")
	fmt.Scanf("%v %v", &height, &width)
	if (height >= 1 && height < 3) || (width >= 1 && width < 3) {
		promt_error = true
	}

	if height <= 0 || width <= 0 {
		return true
	}

	container.Height = height
	container.Width = width

	container.Initial_Array = make([][]rune, height)
	container.Data_Array = make([][]int, height)
	container.State_Array = make([][]bool, height)
	var line string
	var counter_bomb int = 0
	for i := 0; i < height; i++ { // Saving into array the characters
		fmt.Scanf("%v", &line)
		if len(line) != width {
			promt_error = true
			break
		}
		container.Initial_Array[i] = make([]rune, width)
		container.Data_Array[i] = make([]int, width)
		container.State_Array[i] = make([]bool, width)

		for j := 0; j < len(line); j++ {
			if !(line[j] == '.' || line[j] == '*') {
				promt_error = true
				break
			}

			if line[j] == '*' {
				counter_bomb++
			}
			container.Initial_Array[i][j] = rune(line[j])
		}

		if promt_error {
			break
		}

	}

	container.Bomb_amount = counter_bomb

	if promt_error || counter_bomb < 2 {
		return true
	}

	return false
}
