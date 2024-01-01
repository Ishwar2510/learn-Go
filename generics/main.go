//The genrics ate used to accept any dataTypes;

package main

import "fmt"

// step 1 add a non genrics function
func SumInts(m map[string]int64) (int64, error) {
	var result int64

	for key, value := range m {
		fmt.Println(key, value)
		result += value
	}
	return result, nil

}
func SumFloats(m map[string]float64) (float64, error) {
	var result float64
	for key, value := range m {
		fmt.Println(key, value)
		result += value
	}
	return result, nil
}
func main() {

	ints := map[string]int64{
		"first":  20,
		"second": 40,
	}
	floats := map[string]float64{
		"first":  10.50,
		"second": 20.60,
	}
	fmt.Println(SumInts(ints))
	fmt.Println(SumFloats(floats))
	// the above func has be re written in generic form
	// but know while calling u need to define the type u are sending so that
	// the function maps K and V to now its type
	fmt.Println(SumIntsOrFloats[string, int64](ints))
	fmt.Println(SumIntsOrFloats[string, float64](floats))

	// Now in  any cases u can omit  this type declaration w hile calling the code
	// because the compiler can infer this
	// but this is not possible always.
	// if you need to call a function which has no argumnets
	fmt.Println()
	fmt.Println()
	fmt.Println(SumIntsOrFloats(ints))
	fmt.Println(SumIntsOrFloats(floats))

	// moveing type declartion as interface so that it cam be reused by other functiuon
	fmt.Println(SumNumbers(ints))
	fmt.Println(SumNumbers(floats))

}

//convert above code To generics;

//Writing generic function to handle all data types

// The function need to declare what type it supports
// And the calling function need to define the type while its calling
// in order to do thus we need to define type parameter to the function
// each type parameter has a type constraints that act as a  kind of meta-type
// each type constraints specifies the permissiable type arguments
//

// the comparable constraints is pre declared in go
// It allows any type whose  value may be used as an operand of the comparison operators
// == and !=
// Go requires that map key to be comparable , so declaring k as Comparable  is
// necessary so you can use K as the key in the map .
// Using | in the V specifies the union of the two types , meanin this type allows
// either types/
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var result V
	for _, v := range m {
		result += v
	}
	return result
}

// Declare a  type constraints
// now we can move the Type declartion to an interface so that it can be used in multiple places
// now we will again re-write the genric function using interface for types
// declare type constraints as interface
type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var result V
	for _, value := range m {
		result += value
	}
	return result

}
