package ejercicios

import (
	"strconv"
)

func PublicFunction(texto string) (int, string) {

	num, err := strconv.Atoi(texto)

	if err != nil {
		return 0, "hubo error" + err.Error()
	}

	if num > 100 {
		return num, "es mayor a 100"

	} else {
		return num, "es menor a 100"

	}

}
