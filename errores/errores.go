package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := dividir2(12, 0)

	if err != nil {
		fmt.Printf("Ocurrio un error %v ", err)

		return
	}

	fmt.Println(result)

}

func dividir(dividendo, divisor uint64) (uint64, error) {
	if divisor == 0 {
		return 0, errors.New("No puedes dividir con 0")
	}

	return dividendo / divisor, nil
}

/*
	los valores que se retornan los podemos devolver con nombres y estos tinen valor cero
*/

func dividir2(dividendo, divisor uint64) (result uint64, err error) {
	if divisor == 0 {
		err = errors.New("No puedes dividir con 0")
		return
	}

	result = dividendo / divisor
	return
}
