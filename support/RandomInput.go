package support

import (
	"Adlet/container"
	"fmt"
	"math/rand"
	"time"
)

func RandomInitialize() {
	fmt.Println("Enter the height and width of the map: ")
	fmt.Scanf("%v %v", &container.Height, &container.Width)

	fmt.Println("Could you give me the amount of the bomb which will be present in the game ?")
	fmt.Println("Enter the number:")
	fmt.Scanf("%v", &container.Bomb_amount)
	container.Initial_Array = make([][]rune, container.Height)
	container.Data_Array = make([][]int, container.Height)
	container.State_Array = make([][]bool, container.Height)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < container.Height; i++ {
		container.Initial_Array[i] = make([]rune, container.Width)
		container.Data_Array[i] = make([]int, container.Width)
		container.State_Array[i] = make([]bool, container.Width)

		for j := 0; j < container.Width; j++ {
			randomNumber := rand.Intn(container.Height * container.Width)
			if randomNumber >= (container.Height*container.Width*3)/10 {
				container.Initial_Array[i][j] = '.'
			} else {
				container.Initial_Array[i][j] = '*'
			}

		}
	}
}