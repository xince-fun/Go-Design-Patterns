package abstract_factory

import "fmt"

const (
	LuxuryCarType   = 1
	FamiliarCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return &LuxuryCar{}, nil
	case FamiliarCarType:
		return &FamiliarCar{}, nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized", v)
	}
}
