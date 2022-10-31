.PHONY: build
build: 
	swig -c++ -intgosize 64 -go cppsgp4/sgp4.i
