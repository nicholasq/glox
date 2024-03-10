package main

type Interpreter struct {
	globals     interface{}
	environment interface{}
	locals      map[interface{}]int
}
