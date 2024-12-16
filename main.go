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

		for _, cell := range cellsMatrix {
			cellV2 := rl.NewVector2(cell.X, cell.Y)
			color := purple
			if !cell.Status {
				color = gray
			}
			rl.DrawRectangleV(cellV2, cellSizeV2, color)
		}
		for i := 0; i <= rowsCount; i++ {
			rl.DrawLine(int32(cellSize*i), 0, int32(cellSize*i), height, rl.RayWhite)
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
