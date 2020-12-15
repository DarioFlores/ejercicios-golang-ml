package main

import "fmt"

func main() {
	suma, resta := sumaYResta(5, 2);
	fmt.Println("Suma:", suma);
	fmt.Println("Resta:", resta);

	suma2, _ := sumaYResta(15, 10);
	fmt.Println("Suma:", suma2);
}

func sumaYResta(x int, y int) (int, int) {
	return x+y, x-y;
}
