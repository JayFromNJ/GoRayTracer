package main

func main() {

}

/*
import (
	"RayTracer/canvas"
	"RayTracer/color"
	"RayTracer/vector"
	"fmt"
	"math"
)

type Environment struct {
	Gravity vector.Vector
	Wind vector.Vector
}

type Projectile struct {
	Velocity vector.Vector
	Current vector.Vector
}

func (proj *Projectile) tick(env Environment) {
	proj.Current = vector.Add(proj.Current, proj.Velocity)
	proj.Velocity = vector.Add(proj.Velocity, vector.Add(env.Gravity, env.Wind))
}

func main() {

	env := Environment{
		Gravity: vector.NewVector(0.0, -0.1, 0.0),
		Wind: vector.NewVector(-0.01, 0.0, 0.0),
	}

	proj := Projectile{
		Velocity: vector.Normalize(vector.NewVector(1.0, 1.8, 0.0)),
		Current: vector.NewPoint(1.0, 1.0, 0.0),
	}

	proj.Velocity = vector.Multiply(proj.Velocity, 11.25)

	c := canvas.NewCanvas(900, 550)

	red := color.NewColor(1.0, 0.0, 0.0)

	for proj.Current.Y() > 0 {
		c.WritePixel(int(math.Round(proj.Current.X())), c.Height() - int(math.Round(proj.Current.Y())), red)

		fmt.Println(proj.Current.ToString())
		proj.tick(env)
	}

	canvas.ToPPM(c, "traj")

}*/
