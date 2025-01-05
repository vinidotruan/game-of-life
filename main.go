package main

import (
	// rand "math/rand"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width = 800
	height = 800
	cellSize = 40
	rowsCount = height / cellSize
	colsCount = width / cellSize
)

var (
	gray = rl.NewColor(31, 35, 53, 255)
	purple = rl.NewColor(157, 124, 216, 255)
	white = rl.NewColor(86, 95, 137, 255)
	cellsMatrix = [rowsCount][colsCount]bool{}
	cellSizeV2 = rl.NewVector2(cellSize, cellSize)
	frameCounter = 0
)

type Cell struct {
	X float32;
	Y float32;
 bool;
}

func main() {
	rl.InitWindow(width, height, "Game of Life")
	defer rl.CloseWindow()

	rl.SetTargetFPS(144)
	initializeCells(&cellsMatrix)
	
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(gray)
		
		for rowIndex, rowValue := range cellsMatrix {
			for colIndex, cell := range rowValue {
				cellV2 := rl.NewVector2(float32(colIndex)*cellSize, float32(rowIndex)*cellSize)
				color := purple
				if !cell {
					color = gray
				}
				rl.DrawRectangleV(cellV2, cellSizeV2, color)
			}
		}

		for i := 0; i <= rowsCount; i++ {
			start := rl.NewVector2(float32(cellSize*i), 0)
			end := rl.NewVector2(float32(cellSize*i), float32(height))
			rl.DrawLineEx(start, end, float32(3), white)
		}

		for j := 0; j <= colsCount; j++ {
			start := rl.NewVector2(0, float32(cellSize*j))
			end := rl.NewVector2(width, float32(cellSize*j))
			rl.DrawLineEx(start, end, float32(3), white)
		}

		if frameCounter % 10 == 0 {
			 calculateStep()
		}
		frameCounter++
		rl.EndDrawing()
	}
}

func initializeCells(cells *[rowsCount][colsCount]bool) *[rowsCount][colsCount]bool {
	for i := 0; i < rowsCount; i++ {
		for j := 0; j < colsCount; j++ {
			cells[i][j] = false
		}
	}

	// cells[0][1] = true
	// cells[1][2] = true
	// cells[2][0] = true
	// cells[2][1] = true
	// cells[2][2] = true

	cells[rowsCount/2][colsCount/2 + 2] = true
	cells[rowsCount/2 + 1][colsCount/2 + 1] = true
	cells[rowsCount/2 + 1][colsCount/2 + 2] = true
	cells[rowsCount/2 + 2][colsCount/2 + 2] = true
	cells[rowsCount/2 + 2][colsCount/2 + 3] = true
		
	return cells
}

func calculateStep() {
	newMatrix := cellsMatrix

	for rowIndex, rowValue := range cellsMatrix {
		for colIndex, cellValue := range rowValue {
			counter := calculateNeighbors(cellsMatrix, colIndex, rowIndex)	
			// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
			if counter < 2 &&  cellValue {
				newMatrix[rowIndex][colIndex] = false
				continue
			}

			// Any live cell with two or three live neighbours lives on to the next generation
			if cellValue && (counter == 2 || counter == 3) {
				newMatrix[rowIndex][colIndex] = true
				continue
			}

			// Any live cell with more than three live neighbours dies, as if by overpopulation
			if cellValue && counter > 3 {
				newMatrix[rowIndex][colIndex] = false
				continue
			}

			// Any dead cell with exactly three live neighbours becomes a live cell as if by reproduction
			if !cellValue && counter == 3 {
				newMatrix[rowIndex][colIndex] = true
				continue
			}
		}
	}

	cellsMatrix = newMatrix
}

func calculateNeighbors(cellsMatrix [rowsCount][colsCount]bool, column int, line int) int {
	counter := 0

	// full left neighbor
	if 	column > 0 {
		if cellsMatrix[line][column-1] {
			counter+=1
		}
	}
	// full right neighbor
	if column < int(colsCount-1) {
		if cellsMatrix[line][column+1] {
			counter+=1
		}
	}
	// top left
	if line > 0 && column > 0 {
		if cellsMatrix[line-1][column-1] {
			counter+=1
		}
	}
	// top one
	if line > 0 {
		if cellsMatrix[line-1][column] {
			counter+=1
		}
	}
	// top right
	if line > 0 && column < int(colsCount-1) {
		if cellsMatrix[line-1][column+1] {
			counter+=1
		}
	}
	// bottom right
	if line < int(rowsCount-1) && column < int(colsCount-1) {
		if cellsMatrix[line+1][column+1] {
			counter+=1
		}
	}
	// bottom one
	if line < int(rowsCount-1) {
		if cellsMatrix[line+1][column] {
			counter+=1
		}
	}
	// bottom left
	if line < int(rowsCount-1) && column > 0 {
		if cellsMatrix[line+1][column-1] {
			counter+=1
		}
	}

	return counter
}
