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
- **Spy tracking for method call inspection**
- **Support for capturing background method results**
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

Run the following command to generate Go code from a YAML specification:

```sh
go run main.go generate -c <path-to-yaml>
```

For example:

```sh
go run main.go generate -c examples/vehicle-example/vehicle_example.yaml
```

### YAML Spec

YAML describes the package, imports, interfaces, and implementers. For details
and examples, see [`examples/vehicle-example`](./examples/vehicle-example/).

## Dependency Injection Example

```go
car := vehicle.NewCar()
driver := driver.NewDriver(driver.WithVehicle(car))
err := driver.Drive()
```

## Mocking Framework

The generated mocks support:

| Method                                  | Description                         |
| --------------------------------------- | ----------------------------------- |
| `enable<Method>Mock()`                  | Enables mock override               |
| `disable<Method>Mock()`                 | Disables mock override              |
| `set<Method>Response(...)`              | Sets a static response              |
| `set<Method>Func(func)`                 | Custom behavior                     |
| `enqueue<Method>Response(...)`          | Queues responses in FIFO order      |
| `enqueue<Method>ResponseWithDelay(...)` | Queues response with optional delay |

> **Note:** If the mock is not enabled, the real method on the underlying
> concrete implementation is used. This allows selective mocking of only the
> methods relevant to the test.

### Additional Mock Helpers

These additional methods give more control over queued responses:

| Method                                          | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- |
| `enqueue<Method>ResponseFunc(func)`             | Enqueues a function-based response              |
| `enqueue<Method>ResponseFuncWithDelay(func, d)` | Enqueues a function-based response with a delay |

These are useful for dynamic logic testing or simulating delayed computation.

### Mocking Example

```go
car := vehicle.NewCar()
mock := vehicle.NewVehicleMock(car)

// This returns the real method's output
mock.Drive()

// Now mock it
mock.enableDriveMock()
mock.setDriveResponse("mocked!")

// This returns the mocked output
mock.Drive() // => "mocked!"
```

### Spying

Spy methods help inspect call history:

| Method                               | Description                                    |
| ------------------------------------ | ---------------------------------------------- |
| `enable<Method>Spy()`                | Start recording method calls                   |
| `disable<Method>Spy()`               | Stop recording method calls                    |
| `get<Method>Calls()`                 | Retrieve all recorded calls                    |
| `capture<Method>CallSpy(t, timeout)` | Async channel that emits once call is observed |

### Capturing Background Results

Capture method outputs in background goroutines:

| Method                    | Description                                    |
| ------------------------- | ---------------------------------------------- |
| `capture<Method>Result()` | Returns a channel that receives method results |

For multi-return methods, result types are wrapped in named structs:

```go
type mockSelfDrivingAccelerateResult struct {
    Output0 int
    Output1 error
}
```

Example:

```go
ch := mock.captureAccelerateResult()
go func() {
    mock.Accelerate(50, "km/h")
}()
result := <-ch
```

This supports background method testing with simple channel assertions.

---

For further examples and a complete walkthrough, see the `examples/` directory.
