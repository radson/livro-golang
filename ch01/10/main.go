// 1.10 fetchall exibindo a saida em um arquivo

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "erro ao criar arquivo: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Inicia uma gorrotina para cada URL
	}
	for range os.Args[1:] {
		s := <-ch
		fmt.Println(s) // Imprime a saida de cada gorrotina
		fmt.Fprintln(file, s)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds()) // Imprime o tempo total
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // Evita vazamentos de recursos
	if err != nil {
		ch <- fmt.Sprintf("erro ao ler %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, nbytes, url)
}
