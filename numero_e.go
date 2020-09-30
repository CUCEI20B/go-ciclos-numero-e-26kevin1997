package main

import "fmt"

func main() {

	var x int
	var aux float64 = 1
	var res float64 = 0

	fmt.Println("¿Cuantos calculos quieres hacer?")
	fmt.Scan(&x)

	for i := 0; i < x; i++ {
		if i != 0 {
			aux = aux / float64(i)
		}
		res = res + aux
	}
	fmt.Println(res)
}
