package driver

import (
	"errors"
	"fmt"
	"log"
	"sync"

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
func (d *Driver) drive() (int, error) {
	cargo, err := d.vehicle.LoadCargo([]string{"clothes", "toiletries", "electronics"})
	if err != nil {
		return 0, fmt.Errorf("drive: failed to load cargo! %w", err)
	}
	fmt.Printf("%+v pieces of cargo retrieved!", cargo)
	cargo, err = d.vehicle.LoadCargo([]string{"clothes", "toiletries", "electronics", "more stuff"})
	if err != nil {
		return 0, fmt.Errorf("drive: failed to load second batch of cargo! %w", err)
	}
	fmt.Printf("%+v pieces of cargo retrieved!", cargo)
	return cargo, nil
}

func (d *Driver) setVehicleDriveSelf() error {
	if sd, ok := d.vehicle.(vehicle.SelfDriving); ok {
		return sd.DriveSelf("Paris")
	}
	return errors.New("vehicle does not support self-driving")
}

func (d *Driver) instructSelfDriver(location, destination string) (string, error) {

	sd, ok := d.vehicle.(vehicle.SelfDriving)
	if !ok {
		return location, errors.New("vehicle does not support self-driving")
	}

	defer func() {
		if err := sd.ParkSelf(); err != nil {
			fmt.Printf("instructSelfDriver ParkSelf error: %s", err)
		} else {
			fmt.Printf("parked self in %s", location)
		}
	}()

	if err := sd.DriveSelf(destination); err != nil {
		return "", err
	}

	location = destination
	return location, nil
}

func (d *Driver) instructSelfDriverWithCleanup(location, destination string) (string, error) {
	sd, ok := d.vehicle.(vehicle.SelfDriving)
	if !ok {
		return location, errors.New("vehicle does not support self-driving")
	}

	defer func() {
		var wg sync.WaitGroup

		criticalTasks := []struct {
			name string
			fn   func() error
		}{
			{"ParkSelf", sd.ParkSelf},
			{"LockDoors", sd.LockDoors},
		}

		// wait for these
		for _, task := range criticalTasks {
			wg.Add(1)
			go func(name string, fn func() error) {
				defer wg.Done()
				if err := fn(); err != nil {
					log.Printf("%s error: %v", name, err)
				} else {
					log.Printf("%s complete", name)
				}
			}(task.name, task.fn)
		}

		// Fire-and-forget
		bgTasks := []struct {
			name string
			fn   func() error
		}{
			{"TurnOffMusic", sd.TurnOffMusic},
			{"CloseWindows", sd.CloseWindows},
			{"TurnOffAC", sd.TurnOffAC},
		}

		for _, task := range bgTasks {
			go func(name string, fn func() error) {
				if err := fn(); err != nil {
					log.Printf("%s failed: %v", name, err)
				} else {
					log.Printf("%s complete", name)
				}
			}(task.name, task.fn)
		}

		wg.Wait() // Only blocks on ParkSelf + LockDoors
	}()

	if err := sd.DriveSelf(destination); err != nil {
		return "", err
	}

	location = destination
	return location, nil
}
