package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	rm "github.com/gen2brain/raylib-go/raymath"
	"isogame/world"
	"sort"
)

var _ = rm.Clamp
var _ = world.X_SIZE
var _ = fmt.Println
var screenV = rl.Vector3{SCREEN_W, SCREEN_H, 0.0}

func gameToScreen(pos rl.Vector3) rl.Vector2 {
	w := world.New("temp")
	center := world.VectorISO(w.Center())
	offset := world.VectorISO(pos)
	return rl.Vector2{
		X: center.X + 75 + offset.X,
		Y: -center.Y + 35 + offset.Y,
	}
}

/*
func ScreenToGame(pos rl.Vector2) rl.Vector3 {
	w := world.New("temp")
	center := world.VectorISO(w.Center())
	return rl.Vector3{}
}*/

func main() {
	rl.InitWindow(SCREEN_W, SCREEN_H, TITLE)
	rl.SetTargetFPS(FPS)
	defer rl.CloseWindow()

	texture := rl.LoadTexture("grass.png")
	defer rl.UnloadTexture(texture)

	w := world.New(TITLE)

	for i := 0; i < len(w.Chunk.Tiles); i++ {
		for j := 0; j < len(w.Chunk.Tiles[i]); j++ {
			for k := 0; k < len(w.Chunk.Tiles[i][j]); k++ {
				if k <= 8 {
					w.Chunk.Tiles[i][j][k] = &world.Tile{}
				}
			}
		}
	}

	w.Chunk.Tiles[8][8][8] = nil

	var buffer RenderBuffer

	for ix, tx := range w.Chunk.Tiles {
		for iy, ty := range tx {
			for iz, tile := range ty {
				if tile != nil {
					buffer = append(buffer, RenderTile{
						texture: texture,
						position: rl.Vector3{
							float32(ix * world.TILE_WIDTH),
							float32(iy * world.TILE_HEIGHT),
							float32(iz * world.TILE_DEPTH),
						},
					})
				}
			}
		}
	}

	sort.Stable(buffer)

	// CENTERING:
	// draw_start = screen center - chunk center
	zero := gameToScreen(rm.VectorZero())

	// tile offset: 75/35
	camera := rl.Camera2D{
		Target: zero,
		Offset: rl.Vector2{SCREEN_W/2 - zero.X, SCREEN_H/2 - zero.Y},
		//Offset: rl.Vector2{(SCREEN_W - center.X) - SCREEN_W/2 - 75, SCREEN_H/2 + center.Y - 35},
		Zoom: 1,
	}

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyRight) {
			camera.Offset.X -= 10 // Camera displacement with player movement
		} else if rl.IsKeyDown(rl.KeyLeft) {
			camera.Offset.X += 10 // Camera displacement with player movement
		}

		if rl.IsKeyDown(rl.KeyDown) {
			camera.Offset.Y -= 10 // Camera displacement with player movement
		} else if rl.IsKeyDown(rl.KeyUp) {
			camera.Offset.Y += 10 // Camera displacement with player movement
		}

		// Camera zoom controls
		camera.Zoom += float32(rl.GetMouseWheelMove()) * 0.05

		if camera.Zoom > 10.0 {
			camera.Zoom = 10.0
		} else if camera.Zoom < 0.01 {
			camera.Zoom = 0.01
		}

		// Camera reset (zoom and position)
		if rl.IsKeyPressed(rl.KeySpace) {
			camera.Zoom = 0.5
			camera.Offset = camera.Target
		}

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			fmt.Println(float32(rl.GetMouseX()), float32(rl.GetMouseY()))
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(93, 148, 241, 255))
		rl.Begin2dMode(camera)

		// Chunk Ortho
		for _, o := range buffer {
			v := o.Position()
			v.Z = -v.Z // Flip z because our axes directions are *$#!ed
			v = world.VectorISO(v)
			rl.DrawTexture(o.Texture(), int32(v.X), int32(v.Y), rl.White)
		}

		rl.DrawCircleV(camera.Target, 5, rl.Black)

		rl.End2dMode()

		// Ortho Axes
		rl.DrawLine(0, SCREEN_H/2, SCREEN_W, SCREEN_H/2, rl.NewColor(230, 41, 55, 128))
		rl.DrawLine(SCREEN_W/2, 0, SCREEN_W/2, SCREEN_H, rl.NewColor(230, 41, 55, 128))

		rl.EndDrawing()
	}
}

type RenderTile struct {
	texture  rl.Texture2D
	position rl.Vector3
}

func (r RenderTile) Position() rl.Vector3 {
	return r.position
}

func (r RenderTile) Texture() rl.Texture2D {
	return r.texture
}

type Renderable interface {
	Position() rl.Vector3
	Texture() rl.Texture2D
}

type RenderBuffer []Renderable

func (b RenderBuffer) Len() int {
	return len(b)
}

func (b RenderBuffer) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b RenderBuffer) Less(i, j int) bool {
	vi := world.VectorISO(b[i].Position())
	vj := world.VectorISO(b[j].Position())

	// Draw renderables from bottom to top of screen, iso coords (+y -> 0)
	return vi.Y < vj.Y
}
