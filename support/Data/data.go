package Data

import "Adlet/container"

func InitializeData() {
	for i := 0; i < len(container.Initial_Array); i++ {
		for j := 0; j < len(container.Initial_Array[i]); j++ {
			if container.Initial_Array[i][j] == '*' {
				container.Data_Array[i][j] = 9
				continue
			}
			if j+1 < len(container.Initial_Array[i]) && container.Initial_Array[i][j+1] == '*' {
				container.Data_Array[i][j]++
			}

			if j-1 >= 0 && container.Initial_Array[i][j-1] == '*' {
				container.Data_Array[i][j]++
			}

			if i-1 >= 0 && container.Initial_Array[i-1][j] == '*' {
				container.Data_Array[i][j]++
			}

			if i+1 < len(container.Initial_Array) && container.Initial_Array[i+1][j] == '*' {
				container.Data_Array[i][j]++
			}

			if i-1 >= 0 && j-1 >= 0 && container.Initial_Array[i-1][j-1] == '*' {
				container.Data_Array[i][j]++
			}

			if i+1 < len(container.Initial_Array) && j-1 >= 0 && container.Initial_Array[i+1][j-1] == '*' {
				container.Data_Array[i][j]++
			}

			if i+1 < len(container.Initial_Array) && j+1 < len(container.Initial_Array[i]) && container.Initial_Array[i+1][j+1] == '*' {
				container.Data_Array[i][j]++
			}

			if i-1 >= 0 && j+1 < len(container.Initial_Array[i]) && container.Initial_Array[i-1][j+1] == '*' {
				container.Data_Array[i][j]++
			}

		}
	}
}
