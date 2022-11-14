package main

import "fmt"


type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}


type square struct {
	sideLenght float64
}

func main() {

	t := triangle { height: 5, base: 2, }
	s := square { sideLenght: 10}

	fmt.Println(getArea(t))
	fmt.Println(getArea(s))
	
}

func getArea(s shape) float64 {
	return s.getArea()
}

func (t triangle) getArea() float64 {
	return 0.5 * (t.base * t.height)
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}