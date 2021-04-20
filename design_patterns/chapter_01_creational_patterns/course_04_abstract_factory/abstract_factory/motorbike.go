package abstract_factory

import (
	"errors"
	"fmt"
)

type Motorbike interface {
	GetMotorbikeType() int
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct {
}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Motorbike of type %d not recognized\n", v))
	}
}
