// Modulo de mi primer programa en Go, no solo saluda ,
// sino que ademas lo traduce a otro idioma
package main

import (
	"fmt"

	"github.com/rserrano97/experiments/mod2"
)
func init () {
	fmt.Println("Ready to launch ...")
	}
func main() { 
	 saludo := "¡Hola!"
	 fmt.Println(saludo)
	 fmt.Println(mod2.Translate(saludo))
 } 