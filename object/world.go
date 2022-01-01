package object

import (
	"RayTracer/color"
	"RayTracer/matrix"
	"RayTracer/ray"
	"RayTracer/vector"
	"sort"
)

type World struct {
	Spheres []Sphere
	Lights  []Light
}

func CreateEmptyWorld() World {
	return World{}
}

func DefaultWorld() World {

	s1 := CreateSphere("s1")
	s1.SetMaterial(CreateMaterial(color.NewColor(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0))

	s2 := CreateSphere("s2")
	s2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))

	return CreateEmptyWorld().
		AddLight(CreatePointLight(vector.NewPoint(-10.0, 10.0, -10.0), color.NewColor(1.0, 1.0, 1.0))).
		AddSphere(s1).
		AddSphere(s2)
}

func (w World) AddSphere(sphere Sphere) World {
	return World{
		Spheres: append(w.Spheres, sphere),
		Lights:  w.Lights,
	}
}

func (w World) AddLight(light Light) World {
	return World{
		Spheres: w.Spheres,
		Lights:  append(w.Lights, light),
	}
}

func (w World) Intersect(ray ray.Ray) []Intersection {

	var intersections []Intersection

	for _, sphere := range w.Spheres {
		intersections = append(intersections, sphere.Intersect(ray)...)
	}

	sort.SliceStable(intersections, func(i, j int) bool {
		return intersections[i].GetTime() < intersections[j].GetTime()
	})

	return intersections
}
