package main

import "fmt"

func main() {
	/*
		LOS ARRAYS	 tienen un tamaño fijo locual despues de declararro no puede cambiar su cantidad de posiciones
	*/
	var frutas [3]string
	frutas[0] = "lima"
	frutas[1] = "sandilla"
	frutas[2] = "naranja"

	fmt.Println(frutas)

	/*
		los arrays literales
		podemos declarar y asignar en una sola linea
	*/

	comidas := [3]string{"amburguesa", "salchipapa", "hotdoc"}

	fmt.Println(comidas)

	/*
		Los eslices mos permiten trabajar con arrays de forma dinamica son apuntadores a arrays y estas no contienen valores
		slices no poseen datos, son apuntadores a array.
	*/

	set := [7]string{"rojo", "amarillo", "verde0", "azul", "american", "cpp", "colors"}

	/*
		este es un scile de colores el segundo parametro es excluyente, ademas se puede omitir el primer el 1er
		parametro en este caso se tomara desde el primer elemento del array, agual manera el segundo parametro se
		puede omitir en este caso se tomara des el primer parametro asta el final del array y ulto se puede omitir los dos
		parametros en este caso apuntara a todo e array
	*/

	colores := set[:4]
	marcas := set[4:]

	marcas[0] = "farba" // esto demuestra los slices no poseen datos, son apuntadores a array.set

	fmt.Println(colores)
	fmt.Println(marcas)
	fmt.Println(set[:]) //si no espesifico ni el indice inicial ni el indice final se implrimiran todos los elemento de array

	/*
		los slices tienen tamaño y capacidad
	*/

	comida := [5]string{"naranja", "limon", "fresa", "amburquesa", "hotdoc"}
	frutas_ := comida[1:3]

	fmt.Println("Tamaño: ", len(frutas_))   // tamaño del slice frutas_
	fmt.Println("capacidad ", cap(frutas_)) //capacidad es el numero de elementos apartir desde donde esta apuntando el slice hasta el final del array

	/*
		apuntara al array comida hasta que rebase su capacidad, cuando rebasa deja de apuntar
		la funcion append retorna un nuevo array[x*x] y el slice frutas_ apunta a ese nuevo array
	*/
	frutas_ = append(frutas_, "platano", "uva")
	fmt.Println(comida)
	fmt.Println(frutas_)
	/*
		se duplica la capacidad del slice frutas_ al rebasar la capasidad inicial que tiene
	*/
	fmt.Println("capacidad ", cap(frutas_))

	/*
		creando slices desde cero y no desde rangos que referencian a otro array sino a un array nuevo
		y como: al igual q un array pero no espesificamos la cantidad de posiciones
		-------------------------


	*/

	/*
		slice literal ----como se ahacen, 1ro construye un array donde alamacena los elementos y retorna un
		slice que referencia a ese array que el creso
	*/
	//deportes := []string{"futbol", "basket", "boley"}
	/*
		creando arrays con la funcion preconstruida make(tipo slice, tamaño, capacidad)
	*/
	libros := make([]string, 0, 3)

	libros = append(libros, "hamlet", "eloquent", "doom", "pio paco")

	fmt.Println("tmaño", len(libros))
	fmt.Println("capacidad ", cap(libros))
	fmt.Println(libros)

	//cual es el valor cero del los slices: nill. cuando no lleva las llaves asi yano lugares := []string{} xp es un slice literal inicialisado
	var lugares []string

	fmt.Println(lugares == nil)

	/*
		MAPAS
		-----------------------------------------
		Son estructuras de clave valor sintaxis
		nombreMapa := make(map[string]string) //espesificar tipo de dato clave y valor
	*/

	luminaria := make(map[string]string)
	luminaria["nombre"] = "foco1"
	luminaria["estado"] = "apagado"

	fmt.Println(luminaria)

	/*
		mapas literales, propiedades son separados con coma asta el ultimo
	*/

	persona1 := map[string]string{
		"nombre":       "blas humani",
		"estado sivil": "soltero",
	}

	//ponemos eliminar sus elemntos de una mapa con delete(nombreMapa, "llave")
	delete(persona1, "estado sivil")
	fmt.Println(persona1)
	fmt.Printf("valor %q\n", persona1["no existe"])

	//obtener elemento de una mapa con la notacion corchete []. retorna 2 valores: valor y  ok. si exite la llave espesificado
	//persona1["nonbre llave no existe"]sino exite la llave en la mapa retorna el valor cero segun el tipo de dato espesificado como valor en este caso estring
	contenido, ok := persona1["nombre"]

	fmt.Print(contenido, " ", ok)

	if _, ok := persona1["sexo"]; !ok {
		fmt.Println("no existe, creemelo")
		persona1["sexo"] = "macho"
	}

	fmt.Println(persona1)

	/*
		ESTRUCTURAS EN GO
		-----------------------------------------------------
		permiten almacenar campos
		sintaxis
	*/

	type curso struct {
		nombre   string
		profesor string
		ciudad   string
		edad     byte
	}

	//para instanciar un estructura creamos una variable y asignamos sus valores
	db := curso{
		nombre:   "bases de datos",
		profesor: "alexys",
		ciudad:   "mexico",
		edad:     50,
	}

	fmt.Printf("%+v\n", db) //para mostrar todos los campos clave y valor

	/*otra forma de crear instancias atrabes de una estructura literal como:
	puedo indicar unicamente el valor que tendra cada uno de los campos,
	*/

	git := curso{"curso gig", "beto", "mexico", 25}
	css := curso{profesor: "alvaro"} //ojo solo este tipo de instancia nos permite asignar un avalor a algunos campos
	fmt.Printf("%+v", git)
	fmt.Printf("%+v", css)
	fmt.Println(css.profesor) //para acceder a un campo espesifico podemos usar la notacion punto

	/*
		podemos crear punteros a estructuras
	*/

	p := &css

	(*p).nombre = "curso de css" // en go en el caso de las estructuras no es nesesario desreferenciar el puntero de esta forma
	p.ciudad = "lima peru"       //la notacion punto es usado para acceder al campo y es automaticamente desreferenciado

	fmt.Printf("%+v", *p)

}
