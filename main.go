package main

import (
	"math"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WinWidth = 900
const WinHeight = 600

// Basic struct for all astronomical objects

type Planet struct {
	posX   int32
	posY   int32
	radius int32
	color  rl.Color
}

func main() {
	rl.InitWindow(WinWidth, WinHeight, "Solar System")
	rl.SetTargetFPS(60)

	// Initialise the sun and one planet

	var sun Planet = Planet{WinWidth / 2, WinHeight / 2, 50, rl.Yellow}
	var mercury Planet = Planet{sun.posX / 2, WinHeight / 2, 20, rl.DarkGray}
	var angle float64 = math.Pi
	var orbitSize float64 = 100

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Draw the Sun and the one planet

		rl.DrawCircle(sun.posX, sun.posY, float32(sun.radius), rl.Yellow)
		rl.DrawCircle(mercury.posX, mercury.posY, float32(mercury.radius), mercury.color)

		//rl.DrawCircle(mercury.posX+int32(150*math.Sin(speed)), mercury.posY, float32(mercury.radius), mercury.color)

		// Update position on orbit
		// X := originX + cos(angle)*radius;
		// Y := originY + sin(angle)*radius;

		mercury.posX = sun.posX + int32(math.Cos(angle)*orbitSize)
		mercury.posY = sun.posY + int32(math.Sin(angle)*orbitSize)
		angle += math.Atan(float64(mercury.posX / mercury.posY))

		rl.DrawText(strconv.Itoa(int(mercury.posX)), 0, 0, 50, rl.Brown)

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
