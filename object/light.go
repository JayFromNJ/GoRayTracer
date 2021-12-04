package object

import (
	"RayTracer/color"
	"RayTracer/vector"
)

type Light struct {
	position  vector.Vector
	intensity color.Color
}

func (l *Light) Position() vector.Vector { return l.position }
func (l *Light) Intensity() color.Color  { return l.intensity }

func CreatePointLight(pos vector.Vector, int color.Color) Light {
	return Light{
		position:  pos,
		intensity: int,
	}
}
