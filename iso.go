package main

import "github.com/gen2brain/raylib-go/raylib"

func VectorISO(v *raylib.Vector3) {
	x := v.X
	v.X = v.X + v.Y
	v.Y = (v.Y-x)/2 + v.Z
}

func VectorOrtho(v *raylib.Vector3) {
	x := v.X
	v.X = v.X/2 - v.Y + v.Z
	v.Y = v.Y + v.X/2 - v.Z
}
