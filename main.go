// Ejercicio de hola mundo
package main

import (
"fmt"

"github.com/esvillamil/goholamundo/holamundo"
)

import _ "github.com/esvillamil/goholamundo/holamundo"

func init () {
fmt.Println("Hola mundo en init ...")
}


func main() {
saludo := "¡Hola mundo!"
fmt.Println( saludo )
//fmt.Println(holamundo.Translate ("Arecibo␣message␣sent"))
holamundo.ImprimeHolaMundo()
holamundo.ImprimeHolaMundoEnJapones()
holamundo.ImprimeHolaMundoEnArabe()
}