package main

import "fmt"

func main() {
	var point = 8

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Println("Tidak lulus, nilai anda %d \n", point)
	}

	var point2 = 6

	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}
