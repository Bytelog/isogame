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

type Map struct {
	Name    string
	Objects []Object
	Tiles   [][][]Tile
}

type Tile struct {
	Class   uint16
	Enabled bool
}

type Object struct {
	Position raylib.Vector3
}

func makeTiles(x, y, z int) [][][]Tile {
	tiles := make([][][]Tile, x)
	for i := 0; i < x; i++ {
		tiles[i] = make([][]Tile, y)
		for j := 0; j < y; j++ {
			tiles[i][j] = make([]Tile, z)

			// TODO: Procedurally generate some real way
			tiles[i][j][0] = Tile{
				Class:   0,
				Enabled: 0,
			}
		}
	}
	return tiles
}
