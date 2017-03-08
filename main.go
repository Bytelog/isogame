package main

import "github.com/gen2brain/raylib-go/raylib"

func main() {
	screenWidth := int32(1300)
	screenHeight := int32(700)

	raylib.InitWindow(screenWidth, screenHeight, "ISOGAME")
	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	img := raylib.LoadImage("grass.png")
	defer raylib.UnloadImage(img)

	raylib.ImageResize(img, img.Width/2, img.Height/2)
	tile := raylib.LoadTextureFromImage(img)
	defer raylib.UnloadTexture(tile)

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.RayWhite)
		raylib.DrawTexture(tile, screenWidth/2-tile.Width/2, screenHeight/2-tile.Height/2, raylib.White)
		raylib.EndDrawing()
	}

}
