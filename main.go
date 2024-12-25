package main

import (
	"fmt"
	rand "math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width = 680
	height = 680
	cellSize = 40
	rowsCount = height / cellSize
	colsCount = width / cellSize
)

var (
	gray = rl.NewColor(31, 35, 53, 255)
	purple = rl.NewColor(157, 124, 216, 255)
	cellsMatrix = []Cell{}
	cellSizeV2 = rl.NewVector2(cellSize, cellSize)
	frameCounter = 0
)

type Cell struct {
	X float32;
	Y float32;
	Status bool;
}


func main() {
	rl.InitWindow(width, height, "Game of Life")
	defer rl.CloseWindow()

	rl.SetTargetFPS(144)
	cellsMatrix = initializeCells(cellsMatrix)

	fmt.Printf("")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(gray)

		
		if frameCounter % 10 == 0 {
			calculateStep()
		}
		for _, cell := range cellsMatrix {
			cellV2 := rl.NewVector2(cell.X, cell.Y)
			color := purple
			if !cell.Status {
				color = gray
			}
			rl.DrawRectangleV(cellV2, cellSizeV2, color)
		}
		for i := 0; i <= rowsCount; i++ {
			rl.DrawLine(int32(cellSize*i), 0, int32(cellSize*i), height, rl.White)
		}

		for j := 0; j <= colsCount; j++ {
			rl.DrawLine(0, int32(cellSize*j), width, int32(cellSize*j), rl.White)
		}

		frameCounter++
		rl.EndDrawing()
	}
}

func initializeCells(cells []Cell) []Cell {
	for i := 0; i < rowsCount; i++ {
		for j := 0; j < colsCount; j++ {
			fmt.Println("Status: ", rand.Int() % 2 == 1)
			cell := Cell { X: float32(i*cellSize), Y: float32(j*cellSize), Status: rand.Int() % 2 == 1}	
			cells = append(cells, cell)
		}
	}

	return cells
}

func calculateStep() {
	//Como que ele esquece do j + I * width
	//linhas + colunas * tamanho da minha linha (quantidade de colunas)
	//pode ser que seja colunas+linhas também, I e J são nomes idiotas.
	for i := 0; i < rowsCount; i++ {
		for j := 0; j < colsCount; j++ {
			counter := calculateNeighbors(cellsMatrix, j, i, colsCount, rowsCount)	
			cell := &cellsMatrix[j + i * colsCount]

			fmt.Printf("X: %f, Y: %f, Neighbor: %d\n", cell.X, cell.Y, counter)

			if counter < 2 && cell.Status {
				cell.Status = false
			} else if cell.Status && (counter == 2 || counter == 3) {
				cell.Status = true
			} else if cell.Status && counter > 3 {
				cell.Status = false
			} else if !cell.Status && counter == 3 {
				cell.Status = true
			}
		}
	}
}

func calculateNeighbors(cellsMatrix []Cell, column int, line int, colsCount int32, rowsCount int32) int {
	counter := 0
	
	// full left neighbor
	if 	column > 0 {
		if cellsMatrix[column-1 + line * int(colsCount)].Status {
			counter+=1
		}
	}

	// full right neighbor
	if column < int(colsCount-1)  {
		if cellsMatrix[(column+1) + line * int(colsCount)].Status {
			counter+=1
		}
	}

	// top one
	if line > 0 {
		if cellsMatrix[column + (line-1) * int(colsCount)].Status {
			counter+=1
		}
	}

	// bottom one
	if line < int(rowsCount-1) {
		if cellsMatrix[column + (line+1) * int(colsCount)].Status {
			counter+=1
		}
	}
	
	// top left
	if line > 0 && column > 0 {
		if cellsMatrix[(column-1) + (line-1) * int(colsCount)].Status {
			counter+=1
		}
	}

	// top right
	if line > 0 && column < int(colsCount-1) {
		if cellsMatrix[(column+1) + (line-1) * int(colsCount)].Status {
			counter+=1
		}
	}
	
	// bottom right
	if line < int(rowsCount-1) && column < int(colsCount-1) {
		if cellsMatrix[(column+1) + (line+1) * int(colsCount)].Status {
			counter+=1
		}
	}

	// bottom left
	if line < int(rowsCount-1) && column > 0 {
		if cellsMatrix[(column-1) + (line+1) * int(colsCount)].Status {
			counter+=1
		}
	}

	return counter
}
