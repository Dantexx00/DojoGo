package main

import "fmt"

func main() {
  fmt.Print("Ingrese el texto: ")
  var input string
  fmt.Scanf("%127s",&input)

  output := input

  fmt.Println(output)
}

/*EJERCICIO: Cambie el anterior programa para que en vez de capturar
	un decimal, capture un texto y lo imprima.*/