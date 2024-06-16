package function

import (
	"Adlet/container"
	"Adlet/support/Data"
	"Adlet/support/Output"
	"fmt"
	"math/rand"
	"time"
)

var OpenedCell int = 0

func settingbomb() {
	for i := 0; i < len(container.State_Array); i++ {
		for j := 0; j < len(container.State_Array); j++ {
			if container.Data_Array[i][j] == 9 {
				container.State_Array[i][j] = true
			}
		}
	}
}

func randomizer(y, x int) {
	container.Initial_Array[y][x] = '.'

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(container.Height * container.Width)
	xrand := randomNumber % container.Height
	yrand := (randomNumber) % container.Width

	for container.Initial_Array[yrand][xrand] != '.' {
		randomNumber := rand.Intn(container.Height * container.Width)
		xrand = randomNumber % container.Height
		yrand = (randomNumber) % container.Width

		if container.Initial_Array[yrand][xrand] == '.' {
			container.Initial_Array[yrand][xrand] = '*'
			break
		}
	}
	container.Data_Array = make([][]int, container.Height)
	for i := 0; i < len(container.Data_Array); i++ {
		container.Data_Array[i] = make([]int, container.Width)
	}

	Data.InitializeData()
}

func Solve() {
	var x, y int
	var moves int = 0
	for (container.Height*container.Width)-container.Bomb_amount != OpenedCell {
		fmt.Println("Enter your coordinates:")
		fmt.Scanf("%v %v", &y, &x)
		moves++
		y--
		x--
		if !(x >= 0 && y >= 0) {
			moves--
			fmt.Println("Invalid input.")
			continue
		}

		if container.Data_Array[y][x] == 9 {
			if moves == 1 {
				randomizer(y, x)
				moves--
				continue
			}
			settingbomb()
			Output.Tablize()
			fmt.Println("Game Over!")
			fmt.Println("Your statistics:")
			fmt.Printf("- Field size: %vx%v\n", container.Height, container.Width)
			fmt.Println("- Number of bombs:", container.Bomb_amount)
			fmt.Println("- Number of moves:", moves)
			return
		} else if container.Data_Array[y][x] == 0 {
			OpenThePath(y, x)
		} else {
			container.State_Array[y][x] = true
			OpenedCell++
		}

		Output.Tablize()

	}

	fmt.Println("You Win!")
	fmt.Printf("- Field size: %vx%v\n", container.Height, container.Width)
	fmt.Println("- Number of bombs:", container.Bomb_amount)
	fmt.Println("- Number of moves:", moves)
}
