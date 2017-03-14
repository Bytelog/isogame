package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
	"sort"
)

func main() {
	screenWidth := int32(1300)
	screenHeight := int32(700)

	tileWidth := int(32) // effective size after scaling by 0.5x,
	tileHeight := int(32)
	tileDepth := int(48)

	raylib.InitWindow(screenWidth, screenHeight, "ISOGAME")
	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	camera := raylib.Camera2D{
		Target:   raylib.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		Offset:   raylib.NewVector2(0, 0),
		Rotation: 0.0,
		Zoom:     0.5,
	}

	img := raylib.LoadImage("grass.png")
	defer raylib.UnloadImage(img)

	raylib.ImageResize(img, img.Width/2, img.Height/2)
	texture := raylib.LoadTextureFromImage(img)
	defer raylib.UnloadTexture(texture)

	world := Map{
		Name:    "World",
		Objects: make([]Object, 0),
		Tiles:   makeTiles(32, 32, 3),
	}

	var buffer RenderBuffer

	for ix, tx := range world.Tiles {
		for iy, ty := range tx {
			for _, t := range ty {
				if t.Enabled {
					buffer = append(buffer, RenderTile{
						texture: texture,
						position: raylib.NewVector3(
							float32(ix*tileWidth),
							float32(iy*tileHeight),
							t.Z*float32(tileDepth),
						),
					})
				}
			}
		}
	}

	sort.Stable(buffer)

	for !raylib.WindowShouldClose() {
		if raylib.IsKeyDown(raylib.KeyRight) {
			camera.Offset.X -= 10 // Camera displacement with player movement
		} else if raylib.IsKeyDown(raylib.KeyLeft) {
			camera.Offset.X += 10 // Camera displacement with player movement
		}

		if raylib.IsKeyDown(raylib.KeyDown) {
			camera.Offset.Y -= 10 // Camera displacement with player movement
		} else if raylib.IsKeyDown(raylib.KeyUp) {
			camera.Offset.Y += 10 // Camera displacement with player movement
		}

		// Camera zoom controls
		camera.Zoom += float32(raylib.GetMouseWheelMove()) * 0.05

		if camera.Zoom > 3.0 {
			camera.Zoom = 3.0
		} else if camera.Zoom < 0.1 {
			camera.Zoom = 0.1
		}

		// Camera reset (zoom and position)
		if raylib.IsKeyPressed(raylib.KeySpace) {
			camera.Zoom = 1.0
			camera.Offset = camera.Target
		}

		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.RayWhite)
		raylib.Begin2dMode(camera)

		for _, o := range buffer {
			v := o.Position()
			v.Z = -v.Z // Flip z because our axes directions are *$#!ed
			VectorISO(&v)
			raylib.DrawTexture(o.Texture(), int32(v.X), int32(v.Y), raylib.White)
		}

		raylib.End2dMode()
		raylib.EndDrawing()
	}

}

type Map struct {
	Name    string
	Objects []Object
	Tiles   [][][]Tile
}

type Tile struct {
	Z       float32
	Class   uint16
	Enabled bool
}

type RenderTile struct {
	texture  raylib.Texture2D
	position raylib.Vector3
}

func (r RenderTile) Position() raylib.Vector3 {
	return r.position
}

func (r RenderTile) Texture() raylib.Texture2D {
	return r.texture
}

type Object struct {
	Position raylib.Vector3
}

type Renderable interface {
	Position() raylib.Vector3
	Texture() raylib.Texture2D
}

type RenderBuffer []Renderable

func (b RenderBuffer) Len() int {
	return len(b)
}

func (b RenderBuffer) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b RenderBuffer) Less(i, j int) bool {
	vi := b[i].Position()
	vj := b[j].Position()

	// Draw renderables from bottom to top of screen, iso coords (+y -> 0)
	VectorISO(&vi)
	VectorISO(&vj)
	return vi.Y < vj.Y
}

func makeTiles(x, y, z int) [][][]Tile {
	tiles := make([][][]Tile, x)
	for i := 0; i < x; i++ {
		tiles[i] = make([][]Tile, y)
		for j := 0; j < y; j++ {
			tiles[i][j] = make([]Tile, 0, z)

			// TODO: Procedurally generate some real way
			tiles[i][j] = append(tiles[i][j], Tile{
				Z:       500 / float32(math.Pow(float64(i-x/2), 2)+math.Pow(float64(j-y/2), 2)),
				Class:   0,
				Enabled: true,
			})
		}
	}
	return tiles
}
