package main

import "fmt"

func main() {
	/*
		operadores aritmeticos y su presedencia
		(), *, /, %, +, -
	*/

	var _ uint8 = 4 + 6*3 // -> 22

	/*
		operadores de asignacion
		=, +=, -=, *=, /=, %=
	*/

	var cantidad = 5

	cantidad += 5

	fmt.Println(cantidad)

	/*
		pos-incremento y pos-decremento ++, --
		ojo no existe pre-incremento ni pre-decremento en go y no son una exprecion sino una declaracion por eso no imprime Println
		esto no esta permitido
		var num = 5
		num = num++ + 5;
	*/

	var c = 5
	c++

	fmt.Println(c)

	/*
		operadores de comparacion
		>, <, ==, !=, >=, <=
		ejm.
	*/

	fmt.Println(4 > 5)

	/*
		opradores logicos
		son los que me permiten agrupar comparaciones para dar una respuesta boleana &&, ||
	*/
	var edad = 20
	fmt.Println("adulto ", edad > 18 && edad < 80)

	/*
		operador unario de negacion !
	*/

}
