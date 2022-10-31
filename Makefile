.PHONY: build
build: 
	swig -c++ -intgosize 64 -go internal/cppsgp4/sgp4.i
