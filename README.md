# SGP4 for Go
SGP4 library port using SWIG

This library wrap the the [SGP4 C++ library](https://github.com/dnwrnr/sgp4), with handlers to catch C++ exceptions into Go errors.

## Update bindings

- Update the library in `internal/cppsgp4`
- Copy / paste the `sgp4.i` file
- Run `make`

## TODO
- [ ] Fix deprecated
- [x] Better readme

## Sources

cppsgp4: [https://github.com/dnwrnr/sgp4](https://github.com/dnwrnr/sgp4)
