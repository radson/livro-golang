// Fetch exibe o conteudo encontrado em cada URL especificada
// 1.8: Prefixo http:// seja acrescentado caso esteja faltando.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	httpPrefix := "http://"

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) {
			url = fmt.Sprintf("%s%s", httpPrefix, url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
