// Exibir o índice e o valor de cada um dos os.Args
//  go run exer-1.3.go ola mundo cruel maneiro cool show de bola

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	const count = 100000 // Aumentamos a carga
	s := ""

	// Teste A: Concatenação Manual (Lenta)
	start_a := time.Now()

	for i := 0; i < count; i++ {
		sep := ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
	}

	// fmt.Println(s)
	fmt.Printf("Manual:\t  %.10fs\n", time.Since(start_a).Seconds())

	// Teste B: strings.Join (Rápida/Otimizada)
	start_b := time.Now()
	for i := 0; i < count; i++ {
		_ = strings.Join(os.Args[1:], " ")
	}
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Join:\t    %.10fs\n", time.Since(start_b).Seconds())
}
