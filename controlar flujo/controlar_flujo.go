package main

import (
	"fmt"
	"os"
)

/*
	ejemplificacion
	- las funciones con defer se guardan en la pila, segun en orden que se enuentra en el programa
	  osea el 2do encontrado se pone encima del primero, asi sisesivamente. el ultimo se ejecuta primero

	PILA
	---------------------
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")
*/
func main() {
	/*
		Aprende a controlar el flujo de tu programa con las funciones defer, panic y recover

		defer .- aplaza la ejcucion de una puncion haste que finaliza todo el programa
	*/

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	/*
		una particularidad que tienen las funciones subidos a la pila con defer es que sus argumentos son elavuados
		inmediatamente y guardados
	*/
	x := 5
	defer fmt.Println("Defer", x)

	x = 10

	fmt.Println("se ejecuto 1ro", x)

	/*
		------------------------------------------------------------------------------------------------------------------------
			casos pracitos del defer
			sualmente usamos para limpiar recursos,
			- para cerrar archivos,
			- cerrar conexion de red
			- cerrar controladores de bases de datos

			ejm
			creamos un archivo > escribimos es archivo y la cerramos con defer
	*/

	file, err := os.Create("hello.txt") //retorna un puntero a una estructura FILE o un error

	if err != nil {
		fmt.Printf("Error al crear el archivo. %v", err)
		return //se detine el programa

		//escribimos el archivo con el metodo write de file
		//esta funcion devuelve una variable n q son en numero de bytes que fueron escritos, y un error
	}

	defer file.Close() //esto se ejecutar al final  y cerrara el archivo

	_, err = file.Write([]uint8("saludos yo la escribi este archivo"))
	if err != nil {
		// sin usar defer se llamaria aqui el metodo close de file para serrar
		fmt.Printf("Error al escribir el archivo. %v", err)
		return //se detine el programa
	}

	/*
		otra bes al finalizar el programa se llamaria al metodo close de file, por el tamaño de nuestro programa
		nos pedemos olvidar de cerrar el archivo
	*/

	fmt.Println(division(12, 2), "_______")
	fmt.Println(division(10, 0), "_______")
	fmt.Println(division(12, 6), "_______")

}

/*
	la sentencia panic detiene la ejecucion del programa y me muestra la pila de llamadas para rastrear donde se producio los errores
*/
func division(dividendo, divisor int) int {

	/*
	 esta funcion anomina se ejecuta
	 - cuando se finaliza el programa osea al llamar validarDivision que esto detiene el progrma cuando se pasa 0 como argumento
	 - o cuando se finaliza o retona la funcion

	*/
	defer func() {
		if r := recover(); r != nil { //recover nos devuelve el valor dem panico
			fmt.Println("Nos recuperamos del panico. ", r)
		}
	}()

	validarDivision(divisor) //si divisor es 0 se finaliza la ejecucion del programa

	return dividendo / divisor
}

func validarDivision(divisor int) {
	if divisor == 0 {
		panic("no se puede dividir por cero") // ojo panic detiene la ejecucion del programa
	}
}

/*
	como hacer que nuestro programa continue su ejecucion despues de producirse un panico
	como nos recuperamos del panico
*/
