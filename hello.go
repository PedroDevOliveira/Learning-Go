package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Pedro"
	var version float32 = 1.0
	var age int
	var testFloat = 1.1
	inferType := "Essa é uma forma diferente de declarar variaveis"

	fmt.Println("O tipo da variavel testFloat é: ", reflect.TypeOf(testFloat))

	fmt.Println("Olá ", name)
	fmt.Println("Este programa está na versão ", version)
	fmt.Println("Sua ideade é ", age)
	fmt.Println(inferType)
}
