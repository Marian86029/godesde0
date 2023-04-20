package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es el primer mensaje")
	defer fmt.Println("Este es el ultimo mensaje")
	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero panic\n %v", reco)

		}
	}() // no puede existir un recover sin un defer y una funcion anonima
	a := 1 // el recover sirve para continuar la logica de la funcion despues de encontrarse con un panic y posteriormente guardar esa informacion en un log para tener un protocolo
	if a == 1 {
		panic("Se encontro el valor 1")
	}

}
