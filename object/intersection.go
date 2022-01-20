package object

import (
	"RayTracer/vector"
	"math"
	"sort"
)

type Intersection struct {
	ho HittableObject
	t  float64
}

func (i *Intersection) GetTime() float64                 { return i.t }
func (i *Intersection) GetObjectID() string              { return i.ho.GetObjectID() }
func (i *Intersection) GetObjectPosition() vector.Vector { return i.ho.GetPosition() }

func CreateIntersection(ho HittableObject, t float64) Intersection {
	return Intersection{
		ho: ho,
		t:  t,
	}
}

func NullIntersection() Intersection {
	nullSphere := NullSphere()
	return Intersection{
		ho: &nullSphere,
		t:  math.MaxFloat64,
	}
}

func IntersectionEquals(i1 Intersection, i2 Intersection) bool {
	if i1.ho.GetObjectID() != i2.ho.GetObjectID() {
		return false
	}
	if i1.ho.GetPosition() != i2.ho.GetPosition() {
		return false
	}
	if i1.t != i2.t {
		return false
	}
	return true
}

func Hit(xs []Intersection) Intersection {
	if len(xs) < 1 {
		return NullIntersection()
	}
	if len(xs) == 1 {
		if xs[0].t < 0.0 {
			return NullIntersection()
		} else {
			return xs[0]
		}
	}

	var after []Intersection

	for i := 0; i < len(xs); i++ {
		if xs[i].t >= 0.0 {
			after = append(after, xs[i])
		}
	}

	if len(after) < 1 {
		return NullIntersection()
	}
	if len(after) == 1 {
		return after[0]
	}

	sort.SliceStable(after, func(i, j int) bool {
		return after[i].t < after[j].t
	})

	return after[0]
}
