package main

import (
	"bufio"
	"net/http"
	"os"
)

// write bufio to optimization
var writer *bufio.Writer

func Println(text string) {
	writer = bufio.NewWriter(os.Stdout)
	writer.WriteString("\n")
	writer.WriteString(text)
	writer.Flush()
}

func main() {
	mux := http.NewServeMux()

	// diretorio fisico
	fs := http.FileServer(http.Dir("./"))

	// mostra no browser localhost:8080/static
	mux.Handle("/", http.StripPrefix("", fs))
	Println("Run Server 8181")
	http.ListenAndServe(":8181", mux)
}
