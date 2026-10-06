// Exibir o nome de todos os arquivos em que cada linha duplicada aparece.
// go run exer-1.4.go ocorrencias.txt ocorrencias.txt

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "%s: %v\n", os.Args[0], err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[fmt.Sprintf("%s\t%s", f.Name(), input.Text())]++
	}
	// Nota: ignorando erros em potencial de input.Err()
}
