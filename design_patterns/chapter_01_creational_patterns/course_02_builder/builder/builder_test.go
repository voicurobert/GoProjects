package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}

	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they are %d \n", car.Wheels)
		return
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' but it is %s \n", car.Structure)
		return
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 but they are %d \n", car.Seats)
		return
	}

	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.Build()

	if motorbike.Seats != 2 {
		t.Errorf("Seats on a motorbike must be 2 and they are %d \n", motorbike.Seats)
		return
	}

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they are %d \n", motorbike.Wheels)
		return
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' but it is %s \n", motorbike.Structure)
		return
	}
}
