package main

/*
	import "fmt"
	estos son los paquetes creados por go, y no podemos hacer lo mismo con los paquetes que emos creado o de
	terceros

	para eso existen los Modulos
	----------------------------
	- que nos permiten administras nuestras dependencias de nuestros paquetes y
	- controlar las versiones de los mismos

	se recomienda tener 1 solo en la raiz de nuestro proyecto
	como se crea asi:
	go mod init github.com/580174/aprendiendo-go.git // se recomienta q el nombre del modulo sea la misma ruta d nuestro repositorio
	para que tengan acceso otros desarrolladores.  por que nuestro repositorio es publico si otro desarrolladors
	quisieran utilizar nuestros paquetes que es tan en nuestro repositorio, unicamente lo importaria atrabes de
	esa ruta..... go no desarrollo nuevas herramientas para la gestion de paquetes utilizo los ya existentes los sistemas
	de control de versiones. para importar un paquete le espesificamos la ruta de nuestro repo GO atrabes de git clona ese
	repositorio y descarga el codigo fuente para nosotros usarlo


*/
import (
	"fmt"

	"github.com/580174/aprendiendo-go/paquetes%modulos/saludar"
)

func main() {
	/*

	 */
	fmt.Println(saludar.Ingles())
}
