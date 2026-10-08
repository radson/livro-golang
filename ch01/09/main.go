// Fetch exibe o conteudo encontrado em cada URL especificada
// 1.9: Mostrar url e status http

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
	httpStatus := ""

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) {
			url = fmt.Sprintf("%s%s", httpPrefix, url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %\n", err)
			os.Exit(1)
		}
		httpStatus = fmt.Sprintf("HTTP Status: %d\n", resp.StatusCode)
		_, err = io.Copy(os.Stdout, resp.Body)
		_, err = io.Copy(os.Stdout, strings.NewReader(httpStatus))
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s[%d]: %v\n", url, httpStatus, err)
			os.Exit(1)
		}
	}
}
