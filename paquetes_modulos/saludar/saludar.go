package saludar

/*
	variables a nivel de paqute se declaran fuera de cualquier funcion,
	para declarar no esta permitido el operador :=
*/

var Variable string = "variable anivel de paquete"

/*
	los paquetes
	------------
	 es una carpeta que continen una coleccion de archivos q nos proveen una funcionalidad en este caso seria
	 la funcinalidad del paquete saludar "saludar en diferentes leguajes"

	 los nombres de paquetes deben ser claros y consisos, que indiquen claramente que hace el paquete

	 ejm
	 de otro lado se llamaria asi:

	 buena practica
	 -------------
	 saludar.Ingles()

	 mala practica
	 --------------
	 saludar.saluderEnIngles() es redundante el prefijo saludar
*/

func Ingles() string {
	return "Hello goog morning"
}

func Italiano() string {
	return "bondia parlachi"
}
