package main

import (
	"github.com/Marian86029/godesde0/middleware"
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



	numero, texto := ejercicios.PublicFunction("500")
	fmt.Println(numero)                                     // Primer Intento de Proyecto
	fmt.Println(texto)

	teclado.IngresoNumeros()     // Scanear info en el teclado

	iteraciones.Iterar()*/ // ciclo for go

	// fmt.Println(ejercicios.PublicParameter())

	// files.GrabaTabla()

	//files.SumaTabla()

	//files.LeoArchivo()

	//funciones.Llamarclosures()

	//funciones.Exponencia(2)

	//arreglos_slices.MuestroArreglos()

	//arreglos_slices.Capacidad()

	//mapas.MostrarMapas()
	// user.AltaUsuario()

	//Pedro := new(modelos.Hombre)
	//e.HumanosRespirando(Pedro)
	// esto es el mismo ejercicio
	//Maria := new(modelos.Mujer)
	//e.HumanosRespirando(Maria)

	//d.EjemploPanic()

	/*canal1 := make(chan bool)

	go g.MiNombreLentooo("Mariano Romera", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy Aqui")*/

	//webserver.MiWebServer()

	middleware.MiMiddleware()

}
