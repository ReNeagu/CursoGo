package main

import "fmt"

func main() {
	fmt.Println("Array e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	array1[1] = "Posição 2"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 1"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	//ARRAYS INTERNOS
	fmt.Println("----------")
	slice3 := make([]float32, 10, 11) //formato, tamanho, limite max
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5) //nao passei a capacidade dele
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //length
	fmt.Println(cap(slice4)) //capacidade
}
