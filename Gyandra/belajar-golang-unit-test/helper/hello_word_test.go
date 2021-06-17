package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWordAssert(t *testing.T) {
	result := HelloWord("Gya")
	assert.Equal(t, "Hello Gya", result, "Result must be 'Hello Gya'")
	fmt.Println("TestHelloWord with Assert Done")
}

func TestHelloWordGya(t *testing.T) {
	result := HelloWord("Gya")

	if result != "Hello Gya" {
		// error
		t.Error("Result must be 'Hello Gya'")
	}

	fmt.Println("TestHelloWordGya Done")
}


func TestHelloWordNaufal(t *testing.T) {
	result := HelloWord("Naufal")

	if result != "Hello Naufal" {
		// error
		t.Fatal("Result must be 'Hello Naufal'")
	}

	fmt.Println("TestHelloWordNaufal Done")
}