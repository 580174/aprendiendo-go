package main

import "fmt"

func main()  {
	/*
	 tipos de datos basicos en go son 3: bool, string, numeric
	*/
	condicion := true
	numero := 12
	cadena := "cadena de texto"

	fmt.Printf("Tipo %T , Valor %v \n", numero, numero);
	fmt.Printf("Tipo %T , Valor %v \n", condicion, condicion);
	fmt.Printf("Tipo %T , Valor %v \n", cadena, cadena);

	/*
		variedades de tipos de datos numerico se agrupan en dos grandes grupos uint y int

		* uint8 el conjunto de todos los enteros de 8 bits sin signo (0 a 255) 
		* uint16 el conjunto de todos los enteros de 16 bits sin signo (0 a 65535) 
		* uint32 el conjunto de todos los enteros de 32 bits sin signo (0 a 4294967295) 
		* uint64 el conjunto de todos enteros de 64 bits sin signo (0 a 18446744073709551615) 

		* int8 el conjunto de todos los enteros de 8 bits con signo (-128 a 127) 
		* int16 el conjunto de todos los enteros de 16 bits con signo (-32768 a 32767) 
		* int32 el conjunto de todos los 32- enteros de bits (-2147483648 a 2147483647) 
		* int64 el conjunto de todos los enteros de 64 bits con signo (-9223372036854775808 a 9223372036854775807) 

		float32 el conjunto de todos los números de punto flotante IEEE-754 de 32 bits 
		float64 el conjunto de todos los IEEE-754 de 64 bits Números de punto flotante

		complex64 el conjunto de todos los números complejos con float32 partes reales e imaginarias 
		complex128 el conjunto de todos los números complejos con float64 partes reales e imaginarias 

		byte alias para uint8 
		rune alias para int32, ademas representa un codigo unicode 'a' = 97
	*/

	var unicode rune = 'A';

	fmt.Println(unicode);

	/*
		no podemos realizar operaciones entre tipos de datos diferentes aun cuando sean numeros, entonses necesitamos
		casetar ejm.
	*/

	var a uint8 = 100
	var b uint16 = 23600

	var c uint16 = uint16(a) + b;

	fmt.Println(c);

	/*
		operador blank _ se usa para mentener la variable sin usar ejm.
	*/

	var _ uint8 = 123
	_ = 124

	/* 
		sino le asignamos un valor a una variable go asigna por defecto un valor
	*/

	var sinValor1 string
	var sinValor2 uint8
	var condicionX bool

	fmt.Printf("valor %q\n", sinValor1)
	fmt.Printf("valor %v\n", sinValor2)
	fmt.Printf("valor %v\n", condicionX)



}