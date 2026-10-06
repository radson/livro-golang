// Exibir o índice e o valor de cada um dos os.Args

package main

import (
	"fmt"
	"os"
)

func main() {
	for i := range os.Args {
		fmt.Printf("Indice [%d], Nome do comando: %s\n", i, os.Args[i])
	}
}
