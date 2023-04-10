package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 59, 64} // Asignacion Directa de todo el conjunto
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33 //Asignacion Directa de un vector
	tabla[2] = 54 // En consola se mostrara este valor sin importar la especificacion en la primera instancia (Asignacion directa del Conjunto)

	tabla2 := [10]string{"Mariano", "Agustin"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ { // Esta es la manera de recorrer vectores
		fmt.Println(tabla[i])
	}

	matriz[5][24] = 15
	fmt.Println(matriz)
}
