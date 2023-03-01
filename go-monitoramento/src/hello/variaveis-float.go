package main

import (
	"fmt"
	"reflect"
)

func main() {

	var precoLeite float32 = 2.99
	var precoOvo float64 = 3.99
	//var precoPao float = 0.99
	precoPao := 0.99

	fmt.Println("O preço do leite é R$", precoLeite)
	fmt.Println("O preço do ovo é R$", precoOvo)
	fmt.Println("O preço do pão é R$", precoPao)
	fmt.Println("variavel precoPao ", reflect.TypeOf(precoPao))
}
