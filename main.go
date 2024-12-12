package main

import (
	"fmt"
	rand "math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width = 620
	height = 620
	cellSize = 40
	rowsCount = height / cellSize
	colsCount = width / cellSize
)

var (
	gray = rl.NewColor(31, 35, 53, 255)
	purple = rl.NewColor(157, 124, 216, 255)
	cellsMatrix = []rl.Vector2{
		getRandomV2(),
		getRandomV2(),
		getRandomV2(),
	}
	cellSizeV2 = rl.NewVector2(cellSize, cellSize)
	frameCounter = 0
)

func main() {
	rl.InitWindow(width, height, "GoL")
	defer rl.CloseWindow()

	rl.SetTargetFPS(144)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(gray)

		fmt.Println("Resultado: ", (width-cellSize)/cellSize)
		fmt.Println("Numero random: ", rand.Intn((width-cellSize)/cellSize))

		for i := 0; i <= rowsCount; i++ {
			rl.DrawLine(int32(cellSize*i), 0, int32(cellSize*i), height, rl.RayWhite)
		}

		for j := 0; j <= colsCount; j++ {
			rl.DrawLine(0, int32(cellSize*j), width, int32(cellSize*j), rl.White)
		}

		for index, k := range cellsMatrix {
			rl.DrawRectangleV(k, cellSizeV2, purple)
			if  frameCounter % 14 == 0 {
				if k.X+cellSize <= width {
					cellsMatrix[index].X += cellSize
				} else {
					cellsMatrix[index].X = 0 
					cellsMatrix[index].Y += cellSize
				}

				if k.Y+cellSize > height {
					cellsMatrix[index].Y = 0 
				}


				fmt.Println("X: ", k.X, "Y: ", k.Y)
			}

		}

		frameCounter++

		rl.EndDrawing()
	}
}

func getRandomV2() rl.Vector2 {
	return rl.NewVector2(
		float32(rand.Intn((width-cellSize)/cellSize)*cellSize),
		float32(rand.Intn((width-cellSize)/cellSize)*cellSize),
	)	

}
