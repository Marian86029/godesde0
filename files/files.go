package files

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Marian86029/godesde0/ejercicios"
)

var fileName string = "/Users/marianoromera/Desktop/godesde0-1/files/txt"

func GrabaTabla() {
	var texto string = ejercicios.PublicParameter()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al mostrar el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.PublicParameter()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")

	}
}
func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	if err != nil {

		fmt.Println("Error durante el WriteString" + err.Error())

		return false
	}
	arch.Close()
	return true

}
func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	fmt.Println(string(archivo))

	/*    ////!!!!!!OTRA SOLUCION!!!!!!!//
		func LeoArchivo (){
			archivo, err := os.Open(fileName)
			if err != nil {
				fmt.Println("Hubo un error leyendo el archivo"+err.Error())
				return
			}
			scanner := bufio.NewScanner(archivo)    //!!!!!! NECESITO EL PACKAGE "BUFIO" EN ESTE CASO!!!!!////

		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Println(">"+registro)
		}

	archivo.Close()
		}





	*/

}
