package composite

import "fmt"

type Athelete struct {
}

func (a *Athelete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmer struct {
	MyAthlete Athelete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swim")
}

type Animal struct {
}

func (animal *Animal) Eat() {
	fmt.Println("Eat")
}

type Shark struct {
	Animal
	Swim func()
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImplementation struct {
}

func (s *SwimmerImplementation) Swim() {
	fmt.Println("Swimming")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}
