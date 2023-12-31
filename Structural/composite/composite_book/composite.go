package composite

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

func Swim() {
	fmt.Println("Swimming!")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

type Trainer interface {
	Train()
}

type Swimmer interface {
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	fmt.Println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

type Animal struct{}

func (r *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p Parent) int {
	return p.SomeField
}
