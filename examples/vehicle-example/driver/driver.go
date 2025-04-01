package driver

import (
	"fmt"

	"github.com/jackclarke/GoStubGen/examples/vehicle-example/vehicle"
)

type Driver struct {
	vehicle vehicle.Vehicle
}

type Option func(*Driver)

// WithVehicle injects a vehicle into the Driver. Pass an initialised vehicle here
func WithVehicle(v vehicle.Vehicle) Option {
	return func(c *Driver) {
		c.vehicle = v
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

// basic driver method to enable test of mocking
func (d *Driver) drive() error {
	cargo, err := d.vehicle.LoadCargo([]string{"clothes", "toiletries", "electronics"})
	if err != nil {
		return fmt.Errorf("drive: failed to load cargo! %w", err)
	}
	fmt.Printf("%+v pieces of cargo retrieved!", cargo)
	return nil
}

// more complex driver method involving startup of background routine

// func (d *Driver) soemthing() {
// 	defer d.vehicle.ApplyBrakes()
// }
