/*
This finder your Mars age. 687 day is 1 year in the mars and 365 day is 1 year in the earth.
It is coculate your age between earht and mars age differenses.
*/
package main

import "fmt"

func mars_age(age int) int {
	var count, ey int
	ey = age * 365
	for ey > 687 {
		ey = ey - 687
		count++
	}
	return count
}

func main() {
	var age int
	fmt.Print("Inter age: ")
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Printf("Your Mars age: %d\n", mars)
}
