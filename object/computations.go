package object

import (
	"RayTracer/color"
	"RayTracer/ray"
	"RayTracer/vector"
)

type Computations struct {
	t       float64
	hitObj  HittableObject
	point   vector.Vector
	eyev    vector.Vector
	normalv vector.Vector
	inside  bool
}

func (c *Computations) GetTime() float64                  { return c.t }
func (c *Computations) GetHittableObject() HittableObject { return c.hitObj }
func (c *Computations) GetPoint() vector.Vector           { return c.point }
func (c *Computations) GetEyeV() vector.Vector            { return c.eyev }
func (c *Computations) GetNormalV() vector.Vector         { return c.normalv }
func (c *Computations) IsInside() bool                    { return c.inside }

func PrepareComputations(intersection Intersection, ray ray.Ray) Computations {
	time := intersection.GetTime()
	hitObj := intersection.ho
	point := ray.Position(intersection.GetTime())
	eyev := vector.Negate(ray.Direction())
	normalv := hitObj.NormalAt(point)

	inside := true
	if vector.Dot(eyev, normalv) < 0.0 {
		normalv = vector.Negate(normalv)
	} else {
		inside = false
	}

	return Computations{
		t:       time,
		hitObj:  hitObj,
		point:   point,
		eyev:    eyev,
		normalv: normalv,
		inside:  inside,
	}
}

func ShadeHit(world World, comps Computations) color.Color {
	return Lighting(comps.hitObj.GetMaterial(), world.Light,
		comps.GetPoint(),
		comps.GetEyeV(),
		comps.GetNormalV())
}
