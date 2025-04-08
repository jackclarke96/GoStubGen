# GoStubGen

## Overview

GoStubGen is a code generation tool that takes a YAML specification of an
interface and its concrete types and generates a structured Go package. The
generated package includes:

- Interface definitions
- Concrete implementations
- Mocking framework for unit testing
- Support for embedded interfaces and structs

Designed for integration into a parent package using the dependency injection
pattern, GoStubGen ensures full control over injected dependencies, making unit
testing more reliable and maintainable.

## Features

- **Automatic interface and struct generation** based on YAML configuration
- **Mocking framework generation** for unit testing
- **Dependency injection support** for flexible runtime behavior

## Installation

Ensure you have Go installed. Then, clone the repository and build the
generator:

```sh
git clone https://github.com/jackclarke/GoStubGen.git
cd GoStubGen
go mod vendor
```

## Usage

TODO example and split up yaml spec

### YAML Spec

For each interface and struct defined, the generator will take the "embeds" and
"implements" field values to calculate the minimum set of methods a struct
requires in order to implement a set of interfaces.

#### Example: Interface Inheritance

```yaml
interfaces:
  - name: Vehicle
    methods:
      - name: Drive
        outputs:
          - type: string

  - name: Car
    embeds: [Vehicle]
    methods:
      - name: OpenTrunk
        outputs:
          - type: error
```

The Car interface now requires:

- `Drive()` from Vehicle
- `OpenTrunk()` from itself

```go
type Car interface {
	Vehicle
	OpenTrunk() error
}
```

#### Example: Struct Implements Combined Interfaces

```yaml
implementers:
  - name: Sedan
    implements: [Car]
    fields:
      - name: color
        type: string
```

The generator knows Sedan must implement:

- `Drive()` string
- `OpenTrunk()` error

```go
func (s *Sedan) Drive() string     { return "" }
func (s *Sedan) OpenTrunk() error  { return nil }
```

#### Example: Struct Implements Combined Interfaces and also embeds Struct

```yaml
interfaces:
  - name: SelfDriving
    embeds: [Car]
    methods:
      - name: ActivateAutopilot
        outputs:
          - type: error

implementers:
  - name: RoboCar
    implements: [SelfDriving]
    embeds: [Sedan]
    fields:
      - name: aiVersion
        type: string
```

then `RoboCar` must implement only the unique method:

- `ActivateAutopilot() error`

Because it embeds Sedan, which already provides:

- `Drive() string`
- `OpenTrunk() error`

#### Full YAML example

```yaml
package_imports: # List of package imports
  - fmt
  - errors

package: <package_name> # The package name for the generated files
importer: <importer_package> # The package that will import this package

# Custom Structs are optional. Reusable types used in method signatures or within interfaces.
custom_structs: # Optional, defines additional custom structs
  - name: <StructName>
    description: "Destription of the struct" # will appear as comment
    fields:
      - name: <FieldName>
        description: "Description of the field" # will appear as comment
        type: <FieldType>


# Each interface can include methods and embed other interfaces.
name: <InterfaceName> # The name of the interface to generate
  embedded: ["<other interfaces>"]
  description: "Interface description" # will appear as comment
  methods:
    - name: <MethodName>
      description: "Description of method" # will appear as comment
      inputs: # omit for methods with none
        - name: <ParamName>
          type: <ParamType>
      outputs: # omit for methods with none
        - type: <ReturnType>

# Implementers declare concrete structs that implement one or more interfaces.
implementers:
  - name: <StructName>
    description: "Description of the struct" # will appear as comment
    embedded: ["<other structs>"]
    implements: ["<list of interfaces>"]
    fields:
      - name: <FieldName>
        type: <FieldType>
```

### Generating Code

Run the following command to generate Go code from a YAML specification:

```sh
go run main.go generate -c <path-to-yaml>
```

For example, using:

```sh
go run main.go generate -c /examples/vehicle-example/vehicle_example.yaml
```

will generate the following files:

- **`vehicle/vehicle.go`** – Defines the `Vehicle` interface.
- **`vehicle/car.go`** – Implements `Vehicle` for a `Car` type.
- **`vehicle/bike.go`** – Implements `Vehicle` for a `Bike` type.
- **`vehicle/custom_types.go`** – Provides dependency injection setup.
- **`importer/vehicle_mock_test.go`** – Generates a mock implementation for unit
  testing.
- **`importer/self_driving_mocktest.go`** - Generates another mock
  implementation for unit testing

These files allow for easy integration into a larger Go project with dependency
injection support.

## Example: Using Dependency Injection

In `examples/vehicle-example` is a package built using the dependency injection
method encouraged by this pakage. You can inject a `Vehicle` implementation into
the `Driver` struct:

```go
import (
    "github.com/jackclarke/GoStubGen/templates/vehicle"
    "github.com/jackclarke/GoStubGen/templates/driver"
    "fmt"
)

func main() {
    // Create car: a concrete implementation of Vehicle
    car := vehicle.NewCar()
    // Inject car into driver package
    d := driver.NewDriver(driver.WithVehicle(car))
    // Invoke Drive method which uses Vehicle methods
    err := d.Drive()
    if err != nil {
        fmt.Println("Error driving:", err)
    }
}
```

This enables seamless swapping of implementations for testing and production.

## Using the Mocking Framework

GoStubGen provides a built-in mocking system that allows you to override methods
dynamically.

The generated mock struct offers the following methods:

| Method                               | Description                             |
| ------------------------------------ | --------------------------------------- |
| `enable<MethodName>Mock()`           | Enables the mock for the method.        |
| `disable<MethodName>Mock()`          | Restores original functionality.        |
| `set<MethodName>Response(...)`       | Defines a static response for the mock. |
| `set<MethodName>Func(func(...) ...)` | Allows complete method override.        |

### Mocking Example

This example shows how to mock a method in a test:

```go
import (
    "errors"
    "testing"
    "github.com/jackclarke/GoStubGen/templates/vehicle"
)

func TestDriverDriveWithMock(t *testing.T) {
    // Create a mock of a Car (which implements Vehicle)
    mockVeh := newVehicleMock(vehicle.NewCar())
    // inject mock Vehicle into driver
    d := driver.NewDriver(driver.WithVehicle(mockVeh))
    // Mock the LoadCargo method to return an error
    mockVeh.enableLoadCargoMock()
    mockVeh.setLoadCargoResponse(0, errors.New("mock error!"))

    err := d.Drive()
    if err == nil {
        t.Fatal("Expected an error but got none")
    }
}
```

Before mocking:

```go
car := vehicle.NewCar()
car.LoadCargo(10)  // Works as expected
```

After mocking:

```go
mockVeh := vehicleMock(vehicle.NewCar())
mockVeh.enableLoadCargoMock()
mockVeh.setLoadCargoResponse(0, errors.New("mock error!"))

err := mockVeh.LoadCargo(10)  // Always returns an error
fmt.Println(err)  // Output: "mock error!"
```
