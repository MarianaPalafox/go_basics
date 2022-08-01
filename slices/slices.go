package main

import (
	"fmt"
)

func main() {
	// Same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// Length
	fmt.Println(len(loons))

	fmt.Println("________________")
	// 0 indexing
	fmt.Println(loons[1])

	fmt.Println("________________")
	// slice everything from index 1 to the end
	fmt.Println(loons[1:])

	fmt.Println("________________")
	// print with for
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("________________")
	// Single value range
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("________________")
	// Double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("________________")
	// Double range ignore index by using _
	for _, name := range loons {
		fmt.Println(name)
	}

	fmt.Println("________________")
	// add to an slice with append
	loons = append(loons, "elmer")
	fmt.Println(loons)
}
