package main

import "fmt"

type circle struct {
	radius int
}

type rectangle struct {
	length, width int
}

func (c circle) area() int {
	return c.radius ^ 2
}

func (c circle) perim() int {
	return 2 * (22 / 7) * c.radius
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (r rectangle) perim() int {
	return 2 * (r.length + r.width)
}

type shape interface {
	area() int
	perim() int
}

func main() {
	fmt.Println("Testing interface knowledge in Golang.")
	c1 := circle{5}
	r1 := rectangle{5, 6}
	shapes := []shape{c1, r1}
	for _, s := range shapes {
		fmt.Println(s.area(), s.perim())
	}
}
