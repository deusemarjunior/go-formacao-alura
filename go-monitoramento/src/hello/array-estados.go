package main

import (
	"fmt"
)

func main() {
	estados := devolveEstadosDoSudeste()
	fmt.Println(estados)
	fmt.Println(exibeEstadosSlice())
}

func devolveEstadosDoSudeste() [4]string {
	var estados [4]string
	estados[0] = "RJ"
	estados[1] = "SP"
	estados[2] = "MG"
	estados[3] = "ES"
	return estados
}

func exibeEstadosSlice() []string {
	estados := []string{"RJ", "SP", "MG", "ES"}
	return estados
}
