package main

import "fmt"

func main()  {
	/*
		los punteros son variables que almacenan en memoria la direccion de un valor
		operador de direccion de memoria &variable
	*/
	fruit := "manzana"
	var p *string // un apuntador a la variable fruta
	p = &fruit
	//como acceder al valor atrabes del operador de desreferenciacion o indireccion *nombrePuntero
	*p = "banana"
	fmt.Printf("tipo: %T valor: %s direccion: %v\n", fruit, fruit, &fruit)
	fmt.Printf("tipo: %T valor: %v desreferenciacion %s", p, p, *p)
}