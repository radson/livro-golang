// Exibir o os.Args[0], que é o nome do comando que o chamou.

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Nome do comando: %s", os.Args[0])
}
