package main

import (
	"fmt"
	"math"
)

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div= %d, mod= %d\n", div, mod)

	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)

	val2 := 10
	double(val2)
	fmt.Println(val2)

	doublePtr(&val2)
	fmt.Println(val2)

	s1, error := sqrt(-1)
	if error != nil {
		fmt.Printf("ERORR: %s\n", error)
	} else {
		fmt.Printf("%.2f", s1)
	}

	s2, error2 := sqrt(64)
	if error2 != nil {
		fmt.Printf("ERORR: %s\n", error2)
	} else {
		fmt.Printf("%.2f", s2)
	}

	//Defer
	worker()
}

// simple function definition
func add(a int, b int) int {
	return a + b
}

// function returns 2 values
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

// with parameters
func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

// using pointers
func doublePtr(n *int) {
	*n *= 2
}

// returning errors
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value (%.2f)", n)
	}

	return math.Sqrt(n), nil
}

func acquire(name string) (string, error) {
	return name, nil
}

func release(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	r1, error := acquire("A")
	if error != nil {
		fmt.Println("ERROR: ", error)
		return
	}
	defer release(r1) // defer is deferred function called, this will get called only when the worker function exits

	r2, error2 := acquire("B")
	if error2 != nil {
		fmt.Println("ERROR: ", error)
		return
	}
	defer release(r2) // if there are more than one defer they are called in reverse order

	fmt.Printf("\nworker\n")
}
