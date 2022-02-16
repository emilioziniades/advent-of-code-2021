package day22

import (
	"github.com/fogleman/ln/ln"
)

func RenderCuboids(cuboids []cuboid) {
	scene := ln.Scene{}

	/* for _, c := range cuboids {
		start, end := cuboidToVectors(c)
		scene.Add(ln.NewCube(start, end))
	} */
	// start, end := cuboidToVectors(cuboids[0])
	// scene.Add(ln.NewCube(start, end))
	scene.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))
	scene.Add(ln.NewCube(ln.Vector{0, 0, 0}, ln.Vector{2, 2, 2}))

	// define camera parameters
	eye := ln.Vector{4, 3, 2}    // camera position
	center := ln.Vector{0, 0, 0} // camera looks at
	up := ln.Vector{0, 0, 1}     // up direction

	// define rendering parameters
	width := 1024.0  // rendered width
	height := 1024.0 // rendered height
	fovy := 50.0     // vertical field of view, degrees
	znear := 0.1     // near z plane
	zfar := 10.0     // far z plane
	step := 0.01     // how finely to chop the paths for visibility testing

	// compute 2D paths that depict the 3D scene
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	// render the paths in an image
	paths.WriteToPNG("out.png", width, height)
}

func cuboidToVectors(c cuboid) (ln.Vector, ln.Vector) {
	sX, sY, sZ := float64(c.start.x), float64(c.start.y), float64(c.start.z)
	eX, eY, eZ := float64(c.end.x), float64(c.end.y), float64(c.end.z)
	return ln.Vector{sX, sY, sZ}, ln.Vector{eX, eY, eZ}
}
