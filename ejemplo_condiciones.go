package main

import "fmt"

func main() {
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 {
      if i % 5 ==0{
      	fmt.Println(i, "GoTeam")
      }else{
      	fmt.Println(i, "Go")
  			}
    } else if i % 5 == 0{
      if i % 3 ==0{
      	fmt.Println(i, "GoTeam")
      }else{
      	fmt.Println(i, "Team")
  			}

    }
  }
}

/*EJERCICIO: Escriba un porograma que recorra los numeros del 1 al 100. 
	Pero para los multiplos de 3 imprima "Go" y para los multiplos de 5
  imprima "Team". Para los numeros que son multiplos de 3 Y 5 imprima
"GoTeam".*/