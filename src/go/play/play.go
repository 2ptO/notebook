package main

import (
	"errors"
	"fmt"
)

func nextInt() func() int {
	i := 1
	return func() int {
		i++
		return i
	}
}

func fib(n int) int {
	// beware of recursion. You may run into stack overflow
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibIter(n int) int {
	//iterative fibonacci
	a, b, c := 0, 1, 1
	for i := 0; i < n; i++ {
		a = b
		b = c
		c = a + b
	}
	return a
}

func addToList(list []int, n int) {
	//add the new number to the list.
	list = append(list, n)

}

func addToListAndReturn(list []int, n int) []int {
	list = append(list, n)
	return list
}

func addToListPtr(list *[]int, n int) {
	//pList := *list
	*list = append(*list, n)

}

// Identifiers starting in Uppercase are exported by default in Go.
type engine struct {
	rpm        int
	bhp        float64
	nCylinders int //number of cylinders
}

func (e engine) String() string {
	return fmt.Sprintf("%d, %fbhp", e.rpm, e.bhp)
}

func (e engine) raiseRPM(newVal int) {
	e.rpm += newVal
}

/*
 * Pass by pointer avoids copying the struct values
 * and also helps to mutate the values of the struct
 */
func (e *engine) raiseRPMPtr(newVal int) {
	e.rpm += newVal
}

func (e *engine) setCylinderCount(newCount int) error {
	if newCount < 4 || newCount > 10 {
		return errors.New("cyclinder count not in the range")
	}
	e.nCylinders = newCount
	return nil

}

const maxBHP = 50.0
const minBHP = 10.0

type bhpError struct {
	newBHP float64
	reason string
}

func (be bhpError) Error() string {
	return fmt.Sprintf("Can't set %.2f as bhp. Reason: %s", be.newBHP, be.reason)
}

func (e *engine) setBHP(newBHP float64) error {
	if newBHP < minBHP {
		return &bhpError{newBHP, "bhp too low"}
	}

	if newBHP > maxBHP {
		return &bhpError{newBHP, "bhp too high"}
	}

	e.bhp = newBHP
	return nil
}

func main() {
	fmt.Println("Hello, playground")
	n := nextInt()

	fmt.Println("Closures:")
	for i := 0; i < 5; i++ {
		fmt.Println(n())
	}

	fmt.Println("Recursion:")
	fmt.Println("Fibonacci(5) =", fib(30))

	fmt.Println("Pointers:")
	odds := []int{1, 3, 5}
	fmt.Println(odds)
	nextOdd := 7
	fmt.Printf("Adding %d into %v\n", nextOdd, odds)
	addToListPtr(&odds, nextOdd)
	fmt.Println(odds)

	fmt.Println("Structs:")
	altEngine := engine{15000, 40.5, 4}
	fmt.Println(altEngine)

	fmt.Println("Methods:")
	fmt.Println("Method by value:")

	fmt.Println("Before :")
	fmt.Println(altEngine)
	fmt.Println("After calling raiseRPM:")
	altEngine.raiseRPM(1000)
	fmt.Println(altEngine)

	fmt.Println("Before :")
	fmt.Println(altEngine)
	fmt.Println("After calling raiseRPMPtr:")
	altEngine.raiseRPMPtr(1000)
	fmt.Println(altEngine)

	fmt.Println("Errors:")
	fmt.Println("Using inbuilt errors package:")
	if err := altEngine.setCylinderCount(20); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Using custom error:")
	if err := altEngine.setBHP(10.5); err != nil {
		fmt.Println(err)
	}

	if err := altEngine.setBHP(5.5); err != nil {
		fmt.Println(err)
		fmt.Println("Extracting the fields from the custom error")
		bhpErr := err.(*bhpError) //type casting err to bhpError
		fmt.Println("value:", bhpErr.newBHP, "Reason:", bhpErr.reason)
	}
}
