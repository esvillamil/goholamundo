package holamundo

import (
"fmt"
"math/rand"
"time"
)

// init initializes the package.
//
// This function is called automatically when the package is imported.
// It sets the random seed using the current time.
func init () {
fmt.Println("HolaMundo en init_servicio␣...")
rand.Seed(time.Now().Unix())
}

// ImprimeHolaMundo prints "HolaMundo en funcion" in the console.
//
// This function does not take any parameters.
// It does not return any values.
func ImprimeHolaMundo () {
fmt.Println("HolaMundo en funcion")
}

// ImprimeHolaMundoEnJapones prints "HolaMundo en Japónes: ハロー・ワールド" in the console.
//
// This function does not take any parameters.
// It does not return any values.
func ImprimeHolaMundoEnJapones () {
fmt.Println("HolaMundo en Japónes: ハロー・ワールド")
}

// ImprimeHolaMundoEnArabe prints "HolaMundo en Arabe: مرحبا بالعالم" in the console.
//
// This function does not take any parameters.
// It does not return any values.
func ImprimeHolaMundoEnArabe () {
fmt.Println("HolaMundo en Arabe: مرحبا بالعالم")
}
