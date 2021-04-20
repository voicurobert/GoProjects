package composite

import "testing"

func TestComposite(t *testing.T) {
	var swimmer = CompositeSwimmer{MySwim: Swim}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	localSwim := Swim

	swimmer2 := CompositeSwimmer{MySwim: localSwim}
	swimmer2.MySwim()

	fish := Shark{Swim: Swim}
	fish.Swim()
	fish.Eat()

	swimmer3 := CompositeSwimmerB{
		Trainer: &Athelete{},
		Swimmer: &SwimmerImplementation{},
	}

	swimmer3.Train()
	swimmer3.Swim()
}
