package driver

import (
	"fmt"

	"github.com/jackclarke/GoStubGen/templates/vehicle"
)

type Driver struct {
	vehicle vehicle.Vehicle
}

type Option func(*Driver)

// WithVehicle injects a vehicle into the Driver. TODO probably use init here.
func WithVehicle(v vehicle.Vehicle) Option {
	return func(c *Driver) {
		c.vehicle = v //.init()
	}
}

// NewDriver constructs a new Driver with the provided options.
func NewDriver(opts ...Option) *Driver {
	d := &Driver{}
	for _, opt := range opts {
		opt(d)
	}

	// default to car if option not provided.
	if d.vehicle == nil {
		d.vehicle = vehicle.NewCar()
	}
	return d
}

func (d *Driver) drive() error {
	cargo, err := d.vehicle.LoadCargo([]string{"clothes", "toiletries", "electronics"})

	if err != nil {
		return fmt.Errorf("drive: failed to load cargo! %w", err)
	}
	fmt.Printf("Cargo retrieved! %+v", cargo)
	return nil
}
