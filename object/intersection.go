package object

import (
	"RayTracer/vector"
	"math"
	"sort"
)

type Intersection struct {
	Object
	t float64
}

func (i *Intersection) GetTime() float64                 { return i.t }
func (i *Intersection) GetObjectID() string              { return i.id }
func (i *Intersection) GetObjectPosition() vector.Vector { return i.position }

func CreateIntersection(obj Object, t float64) Intersection {
	return Intersection{
		Object: Object{
			id:        obj.id,
			position:  obj.position,
			transform: obj.transform,
			material:  obj.material,
		},
		t: t,
	}
}

func NullIntersection() Intersection {
	return Intersection{
		Object: Object{
			id:       "null",
			position: vector.NewPoint(0.0, 0.0, 0.0),
		},
		t: math.MaxFloat64,
	}
}

func Equals(i1 Intersection, i2 Intersection) bool {
	if i1.id != i2.id {
		return false
	}
	if i1.position != i2.position {
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
