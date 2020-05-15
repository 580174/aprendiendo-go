package main

import "fmt"

func main()  {
	/*
		no son nesesarios los parentesis
	*/

	dia := "miercoles"

	if dia == "lunes" {
		fmt.Println("es el dia lunes")
	} else if dia =="martes"{
		fmt.Println(" es el dia martes")
	}else{
		fmt.Println("es el dia ", dia)
	}

	/*
		podemos declarar variables dentro del scope del if asi como en este caso la variable sexo que 
		solo vivira dentro desta condicion numtiple
	*/

	if sexo := "c"; sexo == "v"{
		fmt.Println("es un varon")
	}else if sexo == "m"{
		fmt.Println("es una mujer")
	}else{
		fmt.Printf("el sexo es: %s\n", sexo)
	}

	/*
		la estructura de ontrol switch
		- no es necesario declarar el break
		- podemos usar operadores logicos y de comparacion en cada uno de los casos. no es neseario colocar la variable que vamos usar
		- podemos evaluar varias veses en cada uno de los casos separado por coma

	*/

	mes := "abril"

	switch mes{
		//aqui se agrupan los casos mas comunes
		case "enero" , "febrero":
			fmt.Println("estamos en verano")
		case "marzo", "abril", "mayo":
			fmt.Println("estamos en invierno")
	}

	colores := "verde"

	switch{
		case colores == "rojo" || colores == "verde" ||colores == "azul":
			fmt.Println("son los colres primarios")
		default:
		fmt.Println("NO son los colres primarios")
	}
}