package world

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func VectorISO(v rl.Vector3) rl.Vector3 {
	return rl.Vector3{
		X: v.X + v.Y,
		Y: (v.Y-v.X)/2 + v.Z,
		Z: v.Z,
	}
}

func VectorOrtho(v rl.Vector3) rl.Vector3 {
	return rl.Vector3{
		X: v.X/2 - v.Y + v.Z,
		Y: v.Y + v.X/2 - v.Z,
		Z: v.Z,
	}
}
