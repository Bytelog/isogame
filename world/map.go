package world

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	rm "github.com/gen2brain/raylib-go/raymath"
)

var _ = rm.Clamp

const (
	X_SIZE        = 16
	Y_SIZE        = 16
	Z_SIZE        = 64
	TILE_WIDTH    = 64
	TILE_HEIGHT   = 64
	TILE_DEPTH    = 96
	TILE_OFFSET_X = 75
	TILE_OFFSET_Y = 35
)

type World struct {
	Name  string
	Chunk Chunk
}

func New(name string) *World {
	return &World{
		Name: name,
	}
}

func (w *World) Center() rl.Vector3 {
	return w.Chunk.Center()
}

type Tile struct {
	ID uint16
}

type Chunk struct {
	Tiles [X_SIZE][Y_SIZE][Z_SIZE]*Tile
}

func (chunk *Chunk) Center() rl.Vector3 {
	return rl.Vector3{
		X_SIZE * TILE_WIDTH / 2.0,
		Y_SIZE * TILE_HEIGHT / 2.0,
		Z_SIZE * TILE_DEPTH / 2.0,
	}
}
