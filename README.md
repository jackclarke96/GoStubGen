# GoStubGen

## Overview

GoStubGen is a code generation tool that takes a YAML specification of an interface and its concrete types and generates a structured Go package. The generated package includes:

- Interface definitions
- Concrete implementations
- Mocking framework for unit testing

Designed for integration into a parent package using the dependency injection pattern, GoStubGen ensures full control over injected dependencies, making unit testing more reliable and maintainable.

## Features
- **Automatic interface and struct generation** based on YAML configuration
- **Mocking framework generation** based on YAML configuration
- **Dependency injection support** for flexible runtime behavior

## Installation

Ensure you have Go installed. Then, clone the repository and build the generator:

```sh
git clone https://github.com/jackclarke/GoStubGen.git
cd GoStubGen
go mod vendor
```

## Usage

### YAML Spec

```yaml
package_imports: # List of package imports
  - fmt
  - errors

package: <package_name> # The package name for the generated files
importer: <importer_package> # The package that will import this package

implementers:
  - name: <StructName>
    description: "Description of the struct"
    fields:
      - name: <FieldName>
        type: <FieldType>

custom_structs: # Optional, defines additional custom structs
  - name: <StructName>
    fields:
      - name: <FieldName>
        type: <FieldType>

name: <InterfaceName> # The name of the interface to generate
methods:
  - name: <MethodName>
    description: "Description of method"
    inputs:
      - name: <ParamName>
        type: <ParamType>
    outputs:
      - type: <ReturnType>
```

### Generating Code

Run the following command to generate Go code from a YAML specification:

```sh
go run main.go generate -c <path-to-yaml>
```

This will add 

### Example YAML Configuration

A sample YAML file defining a `Vehicle` interface is available in `/config/example.yaml`. 

Running:

```sh
go run main.go generate -c config/example.yaml
```

will generate 

- `imported/vehicle.go`: The interface definition
- `imported/car.go`: Concrete type implementation for Car
- `imported/bike.go`: Concrete type implementation for Bike
- `imported/custom_types.go`: Dependency injection setup
- `importer/mocks.go`: Mock implementation for unit testing

#### Example: Using Dependency Injection

In `/templates` folder is...

You can inject a `Vehicle` implementation into the `Driver` struct:

```go
import "github.com/jackclarke/GoStubGen/templates/vehicle"
import "github.com/jackclarke/GoStubGen/templates/driver"

func main() {
    car := vehicle.NewCar()
    d := driver.NewDriver(driver.WithVehicle(car))
    err := d.Drive()
    if err != nil {
        fmt.Println("Error driving:", err)
    }
}
```

#### Using the Mocking Framework

```go
import (
    "errors"
    "testing"
    "github.com/jackclarke/GoStubGen/templates/vehicle"
    "github.com/jackclarke/GoStubGen/templates/driver"
)

func TestDriverDriveWithMock(t *testing.T) {
    mockVeh := vehicleMock(vehicle.NewCar())
    d := driver.NewDriver(driver.WithVehicle(mockVeh))
    mockVeh.enableLoadCargoMock()
    mockVeh.setLoadCargoResponse(0, errors.New("mock error!"))

    err := d.Drive()
    if err == nil {
        t.Fatal("Expected an error but got none")
    }
}
```
