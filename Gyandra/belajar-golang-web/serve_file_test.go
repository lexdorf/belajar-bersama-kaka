package belajar_golang_web

import (
	"fmt"
	_ "embed"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resourcase/ok.html")
	}else {
		http.ServeFile(writer, request, "./resourcase/notfound.html")
	}
}

func TestServerFileServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8888",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resourcase/ok.html
var resourcaseOk string

//go:embed resourcase/notfound.html
var resourcaseNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourcase)
	}else {
		fmt.Fprint(writer, resourcaseNotFound)
	}
}

func TestServerFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8888",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}