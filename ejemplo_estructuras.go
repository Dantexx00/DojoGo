package main

import "fmt"

type persona struct{
	nombre string
	estatura float32
}

func(c *persona) correr() string {
	return c.nombre+" corriendo"
}

func main() {
	
	c:=persona{"Raul",12}
	fmt.Println(c.correr())

}



/*EJERCICIO: Escribir un programa que encuentre el numero menor en el 
siguiente arreglo y lo imprima:

x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}*/