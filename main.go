package main

import (
	"fmt"

	"github.com/Marian86029/godesde0/ejercicios"
)

func main() {

	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/ // Primer Test

	///    !!!!! IMPORTANT !!!!!   ///

	// EXPLICACION IF, ELSE , ELSE IF, SWITCH //

	/*if os := runtime.GOOS; os == "Linux" {
		fmt.Println("esto no es windows", os)

	} else {
		fmt.Println("esto es windows")

	}

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("esto es Linux")
	case "Darwin":
		fmt.Println("esto es Darwin")
	default:
		fmt.Printf("%s\n", os)

	}
	*/

	numero, texto := ejercicios.PublicFunction(500)
	fmt.Println(numero)
	fmt.Println(texto)

}
