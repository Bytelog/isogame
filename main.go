package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"sort"
)

func main() {
	screenWidth := int32(1300)
	screenHeight := int32(700)

	//tileWidth := int(34) // effective size after scaling by 0.5x, orthogonal
	//tileHeight := int(34)
	//tileDepth := int(82) - tileHeight // 82 (full height) - tileHeight

	//worldWidth := int(32)
	//worldHeight := int(32)
	//worldDepth := int(3)

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

	/*world := Map{
		Name:    "World",
		Objects: make([]Object, 0),
		Tiles:   makeTiles(worldWidth, worldHeight, worldDepth),
	}*/

	var buffer RenderBuffer

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			pos := raylib.NewVector3(float32(i*32), float32(j*32), 0)
			buffer = append(buffer, RenderTile{
				texture:  texture,
				position: pos,
			})
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

	if vi.Z != vj.Z {
		return vi.Z < vj.Z
	}

	return vi.Y-vi.X < vj.Y-vj.X
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
				Enabled: true,
			}

			if (i+j)%2 == 0 {
				tiles[i][j][1] = Tile{
					Class:   0,
					Enabled: true,
				}
			}
			if (i+j)%2 == 1 {
				tiles[i][j][2] = Tile{
					Class:   0,
					Enabled: true,
				}
			}
		}
	}
	return tiles
}
