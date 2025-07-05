package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [5]int
	arr1[0] = 10
	arr1[1] = 20
	fmt.Println("Array 1:", arr1)

	arr2 := [5]string{"one", "two", "three", "four", "five"}
	fmt.Println("Array 2:", arr2)
	
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array 3:", arr3)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)

	fmt.Println(reflect.TypeOf(arr1)) // Type of arr1
	fmt.Println(reflect.TypeOf(slice)) // Type of slice

	slice = append(slice, 6, 7, 8)

	fmt.Println("Updated Slice:", slice)

	sliceArray2 := arr2[0:3]
	fmt.Println("Slice from Array 2:", sliceArray2)

	sliceArray3 := make([]float32, 10, 11)
	sliceArray3 = append(sliceArray3, 1.1, 2.2, 3.3)
	fmt.Println("Slice from Array 3:", sliceArray3, "Length:", len(sliceArray3), "Capacity:", cap(sliceArray3))

	sliceArray4 := make([]float32, 5)
	fmt.Println("Slice from Array 4:", sliceArray4, "Length:", len(sliceArray4), "Capacity:", cap(sliceArray4))
}
