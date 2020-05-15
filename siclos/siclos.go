package main

import "fmt"

func main() {
	/*
		en caso de go solo existe el siclo for pero con diferentes variantes

		for clasico
		-----------
	*/

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/*
		siclo for continuo
		------------------
		es similar al while en otros lenguages
	*/

	cantidad := 10

	for cantidad >= 0 {
		fmt.Println(cantidad)
		cantidad--
	}

	/*
		siclo forever
		-------------
		es un for que dura parasimpre < uso en casos como tienen que ser escuchados permanentemente como sockets
	*/

	x := 0

	for false {
		fmt.Println(x)
		x++
	}

	/*
		el siclo for range

		me permite iterar sobre mapas slices y strings
	*/

	slice := []uint8{2, 4, 6, 8, 10}

	//la instrucion for range me devuleve 2 valores: indece y valor < este valor es un copia
	for i, v := range slice {

		v *= 2        //el valor que devulve v es una copia de valor del array slice esto no afecta nada al array slice
		slice[i] *= 2 //esto reemplasa con nuevos valores al array slice

	}

	fmt.Println(slice)

	//iterar sobre mapas
	deportes := map[string]string{
		"futbol":  "balon",
		"basquet": "balon basquet",
	}
	for k, v := range deportes {
		fmt.Println(k, v)
	}

	//iterar strings

	saludo := "hola a todos"

	for _, v := range saludo {
		fmt.Print(string(v)) //tenemos que usar el operador de casting para convertir los numeros unicode a string
	}
}
