package abstract_factory

import "fmt"

type VehicleFactory interface {
	GetVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return &CarFactory{}, nil
	case MotorbikeFactoryType:
		return &MotorbikeFactory{}, nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)
	}
}
