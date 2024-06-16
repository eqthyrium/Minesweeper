package function

import "Adlet/container"

func solve(i, j int) {
	if container.Data_Array[i][j] != 0 {
		container.State_Array[i][j] = true
		OpenedCell++
		return
	}

	if container.State_Array[i][j] == false {
		container.State_Array[i][j] = true
		OpenedCell++

		if i+1 < len(container.State_Array) && container.State_Array[i+1][j] == false {
			solve(i+1, j)
		}

		if i-1 >= 0 && container.State_Array[i-1][j] == false {
			solve(i-1, j)
		}

		if j+1 < len(container.State_Array[i]) && container.State_Array[i][j+1] == false {
			solve(i, j+1)
		}

		if j-1 >= 0 && container.State_Array[i][j-1] == false {
			solve(i, j-1)
		}

		if i+1 < len(container.State_Array) && j-1 >= 0 && container.State_Array[i+1][j-1] == false {
			solve(i+1, j-1)
		}

		if i+1 < len(container.State_Array) && j+1 < len(container.State_Array[i]) && container.State_Array[i+1][j+1] == false {
			solve(i+1, j+1)
		}

		if i-1 >= 0 && j+1 < len(container.State_Array[i]) && container.State_Array[i-1][j+1] == false {
			solve(i-1, j+1)
		}

		if i-1 >= 0 && j-1 >= 0 && container.State_Array[i-1][j-1] == false {
			solve(i-1, j-1)
		}
	}
}

func OpenThePath(y, x int) {
	solve(y, x)
}
