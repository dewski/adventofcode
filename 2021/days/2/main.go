package main

import (
	"fmt"

	"github.com/dewski/adventofcode/2021/inputs"
)

type PositionalSubmarine struct {
	X int
	Y int
}

func (ps *PositionalSubmarine) HandleAction(courseAction inputs.CourseAction) {
	switch courseAction.Action {
	case inputs.DownAction:
		ps.Y += courseAction.Value
	case inputs.UpAction:
		ps.Y -= courseAction.Value
	case inputs.ForwardAction:
		ps.X += courseAction.Value
	}
}

type DirectionalSubmarine struct {
	X   int
	Y   int
	Aim int
}

func (ds *DirectionalSubmarine) HandleAction(courseAction inputs.CourseAction) {
	switch courseAction.Action {
	case inputs.DownAction:
		ds.Aim += courseAction.Value
	case inputs.UpAction:
		ds.Aim -= courseAction.Value
	case inputs.ForwardAction:
		ds.X += courseAction.Value
		ds.Y += courseAction.Value * ds.Aim
	}
}

func main() {
	ps := PositionalSubmarine{
		X: 0,
		Y: 0,
	}

	for _, courseAction := range inputs.DayTwoCourseActions {
		ps.HandleAction(courseAction)
	}

	fmt.Printf("Horizontal position: %d\n", ps.X)
	fmt.Printf("Depth: %d\n", ps.Y)
	fmt.Printf("Answer: %d\n", ps.X*ps.Y)

	ds := DirectionalSubmarine{
		X:   0,
		Y:   0,
		Aim: 0,
	}

	for _, courseAction := range inputs.DayTwoCourseActions {
		ds.HandleAction(courseAction)
	}

	fmt.Printf("Horizontal position: %d\n", ds.X)
	fmt.Printf("Depth: %d\n", ds.Y)
	fmt.Printf("Answer: %d\n", ds.X*ds.Y)
}
