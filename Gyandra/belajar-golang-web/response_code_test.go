package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
	"io"
	"net/http/httptest"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
 	name := request.URL.Query().Get("name")

 	if name == "" {
 		writer.WriteHeader(http.StatusBadRequest)
 		fmt.Fprint(writer, "name is empty")
 	} else {
 		fmt.Fprintf(writer, "Hello %s", name)
 	}
 }

 func TestResponseCodeInvalid(t *testing.T) {
 	request := httptest.NewRequest("GET", "http://localhost:8888", nil)
 	recorder := httptest.NewRecorder()

 	ResponseCode(recorder, request)

 	response := recorder.Result()
 	body, _ := io.ReadAll(response.Body)

 	fmt.Println(response.StatusCode)
 	fmt.Println(response.Status)
 	fmt.Println(string(body))
 }