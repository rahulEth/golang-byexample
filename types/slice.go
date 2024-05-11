package main

import "fmt"

func main() {
	var i = [6]int{1, 2, 3, 4, 5, 6}

	var slice []int = i[1:4] //incldue low indice or bound exclude high bound
	fmt.Println(slice)

	var str = [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	slice1 := str[0:2]
	slice2 := str[1:3]
	fmt.Println(slice1, slice2)
	slice1[0] = "XXX" // Changing the elements of a slice modifies the corresponding elements of its underlying array.

	fmt.Println(slice1, slice2)
	fmt.Println(str)

	q := []int{2, 3, 4, 5, 6}
	fmt.Println(q)
	r := []bool{true, false, true}
	fmt.Println(q, r)
	s := []struct {
		i int
		j bool
	}{
		{2, true},
		{3, false},
		{4, true},
	}
	fmt.Println("sssssss ", s)

	t := []int{2, 3, 4, 5, 6, 7}
	t = t[1:4]
	fmt.Println("1....... ", t)
	t = t[:2]
	fmt.Println("2........ ", t)
	t = t[1:]
	fmt.Println("3........ ", t)
	t = t[:]
	fmt.Println("4....... ", t)

	// below slice expression are equilentconst

	// var a[10]int
	// a[0:10]
	// a[:10]
	// a[0:]
	// a[:]

}
