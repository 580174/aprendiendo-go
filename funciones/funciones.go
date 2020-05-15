package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	saludo := "Buenso Dias  ComPañeRos"
	cambiar(&saludo) //para pasar la direccion de saludo usamos el operador de direccion
	fmt.Println(saludo)
	fmt.Println(sumar(4, 5))

	min, mayus := convertir(saludo)
	fmt.Println(min, "\n", mayus)

	//retona dos datos el primero es el contenido y el 2do es el error
	contenido, err := ioutil.ReadFile("apuntes.txt")

	//si ocurre el error entonses la variable err es diferente de nil, entonses detenemos el programa con la sentencia return
	if err != nil {
		fmt.Println("ocurrio un error ")
		return
	}

	fmt.Println(string(contenido))

	numeros := []int{2, 4, 6, 0, 50, 2, 4, 7, 9, 11, 60, 30}

	resultadov := filtrar(numeros, func(num int) bool {
		if num >= 10 {
			return true
		} else {
			return false
		}

	})

	fmt.Println(resultadov)

	r := hello("blas")(25)

	fmt.Println(r)

	fmt.Println(sumarArgumentos())

	funcAnonimas()
}

/*
	funcion de paso de parametros por valor, por que estring que se resive por parametro es una
	copia de valor d la variable saludo
*/
func saludar(saludo string) {
	fmt.Println(saludo)
}

/*
	funciones por referencia
*/
func cambiar(saludo *string) {
	*saludo = "Soy automata saludos"
}

/*
	funcions que retornan valores
	ojo si los parametros son del mismo tipo solo nesesitmos espesificar el ultimo parametro con su tipo
	y se sobreentiende que todos son de tipo espesifico
*/

func sumar(num1, num2 int) int {
	return num1 + num2
}

/*
	opedemos retornar mutiples valores en las funciones

	- si una funcion retorna mas de un valor deve espesificar el tipo de retorno entre ()
		separado por coma

*/

func convertir(texto string) (string, string) {
	min := strings.ToLower(texto)
	may := strings.ToUpper(texto)

	return min, may
}

/*
	las funciones son tipos de datos asi que podemos resivir como parametros y retornalas

	ejm resibir como parametro una funcion
*/
func filtrar(nums []int, callback func(int) bool) []int {
	result := []int{}

	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}

	return result
}

/*
	funcion que retorna una funcion
*/
func hello(nombre string) func(uint8) string {
	return func(edad uint8) string {
		return "holaa " + nombre + " tines " + string("25") + " años"
	}
}

/*
	funciones variaticas
	--------------------
	nos permiten resivir un numero n de variablesdes del mismo tipo
	en este caso el parametro numeros es un slice de tipo de datos int
*/

func sumarArgumentos(numeros ...int) int {
	r := 0
	for _, v := range numeros {
		r += v
	}

	return r
}

func funcAnonimas() {

	//funcion anonima
	x := func() string {
		return "texto"
	}

	//funcion anonima autoejecutable, tanbien pueden resivir argumentos
	func(txt string) {
		fmt.Println("yo me auto ejecuto", txt)
	}("texto")

	fmt.Println(x())
}
